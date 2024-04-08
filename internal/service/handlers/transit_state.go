package handlers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/registration-relayer/internal/service/requests"
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

	var txd txData
	txd.dataBytes, err = hexutil.Decode(req.Data.TxData)
	if err != nil {
		Log(r).WithError(err).Error("failed to decode data")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	RelayerConfig(r).LockNonce()
	defer RelayerConfig(r).UnlockNonce()

	err = confGas(r, &txd, RelayerConfig(r).LightweightStateAddress)
	if err != nil {
		Log(r).WithError(err).Error("failed to configure gas and gasPrice")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := sendTx(r, &txd, RelayerConfig(r).LightweightStateAddress)
	if err != nil {
		Log(r).WithError(err).Error("failed to send tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	RelayerConfig(r).IncrementNonce()

	ape.Render(w, newTxResponse(tx))
}
