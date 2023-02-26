package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mockdb "github.com/MeganViga/Plipkart/db/mock"
	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/MeganViga/Plipkart/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)
type eqCreateUserParamsMatcher struct{
	arg db.CreateUserParams
	password string
}
func (e eqCreateUserParamsMatcher)Matches(x interface{})bool{
	arg, ok := x.(db.CreateUserParams)
	if !ok{
		return false
	}
	err := util.VerifyPassword(e.password,arg.Password)
	if err != nil{
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg,arg)
}

func (e eqCreateUserParamsMatcher)String()string{
	return fmt.Sprintf("matches arg %v and password %v",e.arg,e.password)
}

func EqCreateUserParams(arg db.CreateUserParams,password string)gomock.Matcher{
	return eqCreateUserParamsMatcher{arg,password}
}
func TestCreateUserAPI(t *testing.T){
	user, password := randomUser(t)
	testCases := []struct{
		name string
		body gin.H
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:"OK",
			body: gin.H{
				"username":user.Username,
				"password":password,
				"first_name":user.FirstName,
				"last_name":user.LastName,
				"role":user.Role,
				"phone_number":user.PhoneNumber,
			},
			buildStubs: func(store *mockdb.MockStore){
				arg := db.CreateUserParams{
					Username: user.Username,
					FirstName: user.FirstName,
					LastName: user.LastName,
					Role: user.Role,
					PhoneNumber: user.PhoneNumber,
				}
				store.EXPECT().CreateUser(gomock.Any(),EqCreateUserParams(arg,password)).Times(1).Return(user,nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t,http.StatusOK,recorder.Code)
			},
		},
	}
	for i := range testCases{
		tc := testCases[i]
		t.Run(tc.name,func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			server := newTestServer(t,store)
			recorder:= httptest.NewRecorder()

			data, err:= json.Marshal(tc.body)
			require.NoError(t,err)

			url := "/signup"
			request, err :=http.NewRequest(http.MethodPost,url,bytes.NewReader(data))
			require.NoError(t,err)
			server.router.ServeHTTP(recorder,request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser(t *testing.T)( user db.UserDatum,password string){
	password = util.RandString(6)
	hashedPassword , err := util.HashPassword(password)
	require.NoError(t,err)
	user = db.UserDatum{
		ID: int64(util.RandInt(1,1000)),
		Username: util.RandString(6),
		Password: hashedPassword,
		FirstName: util.RandString(6),
		LastName: util.RandString(6),
		Role: util.RandString(3),
		PhoneNumber: util.RandString(10),
	}
	return
}