package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"user-service/internal/repository/grpcConnection"
	"user-service/pkg/utils"

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
	return
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
		otpVerifyResponse := domain.VerifyOTPResponse{}
		log.Debug().Fields(*grpcResp).Msg("VerifyOTP Response")
		reqUserAccount := userAuth.GetUserAccountByPhoneRequest{PhoneNumber: grpcResp.PhoneNumber}
		grpcRespUserAccount, err := grpcAuthClient.GetUserAccountByPhone(context.Background(), &reqUserAccount)
		if err != nil || grpcRespUserAccount.UserAccount == nil || grpcRespUserAccount.UserAccount.UserID == 0 {

			tokenByPhoneRequest := userAuth.GetTokenByPhoneRequest{PhoneNumber: grpcResp.PhoneNumber}
			grpcRespTokenByPhone, err := grpcAuthClient.GetTokenByPhone(context.Background(), &tokenByPhoneRequest)
			if err != nil {
				errMsg := err.Error() + req.String()
				log.Error().Any("grpc", errMsg)
				resp.ErrCode = 4031
				resp.ErrMsg = errMsg
				c.JSON(http.StatusBadRequest, resp)
				return
			}
			otpVerifyResponse.Token = grpcRespTokenByPhone.Token
			otpVerifyResponse.TempToken = true
		} else {
			tokenByUserIDRequest := userAuth.GetTokenByUserIDRequest{UserID: grpcRespUserAccount.UserAccount.UserID}
			grpcRespTokenByUserID, err := grpcAuthClient.GetTokenByUserID(context.Background(), &tokenByUserIDRequest)
			if err != nil {
				errMsg := err.Error() + req.String()
				log.Error().Any("grpc", errMsg)
				resp.ErrCode = 4032
				resp.ErrMsg = errMsg
				c.JSON(http.StatusBadRequest, resp)
				return
			}
			otpVerifyResponse.Token = grpcRespTokenByUserID.Token
			otpVerifyResponse.TempToken = false
			userAccount := domain.UserAccount{}
			utils.CopyFields(grpcRespUserAccount.UserAccount, &userAccount)
			otpVerifyResponse.UserAccount = &userAccount
		}

		resp := domain.CommDataResp{}

		resp.Data = otpVerifyResponse
		c.JSON(http.StatusOK, resp)
		return

	}

	resp.ErrCode = 9001
	resp.ErrMsg = "otp is invalid"
	c.JSON(http.StatusOK, resp)
	return
}

func LoginHandler(c *gin.Context) {
	params := domain.LoginRequest{}
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

	resp := domain.CommDataResp{}
	req := userAuth.VerifyUserEmailAndPasswordRequest{}
	err = utils.CopyFields(&params, &req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"errCode": 8003, "errMsg": err.Error()})
		return
	}

	grpcResp, err := grpcAuthClient.VerifyUserEmailAndPassword(context.Background(), &req)
	if err != nil {
		errMsg := err.Error() + req.String()
		log.Error().Any("grpc", errMsg)
		resp.ErrCode = 8002
		resp.ErrMsg = errMsg
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if grpcResp.UserAccount != nil && grpcResp.UserAccount.UserID != 0 {
		otpVerifyResponse := domain.VerifyOTPResponse{}
		tokenByUserIDRequest := userAuth.GetTokenByUserIDRequest{UserID: grpcResp.UserAccount.UserID}
		grpcRespTokenByUserID, err := grpcAuthClient.GetTokenByUserID(context.Background(), &tokenByUserIDRequest)
		if err != nil {
			errMsg := err.Error() + req.String()
			log.Error().Any("grpc", errMsg)
			resp.ErrCode = 4032
			resp.ErrMsg = errMsg
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		otpVerifyResponse.Token = grpcRespTokenByUserID.Token
		otpVerifyResponse.TempToken = false
		userAccount := domain.UserAccount{}
		utils.CopyFields(grpcResp.UserAccount, &userAccount)
		otpVerifyResponse.UserAccount = &userAccount

		resp := domain.CommDataResp{}
		resp.Data = otpVerifyResponse
		c.JSON(http.StatusOK, resp)
		return

	} else {
		resp.ErrCode = 8005
		resp.ErrMsg = "user not found"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

}

func CreateAccountHandler(c *gin.Context) {
	params := domain.CreateAccountRequest{}
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

	req := userAuth.CreateUserByEmailRequest{}
	req.UserAccount = &userAuth.UserAccount{}
	err = utils.CopyFields(&params, req.UserAccount)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"errCode": 8003, "errMsg": err.Error()})
		return
	}
	phone, ok := c.Get("phone")
	if ok {
		req.UserAccount.PhoneNumber = phone.(string)
	}

	resp := domain.CommDataResp{}

	grpcResp, err := grpcAuthClient.CreateUserByEmail(context.Background(), &req)
	if err != nil {
		errMsg := err.Error() + req.String()
		log.Error().Any("grpc", errMsg)
		resp.ErrCode = 8002
		resp.ErrMsg = errMsg
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	log.Error().Msg("Resp Value " + strconv.Itoa(int(grpcResp.UserAccount.UserID)) + " - " + grpcResp.UserAccount.Name)
	if grpcResp.UserAccount != nil && grpcResp.UserAccount.UserID != 0 {
		otpVerifyResponse := domain.VerifyOTPResponse{}
		tokenByUserIDRequest := userAuth.GetTokenByUserIDRequest{UserID: grpcResp.UserAccount.UserID}
		grpcRespTokenByUserID, err := grpcAuthClient.GetTokenByUserID(context.Background(), &tokenByUserIDRequest)
		if err != nil {
			errMsg := err.Error() + req.String()
			log.Error().Any("grpc", errMsg)
			resp.ErrCode = 4032
			resp.ErrMsg = errMsg
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		otpVerifyResponse.Token = grpcRespTokenByUserID.Token
		otpVerifyResponse.TempToken = false
		userAccount := domain.UserAccount{}
		utils.CopyFields(grpcResp.UserAccount, &userAccount)
		otpVerifyResponse.UserAccount = &userAccount

		resp := domain.CommDataResp{}
		resp.Data = otpVerifyResponse
		c.JSON(http.StatusOK, resp)
		return

	} else {
		resp.ErrCode = 8005
		resp.ErrMsg = "user not created"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
}
