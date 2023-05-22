package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"user-service/proto/userAuth"
)

func SignupHandler(c *gin.Context) {

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
	etcdConn, err := grpc.DialContext(context.Background(), "user-auth-grpc:10021", grpc.WithInsecure())
	if err != nil {
		log.Println("Failed to dial gRPC server: %v", err)
	}
	defer etcdConn.Close()
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
