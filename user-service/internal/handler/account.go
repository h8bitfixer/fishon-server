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

	resp.Data = make(map[string]interface{})
	resp.Data["pinToken"] = grpcResp.PinToken
	resp.Data["status"] = grpcResp.Status
	c.JSON(http.StatusOK, resp)

}

func LoginHandler(c *gin.Context) {

	req := userAuth.GetOTPRequest{}
	// Handle login logic here strings.Join(config.Config.Etcd.EtcdAddr, ",")
	//_, err := getcdv3.ResolveEtcd([]string{"127.0.0.1:2379"}, "user-auth-grpc")
	////"fishOn", strings.Join([]string{"127.0.0.1:2379"}, ","), "user-auth-grpc", "req.OperationID")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "not connected"})
	//	return
	//}
	// Create a gRPC connection to one of the resolved endpoints
	grpcConnection := grpcConnection.GRPCConnectionsImpl{}
	grpcAuthClient, err := grpcConnection.GetUserAuthGRPcConnection(context.Background())
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"errCode": 8001, "errMsg": err.Error()})
		return
	}
	reply, err := grpcAuthClient.GetOTP(context.Background(), &req)
	if err != nil {
		errMsg := " UserToken failed " + err.Error() + req.String()
		//log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 8002, "errMsg": errMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": reply.PinToken,
	})
}

func CreateAccountHandler(c *gin.Context) {
	// Handle account creation logic here
}
