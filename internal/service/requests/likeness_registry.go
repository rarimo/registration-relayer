package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/registration-relayer/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewLikenessRequest(r *http.Request) (req resources.LikenessRequest, err error) {
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, errors.Wrap(err, "failed to unmarshal")
	}

	return req, validation.Errors{
		"data/tx_data": validation.Validate(req.Data.TxData, validation.Required, validation.Match(calldataRegexp)),
	}.Filter()
}
