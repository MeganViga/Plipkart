package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/MeganViga/Plipkart/db/mock"
	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/MeganViga/Plipkart/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddProductApi(t *testing.T){
	product := randomProduct()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)
	arg := addProductRequestParams{
		ProductName: product.ProductName,
		Price: product.Price,
	}
	//stubs
	store.EXPECT().
	AddProduct(gomock.Any(),arg).
	Times(1).
	Return(product)

	// start test server and send request
	server := newTestServer(t,store)

	recorder := httptest.NewRecorder()
	databody , err := json.Marshal(&arg)
	require.NoError(t,err)
	
	body := bytes.NewReader(databody)
	url := "/addproduct"
	request, err := http.NewRequest(http.MethodGet,url,body)
	require.NoError(t,err)

	server.router.ServeHTTP(recorder,request)
	require.Equal(t, http.StatusOK,recorder.Code)
}

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

func randomProduct()db.Product{
	return db.Product{
		ID: int64(util.RandInt(1,1000)),
		ProductName: util.RandString(10),
		Price: int64(util.RandInt(100,200)),
		CreatedAt: time.Now(),
		ModifiedAt: time.Now(),
	}
}

func randomProductWithColorAndQuantity()db.GetProductsRow{
	return db.GetProductsRow{
		ID: int64(util.RandInt(1,1000)),
		ProductName: util.RandString(10),
		Price: int64(util.RandInt(100,200)),
		Color: util.RandString(6),
		Quantity: int32(util.RandInt(1,10)),
	}
}

