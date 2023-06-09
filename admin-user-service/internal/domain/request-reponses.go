package domain

type CommResp struct {
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

type CommDataResp struct {
	CommResp
	Data interface{} `json:"data"`
}

type GetOTPResponse struct {
	PinToken string `json:"pinToken"`
	Status   int32  `json:"status"`
}

type VerifyOTPResponse struct {
	Token       string       `json:"token"`
	TempToken   bool         `json:"tempToken"`
	UserAccount *UserAccount `json:"userAccount,omitempty"`
}
