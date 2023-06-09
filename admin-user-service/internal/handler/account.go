package handler

import (
	"admin-user-service/internal/repository/grpcConnection"
	"admin-user-service/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"admin-user-service/internal/domain"
	"admin-user-service/proto/userAuth"
	"net/http"
)

// LoginHandler get otp for signUp user
// @Summary get otp
// @Description Get all users from the database.
// @Tags OTP
// @Produce json
// @Success 200 {string} domain.CommDataResp{}
// @Router /get_otp [post]
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
