# Plipkart
## Database Schema
![alt text](https://github.com/MeganViga/Plipkart/blob/main/PlipKart.png?raw=True)

## Available API's
- /signup
- /login
- /listproducts
- /getproductsbyname
- /getproductsbycolor
##### API's that only works with JWT Token
- /addaddress
- /deleteaddress/:id
- /addproduct
- /addproductcategory

  Order related API's yet to be developed
  
 ## API Examples
 
 * #### Get products API, which doesn't need JWT Token
 
 &nbsp;
 ![alt text](https://github.com/MeganViga/Plipkart/blob/main/GetAPIListProducts.png?raw=True)
 * #### Login API with Invalid Password
 &nbsp;
 ![alt text](https://github.com/MeganViga/Plipkart/blob/main/LoginAPIwithInvalidPassword.png?raw=True)
 
 * #### Login API, which give JWT Token login is Successfull
 &nbsp;
 ![alt text](https://github.com/MeganViga/Plipkart/blob/main/LoginAPI.png?raw=True)
 
 * #### List address API without Authorization Header(JWT Token)
 &nbsp;
 ![alt text](https://github.com/MeganViga/Plipkart/blob/main/ListAddressAPIWithoutAuthorizationHeader.png?raw=True)
 
 * #### List address API with invalid authorization header(JWT Token)
 &nbsp;
 ![alt text](https://github.com/MeganViga/Plipkart/blob/main/ListAddressAPIwithInvalidAuthorizationHeader.png?raw=True)
 
 * #### List Address API with valid authorizaion header(JWT Token)
 &nbsp;
 ![alt text](https://github.com/MeganViga/Plipkart/blob/main/ListAddressAPIWithValidAuthorizationHeader.png?raw=True)
 
 * #### List Address API with Expired Auhorization Header(JWT Token)
 &nbsp;
 ![alt text](https://github.com/MeganViga/Plipkart/blob/main/ListAddressAPIWithExpiredToken.png?raw=True)
 
 ## Example Unit test return using go mock for List Product API
 &nbsp;
 ```python
 func TestListProductApi(t *testing.T){
	products := []db.GetProductsRow{}
	for i :=0 ;i < 5;i++{
		products = append(products,randomProductWithColorAndQuantity())
	}
	//reponseList := createListProductsJson(products)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)
	//stubs
	store.EXPECT().
	GetProducts(gomock.Any()).
	Times(1).
	Return(products,nil)

	// start test server and send request
	server := newTestServer(t,store)

	//recorder to record the response
	recorder := httptest.NewRecorder()
	url := "/listproducts"
	request, err := http.NewRequest(http.MethodGet,url,nil)
	require.NoError(t,err)

	//request and Record response using recorder
	server.router.ServeHTTP(recorder,request)

	//compare response status with expected status
	require.Equal(t, http.StatusOK,recorder.Code)
}
```

## Dependencies
- Golang Version
```
go version go1.17.1 darwin/amd64
```
- Go Mock
```
go get github.com/golang/mock/mockgen@v1.6.0
```
- Require Package
```
github.com/stretchr/testify/require
```
- Gin Framework
```
github.com/gin-gonic/gin
```
- Golang Migrate
```
https://github.com/golang-migrate/migrate
```
- Viper
```
go get github.com/spf13/viper
```
- JWT Golang
```
go get -u github.com/golang-jwt/jwt/v4
```
- sqlc
```
https://sqlc.dev
```
## Run the Server
```
go run main.go
```