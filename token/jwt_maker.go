package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Jwt_Maker struct{
	secretKey string
}
type Claims struct{
	Username string
	jwt.RegisteredClaims
}
const minSecretKeySize = 32
func NewJwt_Maker(secretKey string)(*Jwt_Maker, error){
	if len(secretKey) < minSecretKeySize{
		return nil, fmt.Errorf("invalid key size: must beat least %d characters",minSecretKeySize)
	}
	return &Jwt_Maker{
		secretKey: secretKey,
	}, nil
}

var (
	ErrInvalidToken = errors.New("Invalid Token")
	ErrExpiredToken = errors.New("Token Expired")
)

var jwtkey = []byte("secret")
func (j *Jwt_Maker)CreateToken(username string, duration time.Duration)(string,*Payload, error){
	payload, err :=NewPayload(username,duration)
	if err != nil{
		return "",payload, nil
	}
	// claims := &Claims{
	// 	Username: payload.Username,
	// 	RegisteredClaims:jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
	// 		IssuedAt: jwt.NewNumericDate(payload.IssuedAt),
	// 		ID: payload.TokenId.String(),
	// 		NotBefore: jwt.NewNumericDate(payload.IssuedAt),
	// 	},
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil{
		return "",payload, err
	}
	return tokenString,payload, nil
}

func(j *Jwt_Maker)VerifyToken(signed string)(*Payload, error){
	keyFunc := func(token *jwt.Token)(interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil,ErrInvalidToken
		} 
		return jwtkey, nil
	}
	jwtToken, err := jwt.ParseWithClaims(signed,&Payload{},keyFunc)
	if err != nil{
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner,ErrExpiredToken){
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok{
		return nil, ErrInvalidToken
	}
	return payload, nil
}