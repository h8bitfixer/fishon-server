package main

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog/log"
	"userAuth-grpc/internal"
	"userAuth-grpc/pkg/db"
)

func main() {
	err := db.InitializeMySQL()
	if err != nil {
		log.Error().AnErr("Failed to initialize MySQL: %v\n", err)
	}
	err = db.InitializeRedis()
	if err != nil {
		log.Error().AnErr("Failed to initialize Redis: %v\n", err)
	}

	rpcPort := flag.Int("port", 10021, "RpcToken default listen port 10800")
	flag.Parse()
	fmt.Println("start auth rpc server, port: ", *rpcPort)
	userAuthServer := internal.NewRpcUserAuthServer(*rpcPort)
	userAuthServer.Run()
}
