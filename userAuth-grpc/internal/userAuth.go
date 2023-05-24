package internal

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"userAuth-grpc/pkg/db"
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
		log.Error().Err(err)
		return
	}
}

func (userAuthServer *UserAuthServer) GetOTP(ctx context.Context, in *userAuth.GetOTPRequest) (*userAuth.GetOTPResponse, error) {
	otpResp := userAuth.GetOTPResponse{
		PinToken: in.GetPhoneNumber(),
		Status:   1,
	}
	// Set a key-value pair in Redis
	err := db.GetRedisDB().Set(ctx, "phoneNumber", in.PhoneNumber, 0).Err()
	if err != nil {
		otpResp.Status = 2
		log.Error().AnErr("Failed to set value in Redis: %v\n", err)
		return &otpResp, err
	}

	return &otpResp, nil
}
