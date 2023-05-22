package internal

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"strings"
	"userAuth-grpc/pkg/grpc-etcdv3/getcdv3"
	"userAuth-grpc/proto/userAuth"
)

type UserAuthServer struct {
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewRpcUserAuthServer(port int) *UserAuthServer {
	return &UserAuthServer{
		rpcPort:         port,
		rpcRegisterName: "UserAuth",                 //config.Config.RpcRegisterName.OpenImAuthName,
		etcdSchema:      "fishOn",                   //config.Config.Etcd.EtcdSchema,
		etcdAddr:        []string{"127.0.0.1:2379"}, //config.Config.Etcd.EtcdAddr,
	}
}

func (userAuthServer *UserAuthServer) Run() {
	//log.NewPrivateLog(constant.OpenImAuthLog)
	//operationID := utils.OperationIDGenerator()
	//log.NewInfo(operationID, "rpc auth start...")

	listenIP := ""
	//if config.Config.ListenIP == "" {
	listenIP = "0.0.0.0"
	//} else {
	//	listenIP = config.Config.ListenIP
	//}
	address := listenIP + ":" + strconv.Itoa(userAuthServer.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("listening err:" + err.Error() + userAuthServer.rpcRegisterName)
	}
	//log.NewInfo(operationID, "listen network success, ", address, listener)
	//grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	//service registers with etcd
	userAuth.RegisterUserAuthServer(srv, userAuthServer)
	rpcRegisterIP := "0.0.0.0"
	//if config.Config.RpcRegisterIP == "" {
	//rpcRegisterIP, err = utils.GetLocalIP()
	//	if err != nil {
	//		log.Error("", "GetLocalIP failed ", err.Error())
	//	}
	//}
	//log.NewInfo("", "rpcRegisterIP", rpcRegisterIP)
	err = getcdv3.RegisterEtcd(userAuthServer.etcdSchema, strings.Join(userAuthServer.etcdAddr, ","), rpcRegisterIP, userAuthServer.rpcPort, userAuthServer.rpcRegisterName, 10)
	if err != nil {
		log.Println(err.Error())
		//log.NewError(operationID, "RegisterEtcd failed ", err.Error(),
		//	rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), rpcRegisterIP, rpc.rpcPort, rpc.rpcRegisterName)
		return
	}
	//log.NewInfo(operationID, "RegisterAuthServer ok ", rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), rpcRegisterIP, rpc.rpcPort, rpc.rpcRegisterName)
	err = srv.Serve(listener)
	if err != nil {
		log.Println(err.Error())
		//log.NewError(operationID, "Serve failed ", err.Error())
		return
	}
	//log.NewInfo(operationID, "rpc auth ok")
}

func (userAuthServer *UserAuthServer) GetOTP(ctx context.Context, in *userAuth.GetOTPRequest) (*userAuth.GetOTPResponse, error) {
	otpResp := userAuth.GetOTPResponse{
		PinToken: "TestingToken",
		Status:   1,
	}
	return &otpResp, nil
}
