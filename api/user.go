package api

import (
	//"fmt"
	"net/http"
	"time"

	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/MeganViga/Plipkart/util"
	"github.com/gin-gonic/gin"
)


type userSignupRequestParams struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Firstname string `json:"first_name" binding:"required"`
	Lastname string `json:"last_name" binding:"required"`
	Role string `json:"role" binding:"required"`
	Phonenumber string `json:"phone_number" binding:"required"`


}
func (s *Server)userSignup(ctx *gin.Context){
	var request userSignupRequestParams
	if err := ctx.ShouldBindJSON(&request); err != nil{
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	password_hash , err := util.HashPassword(request.Password)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	arg := db.CreateUserParams{
		Username: request.Username,
		Password: password_hash,
		FirstName: request.Firstname,
		LastName: request.Lastname,
		Role: request.Role,
		PhoneNumber: request.Phonenumber,
		CreatedAt: time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),

	}
	user, err := s.store.CreateUser(ctx,arg)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,user)
}

type userLoginRequestParams struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (s *Server)userLogin(ctx *gin.Context){
	var request userLoginRequestParams
	if err := ctx.ShouldBindJSON(&request); err != nil{
		ctx.JSON(http.StatusBadRequest, s.ShowError(err))
		return
	}

	user, err := s.store.GetUserByName(ctx,request.Username)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
		return
	}

	if err := util.VerifyPassword(request.Password,user.Password); err != nil{
		print("Error", err)
		ctx.JSON(http.StatusInternalServerError, s.ShowError(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}