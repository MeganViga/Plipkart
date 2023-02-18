package api
import (
	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/gin-gonic/gin"
	"database/sql"
)
type Server struct{
	store db.Store
	router *gin.Engine
}

func NewServer(DB *sql.DB)Server{
	server := Server{
		store: db.NewStore(DB),
	}


	router := gin.Default()
	router.POST("/signup",server.userSignup)
	router.POST("/login",server.userLogin)
	router.POST("/addaddress",server.addAddress)
	router.GET("/deleteaddress/:id",server.deleteAddress)
	router.POST("/addproduct",server.addProduct)
	router.POST("/addproductcategory",server.addProductCategory)
	router.GET("/listproducts",server.listProducts)
	router.GET("/getproductsbyname",server.listProductsByName)
	router.GET("/getproductsbycolor",server.listProductsByColor)

	server.router = router
	return server
}

func (s *Server)StartServer(){
	s.router.Run(":4789")
}

func (s *Server)ShowError(err error)gin.H{
	return gin.H{"error":err.Error()}
}