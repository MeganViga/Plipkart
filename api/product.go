package api

import (
	"net/http"
	"time"

	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/gin-gonic/gin"
)

type addProductRequestParams struct{
	ProductName string `json:"product_name" binding:"required"`
	Price	int64	`json:"price" binding:"required"`
}
func(s *Server)addProduct(ctx *gin.Context){
	var request addProductRequestParams
	if err := ctx.ShouldBindJSON(&request); err != nil{
		ctx.JSON(http.StatusBadRequest,s.ShowError(err))
		return
	}
	arg := db.AddProductParams{
		ProductName: request.ProductName,
		Price: request.Price,
		CreatedAt: time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
	}
	product, err := s.store.AddProduct(ctx,arg)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
	}
	ctx.JSON(http.StatusOK,product)
}
type addProductCategoryRequestParams struct{
	ProductID int64 `json:"product_id" binding:"required"`
	Color string	`json:"color" binding:"required"`
	Quantity int32 	`json:"quantity" binding:"required"`
}
func (s *Server)addProductCategory(ctx *gin.Context){
	var request addProductCategoryRequestParams
	if err := ctx.ShouldBindJSON(&request); err != nil{
		ctx.JSON(http.StatusBadRequest,s.ShowError(err))
		return
	}
	arg := db.AddProductCategoryParams{
		ProductID: request.ProductID,
		Color: request.Color,
		Quantity: request.Quantity,
	}
	product_category, err := s.store.AddProductCategory(ctx,arg)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
		return
	}
	ctx.JSON(http.StatusOK,product_category)
}

func (s *Server)listProducts(ctx *gin.Context){
	products, err := s.store.GetProducts(ctx)
	listProductsResponse := createListProductsJson(products)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
		return
	}
	
	ctx.JSON(http.StatusOK,listProductsResponse)
}
type listProductByNameRequestParams struct{
	ProductName string `form:"product_name" binding:"required"`
}
func (s *Server)listProductsByName(ctx *gin.Context){
	var request listProductByNameRequestParams
	if err := ctx.ShouldBindQuery(&request); err != nil{
		ctx.JSON(http.StatusBadRequest,s.ShowError(err))
		return
	}
	product, err := s.store.GetProductByName(ctx,request.ProductName)
	listProductsByNameResponse := createListProductsByNameJson(product)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
		return
	}
	
	ctx.JSON(http.StatusOK,listProductsByNameResponse)
}
type listProductsByColorRequestParams struct{
	ProductID int64 `form:"product_id" binding:"required"`
	Color string	`form:"color" binding:"required"`
}
func (s *Server)listProductsByColor(ctx *gin.Context){
	var request listProductsByColorRequestParams
	if err := ctx.ShouldBindQuery(&request); err != nil{
		ctx.JSON(http.StatusBadRequest,s.ShowError(err))
		return
	}
	arg := db.GetProductByColorParams{
		ProductID: request.ProductID,
		Color: request.Color,
	}
	product, err := s.store.GetProductByColor(ctx,arg)
	listProductsByColorResponse := createListProductsByColorJson(product)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,s.ShowError(err))
		return
	}
	
	ctx.JSON(http.StatusOK,listProductsByColorResponse)
}