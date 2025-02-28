package handlers

import (
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/registration-relayer/internal/service/requests"
	"github.com/rarimo/registration-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type txData struct {
	dataBytes []byte
	gasPrice  *big.Int
	gas       uint64
}

func Registration(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistrationRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	logF := logan.F{
		"user-agent": r.Header.Get("User-Agent"),
		"calldata":   req.Data.TxData,
	}
	if req.Data.Meta != nil {
		logF = logF.Merge(*req.Data.Meta)
	}

	log := Log(r).WithFields(logF)
	log.Debug("registration request")

	// `RelayerConfig(r).RegistrationAddress` is default value for target contract
	// if destination not specified this value will be used
	// this value is required in config
	registrationAddress := RelayerConfig(r).RegistrationAddress
	if req.Data.Destination != nil {
		if !RelayerConfig(r).WhiteList.IsPresent(*req.Data.Destination) {
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"data/destination": fmt.Errorf("specified contract address not allowed"),
			})...)
			return
		}

		// destination is valid hex address because of request validation
		registrationAddress = common.HexToAddress(*req.Data.Destination)
	}

	var txd txData
	txd.dataBytes, err = hexutil.Decode(req.Data.TxData)
	if err != nil {
		Log(r).WithError(err).Error("failed to decode data")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	RelayerConfig(r).LockNonce()
	defer RelayerConfig(r).UnlockNonce()

	err = confGas(r, &txd, &registrationAddress)
	if err != nil {
		Log(r).WithError(err).Error("failed to configure gas and gasPrice")
		// `errors.Is` is not working for rpc errors, they passed as a string without additional wrapping
		// because of this we operate with raw strings
		if strings.Contains(err.Error(), vm.ErrExecutionReverted.Error()) {
			errParts := strings.Split(err.Error(), ":")
			contractName := strings.TrimSpace(errParts[len(errParts)-2])
			errMsg := errors.New(strings.TrimSpace(errParts[len(errParts)-1]))
			ape.RenderErr(w, problems.BadRequest(validation.Errors{contractName: errMsg}.Filter())...)
			return
		}
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := sendTx(r, &txd, &registrationAddress, req.Data.NoSend)
	if err != nil {
		Log(r).WithError(err).Error("failed to send tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newTxResponse(tx))
}

func newTxResponse(tx *types.Transaction) resources.TxResponse {
	return resources.TxResponse{
		Data: resources.Tx{

			Key: resources.Key{
				ID:   tx.Hash().String(),
				Type: resources.TXS,
			},
			Attributes: resources.TxAttributes{
				TxHash: tx.Hash().String(),
			},
		},
	}
}

func confGas(r *http.Request, txd *txData, receiver *common.Address) (err error) {
	txd.gasPrice, err = RelayerConfig(r).RPC.SuggestGasPrice(r.Context())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}

	txd.gas, err = RelayerConfig(r).RPC.EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(RelayerConfig(r).PrivateKey.PublicKey),
		To:       receiver,
		GasPrice: txd.gasPrice,
		Data:     txd.dataBytes,
	})
	if err != nil {
		return fmt.Errorf("failed to estimate gas: %w", err)
	}
	txd.gas = uint64(float64(txd.gas) * RelayerConfig(r).GasLimitMultiplier)

	return nil
}

func sendTx(r *http.Request, txd *txData, receiver *common.Address, noSend bool) (tx *types.Transaction, err error) {
	tx, err = signTx(r, txd, receiver)
	if err != nil {
		return nil, fmt.Errorf("failed to sign new tx: %w", err)
	}

	if noSend {
		Log(r).WithField("hash", tx.Hash().String()).Warn("transaction sending disabled")
		return tx, nil
	}

	if err = RelayerConfig(r).RPC.SendTransaction(r.Context(), tx); err != nil {
		if strings.Contains(err.Error(), "nonce") {
			if err = RelayerConfig(r).ResetNonce(RelayerConfig(r).RPC); err != nil {
				return nil, fmt.Errorf("failed to reset nonce: %w", err)
			}

			tx, err = signTx(r, txd, receiver)
			if err != nil {
				return nil, fmt.Errorf("failed to sign new tx: %w", err)
			}

			if err := RelayerConfig(r).RPC.SendTransaction(r.Context(), tx); err != nil {
				return nil, fmt.Errorf("failed to send transaction: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to send transaction: %w", err)
		}
	}

	RelayerConfig(r).IncrementNonce()

	return tx, nil
}

func signTx(r *http.Request, txd *txData, receiver *common.Address) (tx *types.Transaction, err error) {
	tx, err = types.SignNewTx(
		RelayerConfig(r).PrivateKey,
		types.NewCancunSigner(RelayerConfig(r).ChainID),
		&types.LegacyTx{
			Nonce:    RelayerConfig(r).Nonce(),
			Gas:      txd.gas,
			GasPrice: txd.gasPrice,
			To:       receiver,
			Data:     txd.dataBytes,
		},
	)
	return tx, err
}
