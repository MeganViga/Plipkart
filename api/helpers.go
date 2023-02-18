package api

import (

	db "github.com/MeganViga/Plipkart/db/sqlc"
)


var listProductsResponse = make(map[int64]listProducts)

type listProducts struct{
	ProductName string
	Price int64
	Colors []colorAndQuantity
}
type colorAndQuantity struct{
	Color string
	Quantity int32
}



func createListProductsJson(products []db.GetProductsRow)map[int64]listProducts{
	for _, item := range products{
		listProductsResponse[item.ID] = listProducts{}
		a := listProductsResponse[item.ID]
		a.Colors = []colorAndQuantity{}
		listProductsResponse[item.ID] = a

	}
	for _, item := range products{
		a := listProductsResponse[item.ID]
		if a.ProductName == ""{
			a.ProductName = item.ProductName
			a.Price = item.Price
		}
		b := colorAndQuantity{
			Color: item.Color,
			Quantity: item.Quantity,
		}
		a.Colors = append(a.Colors,b)
		listProductsResponse[item.ID] = a
	}
	return listProductsResponse
}

func createListProductsByNameJson(products []db.GetProductByNameRow)map[int64]listProducts{

	for _, item := range products{
		listProductsResponse[item.ID] = listProducts{}
		a := listProductsResponse[item.ID]
		a.Colors = []colorAndQuantity{}
		listProductsResponse[item.ID] = a

	}
	for _, item := range products{
		a := listProductsResponse[item.ID]
		if a.ProductName == ""{
			a.ProductName = item.ProductName
			a.Price = item.Price
		}
		b := colorAndQuantity{
			Color: item.Color,
			Quantity: item.Quantity,
		}
		a.Colors = append(a.Colors,b)
		listProductsResponse[item.ID] = a
	}
	return listProductsResponse
}

func createListProductsByColorJson(products []db.GetProductByColorRow)map[int64]listProducts{

	for _, item := range products{
		listProductsResponse[item.ID] = listProducts{}
		a := listProductsResponse[item.ID]
		a.Colors = []colorAndQuantity{}
		listProductsResponse[item.ID] = a

	}
	for _, item := range products{
		a := listProductsResponse[item.ID]
		if a.ProductName == ""{
			a.ProductName = item.ProductName
			a.Price = item.Price
		}
		b := colorAndQuantity{
			Color: item.Color,
			Quantity: item.Quantity,
		}
		a.Colors = append(a.Colors,b)
		listProductsResponse[item.ID] = a
	}
	return listProductsResponse
}




