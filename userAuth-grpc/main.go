package main

import (
	"flag"
	"fmt"
	"userAuth-grpc/internal"
)

func main() {
	//lis, err := net.Listen("tcp", ":10021")
	//if err != nil {
	//	log.Fatalf("Failed to listen: %v", err)
	//}
	//
	//grpcServer := grpc.NewServer()
	//userAuthServer := internal.UserAuthServer{}
	//userAuth.RegisterUserAuthServer(grpcServer, &userAuthServer)
	//
	//if err := grpcServer.Serve(lis); err != nil {
	//}

	//log.NewPrivateLog(constant.OpenImAuthLog)
	//defaultPorts := config.Config.RpcPort.OpenImAuthPort
	//rpcPort := flag.Int("port", defaultPorts[0], "RpcToken default listen port 10800")
	rpcPort := flag.Int("port", 10021, "RpcToken default listen port 10800")
	flag.Parse()
	fmt.Println("start auth rpc server, port: ", *rpcPort)
	userAuthServer := internal.NewRpcUserAuthServer(*rpcPort)
	userAuthServer.Run()
}
