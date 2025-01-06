package controller

import (
	"tudo-thrfting-clothes-server/internal/model"
	"tudo-thrfting-clothes-server/internal/service"
	"tudo-thrfting-clothes-server/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		productService: service.NewProductService(),
	}
}

// GetListProduct handles the API endpoint for getting the list of products
func (pc *ProductController) GetListProduct(c *gin.Context) {
	// Get the list of products and handle any errors
	products, err := pc.productService.GetListProduct()

	if err != nil {
		// Call the error response function with a custom error message
		response.ErrorReponse(c, 500, "Failed to retrieve products from the database")
		return
	}

	// Return success response if products are retrieved successfully
	response.SuccessReponse(c, 200, products)
}

// CreateProduct handles the API endpoint for creating a new product

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var data model.CreateProduct

	if err := c.ShouldBindJSON(&data); err != nil {
		// Call the error response function with a custom error message
		response.ErrorReponse(c, response.ErrPayloadData, err.Error())
		return
	}
	// Create a new product and handle any errors
	product, err := pc.productService.CreateProduct(data)

	if err != nil {
		// Call the error response function with a custom error message
		response.ErrorReponse(c, 500, "Failed to create product in the database")
		return
	}

	// Return success response if product is created successfully
	response.SuccessReponse(c, 200, product)
}
