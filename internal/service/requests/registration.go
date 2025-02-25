package requests

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/registration-relayer/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	calldataRegexp = regexp.MustCompile("^0x[0-9a-fA-F]+$")
	addressRegexp  = regexp.MustCompile("^0x[0-9a-fA-F]{40}")
)

func NewRegistrationRequest(r *http.Request) (req resources.RegistrationRequest, err error) {
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, errors.Wrap(err, "failed to unmarshal")
	}

	if req.Data.Destination != nil {
		*req.Data.Destination = strings.ToLower(*req.Data.Destination)
	}

	return req, validation.Errors{
		"data/tx_data":     validation.Validate(req.Data.TxData, validation.Required, validation.Match(calldataRegexp)),
		"data/destination": validation.Validate(req.Data.Destination, validation.When(req.Data.Destination != nil, validation.Required, validation.Match(addressRegexp))),
	}.Filter()
}
