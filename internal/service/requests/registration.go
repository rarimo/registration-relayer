package requests

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RegistrationRequestData struct {
	TxData string `json:"tx_data"`
}

type RegistrationRequest struct {
	Data RegistrationRequestData `json:"data"`
}

func NewRegistrationRequest(r *http.Request) (RegistrationRequest, error) {
	var request RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}
