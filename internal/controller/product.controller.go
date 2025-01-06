package controller

import (
	"tudo-thrfting-clothes-server/global"
	"tudo-thrfting-clothes-server/internal/model"
	"tudo-thrfting-clothes-server/internal/service"
	"tudo-thrfting-clothes-server/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		productService: service.NewProductService(),
	}
}

// Get list of products
// @Summary Get list of products
// @Description Get list of products
// @Tags Product
// @Accept  json
// @Produce  json
// @Success 200 {object} []model.Products "List of products"
// @Router /product [get]
func (pc *ProductController) GetListProduct(c *gin.Context) {
	// Get the list of products and handle any errors
	products, err := pc.productService.GetListProduct()

	if err != nil {
		// Call the error response function with a custom error message
		response.ErrorReponse(c, 500, "Failed to retrieve products from the database")
		return
	}

	// Return success response if products are retrieved successfully
	response.SuccessReponse(c, response.ErrCodeSuccess, products)
}

// AddProduct handles adding a new product.
// @Summary Add a new product
// @Description Add a product to the inventory
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product body model.CreateProduct true "Product Data"
// @Success 201 {object} model.CreateProduct "Created Product"
// @Router /product [post]
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

// CreateProductBrand handles adding a new product brand.
// @Summary Add a new product brand
// @Description Add a product brand to the inventory
// @Tags Product
// @Accept  json
// @Produce  json
// @Param brand body model.CreateProductBrand true "Product brand Data"
// @Success 201 {object} model.CreateProductBrand "Created Product brand"
// @Router /product/brand [post]
func (pc *ProductController) CreateProductBrand(c *gin.Context) {
	var data model.CreateProductBrand

	if err := c.ShouldBindJSON(&data); err != nil {
		// Call the error response function with a custom error message
		global.Logger.Error("Failed to bind JSON data", zap.Error(err))
		response.ErrorReponse(c, response.ErrPayloadData, err.Error())
		return
	}
	// Create a new product and handle any errors
	brand, err := pc.productService.CreateProductBrand(data)

	if err != nil {
		// Call the error response function with a custom error message
		response.ErrorReponse(c, 500, "Failed to create product in the database")
		return
	}

	// Return success response if product is created successfully
	response.SuccessReponse(c, response.ErrCodeSuccess, brand)
}
