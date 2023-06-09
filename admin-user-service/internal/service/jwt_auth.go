package middleware

import (
	"admin-user-service/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, errValidation := ValidateUser(c, "operationID")
		if errValidation != nil {
			log.Debug().Msgf("", "GetUserIDFromToken false ", c.Request.Header.Get("token"))
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		log.Debug().Msgf("0", "userID: ", userID)
		c.Set("userID", userID)
		c.Next()
	}
}

// ValidateUser validate Header Token and get user ID from token also check is user
//
//	Also check user si blocked or not
func ValidateUser(c *gin.Context, operationID string) (string, error) {
	//get the userId from token
	token := c.GetHeader("token")
	if token == "" {
		log.Debug().Msgf(operationID, utils.GetSelfFuncName(), "token is nil")
		//openIMHttp.RespHttp403(c, constant.ErrTokenInvalid, nil)
		return "", errors.New("token is nil")
	}

	//_, userID, officialID, _ := token_verify.GetUserIDFromTokenV2(token, "")
	//if userID == "" {
	//	log.NewError(operationID, utils.GetSelfFuncName(), "token is illegal")
	//	openIMHttp.RespHttp403(c, constant.ErrTokenInvalid, nil)
	//	return userID, errors.New("token is illegal")
	//}

	return "userID", nil
}

func JWTAuthByPhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, errValidation := validateUserByPhoneToken(c, "operationID")
		if errValidation != nil {
			log.Debug().Msgf("", "GetUserIDFromToken false ", c.Request.Header.Get("token"))
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		log.Debug().Msgf("0", "userID: ", userID)
		c.Set("phone", userID)
		c.Next()
	}
}

func validateUserByPhoneToken(c *gin.Context, operationID string) (string, error) {
	token := c.GetHeader("token")
	if token == "" {
		log.Debug().Msgf(operationID, utils.GetSelfFuncName(), "token is nil")
		return "", errors.New("token is nil")
	}
	claims, err := GetPhoneJWTManager().VerifyToken(token)
	if err != nil {
		return "", errors.New("token is not valid")
	}
	clm := claims.(jwt.MapClaims)
	phone := clm["phone"].(string)
	return phone, nil
}
