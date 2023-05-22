package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"user-service/pkg/grpc-etcdv3/getcdv3"
	"user-service/proto/userAuth"
)

func SignupHandler(c *gin.Context) {

}

func LoginHandler(c *gin.Context) {

	req := userAuth.GetOTPRequest{}
	// Handle login logic here strings.Join(config.Config.Etcd.EtcdAddr, ",")
	etcdConn := getcdv3.GetConn("fishOn", strings.Join([]string{"127.0.0.1:2379"}, ","), "UserAuth", "req.OperationID")
	if etcdConn == nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "not connected"})
		return
	}
	client := userAuth.NewUserAuthClient(etcdConn)
	reply, err := client.GetOTP(context.Background(), &req)
	if err != nil {
		errMsg := " UserToken failed " + err.Error() + req.String()
		//log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": reply.PinToken,
	})
}

func CreateAccountHandler(c *gin.Context) {
	// Handle account creation logic here
}
