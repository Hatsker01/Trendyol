package auth

import (
	"fmt"
	"time"

	"github.com/Trendyol/Api/pkg/logger"
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtHandler struct {
	Sub       string
	Iss       string
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	Token     string
	SigninKey string
	Log       logger.Logger
}

//Generate auth...
// func (JwtHandler *JwtHandler) GenerateJwt() (accesstoken, refreshtoken string, err error) {
// 	var (
// 		accessToken, refreshToken *jwt.Token
// 		claims                    jwt.MapClaims
// 	)
// 	accessToken = jwt.New(jwt.SigningMethodHS256)
// 	refreshToken = jwt.New(jwt.SigningMethodHS256)

// 	claims = accessToken.Claims.(jwt.MapClaims)
// 	claims["sub"] = JwtHandler.Sub
// 	claims["iss"] = JwtHandler.Iss
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
// 	claims["iat"] = time.Now().Unix()
// 	claims["role"] = JwtHandler.Role
// 	claims["aud"] = JwtHandler.Aud

// 	access, err := accessToken.SignedString([]byte(JwtHandler.SigninKey))
// 	if err != nil {
// 		JwtHandler.Log.Error("error generating access token", logger.Error(err))
// 		return
// 	}

// 	refresh, err := refreshToken.SignedString([]byte(JwtHandler.SigninKey))
// 	if err != nil {
// 		JwtHandler.Log.Error("error creating refresh token", logger.Error(err))
// 		return

// 	}

// 	return access, refresh, nil
// }
func (jwtHandler *JwtHandler) GenerateAuthJWT() (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
	)

	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	claims["iss"] = jwtHandler.Iss
	claims["sub"] = jwtHandler.Sub
	claims["exp"] = time.Now().Add(time.Hour * 500).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHandler.Role
	claims["aud"] = jwtHandler.Aud
	access, err = accessToken.SignedString([]byte(jwtHandler.SigninKey))
	if err != nil {
		jwtHandler.Log.Error("error generating access token", logger.Error(err))
		return
	}

	refresh, err = refreshToken.SignedString([]byte(jwtHandler.SigninKey))
	if err != nil {
		jwtHandler.Log.Error("error generating refresh token", logger.Error(err))
		return
	}

	return
}

func (JwtHandler *JwtHandler) ExtractClaims() (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	token, err = jwt.Parse(JwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtHandler.SigninKey), nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		JwtHandler.Log.Error("invalid jwt token")
		return nil, err
	}

	return claims, nil
}
func ExtractClaim(tokenStr string, signinigKey []byte) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	fmt.Println("senma", tokenStr)
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signinigKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}
	return claims, nil
}
