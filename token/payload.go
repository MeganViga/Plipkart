package token

import (
	"time"

	"github.com/google/uuid"
)
type Payload struct{
	TokenId uuid.UUID
	Username string
	IssuedAt time.Time
	ExpiredAt time.Time
}

func NewPayload(username string,duration time.Duration)(*Payload, error){
	token, err := uuid.NewRandom()
	if err != nil{
		return nil, err
	}
	return &Payload{
		TokenId: token,
		Username: username,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (p Payload)Valid()error{
	if time.Now().After(p.ExpiredAt){
		return ErrExpiredToken
	}
	return nil
}