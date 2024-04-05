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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/registration-relayer/internal/service/requests"
	"github.com/rarimo/registration-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
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

	var txd txData
	txd.dataBytes, err = hexutil.Decode(req.Data.TxData)
	if err != nil {
		Log(r).WithError(err).Error("failed to decode data")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	RelayerConfig(r).LockNonce()
	defer RelayerConfig(r).UnlockNonce()

	err = confGas(r, &txd, &RelayerConfig(r).RegistrationAddress)
	if err != nil {
		Log(r).WithError(err).Error("failed to configure gas and gasPrice")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := sendTx(r, &txd, &RelayerConfig(r).RegistrationAddress)
	if err != nil {
		Log(r).WithError(err).Error("failed to send tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	RelayerConfig(r).IncrementNonce()

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

	return nil
}

func sendTx(r *http.Request, txd *txData, receiver *common.Address) (tx *types.Transaction, err error) {
	tx, err = signTx(r, txd, receiver)
	if err != nil {
		return nil, fmt.Errorf("failed to sign new tx: %w", err)
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
