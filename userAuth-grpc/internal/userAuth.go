package internal

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
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
		rpcRegisterName: "user-auth-grpc",      //config.Config.RpcRegisterName.OpenImAuthName,
		etcdSchema:      "fishOn",              //config.Config.Etcd.EtcdSchema,
		etcdAddr:        []string{"etcd:2379"}, //config.Config.Etcd.EtcdAddr,
	}
}

func (userAuthServer *UserAuthServer) Run() {
	listenIP := ""
	address := listenIP + ":" + strconv.Itoa(userAuthServer.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("listening err:" + err.Error() + userAuthServer.rpcRegisterName)
	}

	srv := grpc.NewServer()
	defer srv.GracefulStop()
	//rpcRegisterIP := userAuthServer.rpcRegisterName + ":" + strconv.Itoa(userAuthServer.rpcPort)
	//service registers with etcd
	userAuth.RegisterUserAuthServer(srv, userAuthServer)

	//err = getcdv3.RegisterEtcd(userAuthServer.etcdAddr, userAuthServer.rpcRegisterName, rpcRegisterIP, 10)
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	err = srv.Serve(listener)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (userAuthServer *UserAuthServer) GetOTP(ctx context.Context, in *userAuth.GetOTPRequest) (*userAuth.GetOTPResponse, error) {
	otpResp := userAuth.GetOTPResponse{
		PinToken: "TestingToken",
		Status:   1,
	}
	return &otpResp, nil
}
