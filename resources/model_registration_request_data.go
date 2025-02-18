/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RegistrationRequestData struct {
	// Address of registration smart contract
	Destination *string `json:"destination,omitempty"`
	// Flag indicates whether transaction should be sent on-chain
	NoSend bool   `json:"no_send"`
	TxData string `json:"tx_data"`
}
