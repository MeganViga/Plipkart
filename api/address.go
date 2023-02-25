package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/gin-gonic/gin"
)

type addAddressRequestParams struct{
	UserID       int64     `json:"user_id"`
	AddressLine1 string    `json:"address_line1"`
	AddressLine2 string    `json:"address_line2"`
	City         string    `json:"city"`
	PostalCode   string    `json:"postal_code"`
	Country      string    `json:"country"`
	PhoneNumber  string    `json:"phone_number"`
}
func (s *Server)addAddress(ctx *gin.Context){
	var request addAddressRequestParams
	if err := ctx.ShouldBindJSON(&request); err != nil{
		ctx.JSON(http.StatusBadRequest,s.ShowError(err))
		return
	}
	arg := db.CreateUserAddressParams{
		UserID: request.UserID,
		AddressLine1: request.AddressLine1,
		AddressLine2: request.AddressLine2,
		City: request.City,
		PostalCode: request.PostalCode,
		Country: request.Country,
		PhoneNumber: request.PhoneNumber,
		CreatedAt: time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
	}

	address, err := s.store.CreateUserAddress(ctx,arg)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
		return
	}
	ctx.JSON(http.StatusOK,address)
}
type deleteAddressRequestParams struct{
	Id int64 `uri:"id" binding:"required"`
}
func (s *Server)deleteAddress(ctx *gin.Context){
	var request deleteAddressRequestParams
	if err := ctx.ShouldBindUri(&request); err != nil{
		ctx.JSON(http.StatusBadRequest,s.ShowError(err))
		return
	}
	fmt.Println(request)
	address, err := s.store.DeleteUserAddress(ctx,request.Id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
		return
	}
	ctx.JSON(http.StatusOK,address)
}
type listAddressRequestParams struct{
	User_id int64 `uri:"id" binding:"required"`
}
func (s *Server)listAddress(ctx *gin.Context){
	var request listAddressRequestParams
	fmt.Println(request)
	if err := ctx.ShouldBindUri(&request); err != nil{
		ctx.JSON(http.StatusBadRequest,s.ShowError(err))
		return
	}

	addresses, err := s.store.ListUserAddresses(ctx,request.User_id)
	fmt.Println(addresses)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
		return
	}
	ctx.JSON(http.StatusOK,addresses)
}