package config

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	vaultapi "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/dig"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RelayerConfiger interface {
	RelayerConfig() *RelayerConfig
}

func NewRelayerConfiger(getter kv.Getter) RelayerConfiger {
	return &ethereum{
		getter: getter,
	}
}

type ethereum struct {
	once   comfig.Once
	getter kv.Getter
}

type whitelist map[string]struct{}

type RelayerConfig struct {
	RPC                     *ethclient.Client
	RegistrationAddress     common.Address
	LightweightStateAddress *common.Address
	ChainID                 *big.Int
	PrivateKey              *ecdsa.PrivateKey
	WhiteList               whitelist
	nonce                   uint64
	GasLimitMultiplier      float64

	mut *sync.Mutex
}

func (e *ethereum) RelayerConfig() *RelayerConfig {
	return e.once.Do(func() interface{} {
		var result RelayerConfig

		networkConfig := struct {
			RPC                     *ethclient.Client `fig:"rpc,required"`
			RegistrationAddress     common.Address    `fig:"registration,required"`
			LightweightStateAddress *common.Address   `fig:"lightweight_state"`
			PrivateKey              *ecdsa.PrivateKey `fig:"private_key"`
			VaultAddress            string            `fig:"vault_address"`
			VaultMountPath          string            `fig:"vault_mount_path"`
			WhiteList               []string          `fig:"whitelist"`
			GasLimitMultiplier      float64           `fig:"gas_limit_multiplier"`
		}{}
		err := figure.
			Out(&networkConfig).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(e.getter, "network")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ethereum config"))
		}

		result.RPC = networkConfig.RPC
		result.RegistrationAddress = networkConfig.RegistrationAddress
		result.LightweightStateAddress = networkConfig.LightweightStateAddress

		result.ChainID, err = result.RPC.ChainID(context.Background())
		if err != nil {
			panic(errors.Wrap(err, "failed to get chain ID"))
		}

		result.PrivateKey = networkConfig.PrivateKey
		if result.PrivateKey == nil {
			result.PrivateKey = extractPrivateKey(networkConfig.VaultAddress, networkConfig.VaultMountPath)
		}

		result.nonce, err = result.RPC.NonceAt(context.Background(), crypto.PubkeyToAddress(result.PrivateKey.PublicKey), nil)
		if err != nil {
			panic(errors.Wrap(err, "failed to get nonce"))
		}

		result.WhiteList = make(whitelist, len(networkConfig.WhiteList))
		for _, address := range networkConfig.WhiteList {
			address = strings.ToLower(address)
			if result.WhiteList.IsPresent(address) {
				continue
			}

			result.WhiteList[address] = struct{}{}
		}

		result.GasLimitMultiplier = networkConfig.GasLimitMultiplier
		if result.GasLimitMultiplier == 0.0 {
			result.GasLimitMultiplier = 1.2
		}

		result.mut = &sync.Mutex{}
		return &result
	}).(*RelayerConfig)
}

func (n *RelayerConfig) LockNonce() {
	n.mut.Lock()
}

func (n *RelayerConfig) UnlockNonce() {
	n.mut.Unlock()
}

func (n *RelayerConfig) Nonce() uint64 {
	return n.nonce
}

func (n *RelayerConfig) IncrementNonce() {
	n.nonce++
}

// // ResetNonce sets nonce to the value received from a node
func (n *RelayerConfig) ResetNonce(client *ethclient.Client) error {
	nonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(n.PrivateKey.PublicKey), nil)
	if err != nil {
		return errors.Wrap(err, "failed to get nonce")
	}
	n.nonce = nonce
	return nil
}

func extractPrivateKey(vaultAddress, vaultMountPath string) *ecdsa.PrivateKey {
	conf := vaultapi.DefaultConfig()
	conf.Address = vaultAddress

	vaultClient, err := vaultapi.NewClient(conf)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize new client"))
	}

	token := struct {
		Token string `dig:"VAULT_TOKEN,clear"`
	}{}

	err = dig.Out(&token).Now()
	if err != nil {
		panic(errors.Wrap(err, "failed to dig out token"))
	}

	vaultClient.SetToken(token.Token)

	secret, err := vaultClient.KVv2(vaultMountPath).Get(context.Background(), "relayer")
	if err != nil {
		panic(errors.Wrap(err, "failed to get secret"))
	}

	vaultRelayerConf := struct {
		PrivateKey *ecdsa.PrivateKey `fig:"private_key,required"`
	}{}

	if err := figure.
		Out(&vaultRelayerConf).
		With(figure.EthereumHooks).
		From(secret.Data).
		Please(); err != nil {
		panic(errors.Wrap(err, "failed to figure out"))
	}

	return vaultRelayerConf.PrivateKey
}

func (w whitelist) IsPresent(address string) bool {
	_, ok := w[address]
	return ok
}
