package api

import (
	"fmt"

	"github.com/MeganViga/Plipkart/token"
	"github.com/MeganViga/Plipkart/util"

	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/gin-gonic/gin"
)
type Server struct{
	config util.Config
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config,store db.Store)(*Server,error){
	tokenMaker, err := token.NewJwt_Maker("aPVLkh3OoBTmRSTqz9n9OqnnQ0WKfNx9")
	if err != nil{
		return nil, fmt.Errorf("cannot create token maker: %w",err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,

	}


	router := gin.Default()
	router.POST("/signup",server.userSignup)
	router.POST("/login",server.userLogin)
	// router.POST("/addaddress",server.addAddress)
	// router.GET("/deleteaddress/:id",server.deleteAddress)
	// router.POST("/addproduct",server.addProduct)
	// router.POST("/addproductcategory",server.addProductCategory)
	router.GET("/listproducts",server.listProducts)
	router.GET("/getproductsbyname",server.listProductsByName)
	router.GET("/getproductsbycolor",server.listProductsByColor)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/addaddress",server.addAddress)
	authRoutes.GET("/deleteaddress/:id",server.deleteAddress)
	authRoutes.POST("/addproduct",server.addProduct)
	authRoutes.POST("/addproductcategory",server.addProductCategory)

	server.router = router
	return server, nil
}

func (s *Server)StartServer(address string){
	s.router.Run(address)
}

func (s *Server)ShowError(err error)gin.H{
	return gin.H{"error":err.Error()}
}

func errResponse(err error)gin.H{
	return gin.H{"error":err.Error()}
}