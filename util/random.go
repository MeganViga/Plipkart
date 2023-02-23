package util

import (
	"math/rand"
	"time"
)


func RandInt(min, max int)int{
	rand.Seed(time.Now().Unix())
	r := rand.Intn(max -min) + min
	return r
}

func RandString(length int)string{
	alphabets := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var rstring []byte
	for i :=0 ;i <  length;i++{
		rstring = append(rstring, alphabets[RandInt(0,length)])
	}
	return string(rstring)
}