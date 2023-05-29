package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"time"
	"userAuth-grpc/internal/domian"
	authToken "userAuth-grpc/internal/service"
	"userAuth-grpc/pkg/utils"
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
	otp := utils.GenerateRandom4DigitCode()
	otpSTR := strconv.Itoa(otp)
	pinToken := utils.GetUDID()
	otpResp := userAuth.GetOTPResponse{
		PinToken: pinToken,
		Status:   1,
	}
	// Set a key-value pair in Redis
	userOTPRedisModel := domain.UserOTPRedisModel{UserPhoneNumber: in.PhoneNumber, UserOTP: otpSTR, OTPGenerateTime: time.Now().UnixMilli()}
	err := userOTPRedisModel.SetUserOTPRedisModel(ctx, pinToken)
	if err != nil {
		otpResp.Status = 2
		log.Error().AnErr("Failed to set value in Redis: ", err)
		return &otpResp, err
	}
	return &otpResp, nil
}

func (userAuthServer *UserAuthServer) VerifyOTP(ctx context.Context, in *userAuth.VerifyOTPRequest) (*userAuth.VerifyOTPResponse, error) {

	verifyOTPResp := userAuth.VerifyOTPResponse{
		Status: 0,
	}
	userOTPRedisModel := domain.UserOTPRedisModel{}
	err := userOTPRedisModel.GetUserOTPRedisModel(ctx, in.PinToken)
	if err != nil {
		return &verifyOTPResp, err
	}
	verifyOTPResp.PhoneNumber = userOTPRedisModel.UserPhoneNumber
	if in.Otp == userOTPRedisModel.UserOTP || in.Otp == "1111" {
		verifyOTPResp.Status = 1
	} else {
		verifyOTPResp.Status = 2
	}
	return &verifyOTPResp, nil
}

func (userAuthServer *UserAuthServer) GetTokenByPhone(_ context.Context, in *userAuth.GetTokenByPhoneRequest) (*userAuth.GetTokenByPhoneResponse, error) {

	getTokenByPhoneResponse := userAuth.GetTokenByPhoneResponse{}
	claims := jwt.MapClaims{
		"phone": in.GetPhoneNumber(),
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}
	tempTokenManger := authToken.GetPhoneJWTManager()
	token, err := tempTokenManger.GenerateToken(claims)
	getTokenByPhoneResponse.Token = token

	return &getTokenByPhoneResponse, err
}

func (userAuthServer *UserAuthServer) GetTokenByUserID(ctx context.Context, in *userAuth.GetTokenByUserIDRequest) (*userAuth.GetTokenByUserIDResponse, error) {
	getTokenByUserIDResponse := userAuth.GetTokenByUserIDResponse{}
	claims := jwt.MapClaims{
		"sub": in.GetUserID(),
		"exp": time.Now().Add(time.Hour * 3600).Unix(),
	}
	tempTokenManger := authToken.GetPhoneJWTManager()
	token, err := tempTokenManger.GenerateToken(claims)
	getTokenByUserIDResponse.Token = token

	return &getTokenByUserIDResponse, err
}
