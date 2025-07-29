package common

var (
	MsgCreateFinalityProvider      = "/babylon.btcstaking.v1.MsgCreateFinalityProvider"
	MsgCreateBTCDelegation         = "/babylon.btcstaking.v1.MsgCreateBTCDelegation"
	MsgCommitPubRandList           = "/babylon.finality.v1.MsgCommitPubRandList"
	MsgBTCUndelegate               = "/babylon.btcstaking.v1.MsgBTCUndelegate"
	MsgSelectiveSlashingEvidence   = "/babylon.btcstaking.v1.MsgSelectiveSlashingEvidence"
	MsgSubmitFinalitySignatureType = "/MsgSubmitFinalitySignature"
)

const (
	BabylonSignType = iota
	SymbioticSignType
)

type SymbioticStakeRequest struct {
	Query string `json:"query"`
}

type SymbioticStakeResponse struct {
	Data struct {
		VaultUpdates []struct {
			VaultTotalActiveStaked string `json:"vaultTotalActiveStaked"`
		} `json:"vaultUpdates"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}
