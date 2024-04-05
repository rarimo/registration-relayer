package handlers

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/registration-relayer/internal/service/requests"
	"github.com/rarimo/registration-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func TransitState(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistrationRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	dataBytes, err := hexutil.Decode(req.Data.TxData)
	if err != nil {
		Log(r).WithError(err).Error("failed to decode data")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	gasPrice, err := RelayerConfig(r).RPC.SuggestGasPrice(r.Context())
	if err != nil {
		Log(r).WithError(err).Error("failed to suggest gas price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	RelayerConfig(r).LockNonce()
	defer RelayerConfig(r).UnlockNonce()

	gas, err := RelayerConfig(r).RPC.EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(RelayerConfig(r).PrivateKey.PublicKey),
		To:       &RelayerConfig(r).LightweightStateAddress,
		GasPrice: gasPrice,
		Data:     dataBytes,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to estimate gas")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := types.SignNewTx(
		RelayerConfig(r).PrivateKey,
		types.NewCancunSigner(RelayerConfig(r).ChainID),
		&types.LegacyTx{
			Nonce:    RelayerConfig(r).Nonce(),
			Gas:      gas,
			GasPrice: gasPrice,
			To:       &RelayerConfig(r).LightweightStateAddress,
			Data:     dataBytes,
		},
	)
	if err != nil {
		Log(r).WithError(err).Error("failed to sign new tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if err := RelayerConfig(r).RPC.SendTransaction(r.Context(), tx); err != nil {
		if strings.Contains(err.Error(), "nonce") {
			if err := RelayerConfig(r).ResetNonce(RelayerConfig(r).RPC); err != nil {
				Log(r).WithError(err).Error("failed to reset nonce")
				ape.RenderErr(w, problems.InternalError())
				return
			}

			tx, err = types.SignNewTx(
				RelayerConfig(r).PrivateKey,
				types.NewCancunSigner(RelayerConfig(r).ChainID),
				&types.LegacyTx{
					Nonce:    RelayerConfig(r).Nonce(),
					Gas:      gas,
					GasPrice: gasPrice,
					To:       &RelayerConfig(r).LightweightStateAddress,
					Data:     dataBytes,
				},
			)
			if err != nil {
				Log(r).WithError(err).Error("failed to sign new tx")
				ape.RenderErr(w, problems.InternalError())
				return
			}

			if err := RelayerConfig(r).RPC.SendTransaction(r.Context(), tx); err != nil {
				Log(r).WithError(err).Error("failed to send transaction")
				ape.RenderErr(w, problems.InternalError())
				return
			}
		} else {
			Log(r).WithError(err).Error("failed to send transaction")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	RelayerConfig(r).IncrementNonce()

	ape.Render(w, resources.Tx{
		Key: resources.Key{
			ID:   tx.Hash().String(),
			Type: resources.TXS,
		},
		Attributes: resources.TxAttributes{
			TxHash: tx.Hash().String(),
		},
	})
}
