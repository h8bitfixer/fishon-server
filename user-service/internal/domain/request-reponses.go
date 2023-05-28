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
