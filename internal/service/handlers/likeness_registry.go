package handlers

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/registration-relayer/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func LikenessRegistry(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewLikenessRequest(r)
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
	log.Debug("likeness registry request")

	if RelayerConfig(r).LikenessRegistryAddress == nil {
		Log(r).Error("likeness registry address is not set in config")
		ape.RenderErr(w, problems.InternalError())
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

	err = confGas(r, &txd, RelayerConfig(r).LikenessRegistryAddress)
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

	tx, err := sendTx(r, &txd, RelayerConfig(r).LikenessRegistryAddress, req.Data.NoSend)
	if err != nil {
		Log(r).WithError(err).Error("failed to send tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newTxResponse(tx))
}
