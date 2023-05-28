package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"user-service/internal/repository/grpcConnection"

	"github.com/rs/zerolog/log"

	"net/http"
	"user-service/internal/domain"
	"user-service/proto/userAuth"
)

// GetOTPHandler get otp for signUp user
// @Summary get otp
// @Description Get all users from the database.
// @Tags OTP
// @Produce json
// @Success 200 {string} domain.CommDataResp{}
// @Router /get_otp [post]
func GetOTPHandler(c *gin.Context) {
	params := domain.OTPRequest{}
	if err := c.BindJSON(&params); err != nil {
		errMsg := " BindJSON failed " + err.Error()
		log.Error().Any("0", errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}

	grpcConnection := grpcConnection.GRPCConnectionsImpl{}
	grpcAuthClient, err := grpcConnection.GetUserAuthGRPcConnection(context.Background())
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"errCode": 8001, "errMsg": err.Error()})
		return
	}

	req := userAuth.GetOTPRequest{}
	resp := domain.CommDataResp{}
	req.PhoneNumber = params.CountryCode + params.PhoneNumber
	grpcResp, err := grpcAuthClient.GetOTP(context.Background(), &req)
	if err != nil {
		errMsg := err.Error() + req.String()
		log.Error().Any("grpc", errMsg)
		resp.ErrCode = 8002
		resp.ErrMsg = errMsg
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	getOTPResponse := domain.GetOTPResponse{}
	getOTPResponse.PinToken = grpcResp.PinToken
	getOTPResponse.Status = grpcResp.Status

	resp.Data = getOTPResponse
	c.JSON(http.StatusOK, resp)

}

// VerifyOTPHandler verify otp for signUp user
// @Summary verify otp
// @Description Get all users from the database.
// @Tags verify OTP
// @Produce json
// @Success 200 {string} domain.CommDataResp{}
// @Router /get_otp [post]
func VerifyOTPHandler(c *gin.Context) {
	params := domain.OTPVerifyRequest{}
	if err := c.BindJSON(&params); err != nil {
		errMsg := " BindJSON failed " + err.Error()
		log.Error().Any("0", errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}

	grpcConnection := grpcConnection.GRPCConnectionsImpl{}
	grpcAuthClient, err := grpcConnection.GetUserAuthGRPcConnection(context.Background())
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"errCode": 8001, "errMsg": err.Error()})
		return
	}

	req := userAuth.VerifyOTPRequest{}
	resp := domain.CommDataResp{}
	req.PinToken = params.PinToken
	req.Otp = params.Otp

	grpcResp, err := grpcAuthClient.VerifyOTP(context.Background(), &req)
	if err != nil {
		errMsg := err.Error() + req.String()
		log.Error().Any("grpc", errMsg)
		resp.ErrCode = 8002
		resp.ErrMsg = errMsg
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if grpcResp.Status == 1 {
		reqUserAccount := userAuth.GetUserAccountByPhoneRequest{PhoneNumber: grpcResp.PhoneNumber}
		grpcRespUserAccount, err := grpcAuthClient.GetUserAccountByPhone(context.Background(), &reqUserAccount)
		if err != nil || grpcRespUserAccount.UserID == 0 {
			//todo
			// create a temp token so user can create user account with it.
			// create token by user phone number
		} else {
			//todo
			// create a user token and let user so user can move to Home screen
			// create user by user ID
		}

	}

	resp.ErrCode = 9001
	resp.ErrMsg = "otp is invalid"
	c.JSON(http.StatusOK, resp)

}

func LoginHandler(c *gin.Context) {

}

func CreateAccountHandler(c *gin.Context) {
}
