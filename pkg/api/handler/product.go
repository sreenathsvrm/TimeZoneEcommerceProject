package handler

import (
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	services "ecommerce/pkg/usecase/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUsecase services.ProductUsecase
}

func NewproductHandler(ProductUsecase services.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: ProductUsecase,
	}
}

// AddNEWCategory
// @Summary Create new product category
// @ID create-category
// @Description Admin can create new category from admin panel
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_name body requests.Category true "New category name"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/category/add [post]
func (cr *ProductHandler) Addcategory(c *gin.Context) {
	var newcategory requests.Category
	err := c.Bind(&newcategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "faild to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	NewCategoery, err := cr.ProductUsecase.Addcategory(c.Request.Context(), newcategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't creat newcategory",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "newCategory Created",
		Data:       NewCategoery,
		Errors:     nil,
	})
}

// UpdateCategory
// @Summary Admin can update category details
// @ID update-category
// @Description Admin can update category details
// @Tags Product Category
// @Accept json
// @Produce json
// @Param id path string true "ID of the Category to be updated"
// @Param category_details body requests.Category true "category info"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/category/update/{id} [patch]
func (cr *ProductHandler) UpdatCategory(c *gin.Context) {
	var category requests.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updatedCategory, err := cr.ProductUsecase.UpdatCategory(c.Request.Context(), category, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Updated",
		Data:       updatedCategory,
		Errors:     nil,
	})
}

// DeleteCategory
// @Summary Admin can delete a category
// @ID delete-category
// @Description Admin can delete a category
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/category/delete/{category_id} [delete]
func (cr *ProductHandler) DeleteCategory(c *gin.Context) {
	parmasId := c.Param("category_id")
	id, err := strconv.Atoi(parmasId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.ProductUsecase.DeleteCatagory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't dlete category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category deleted",
		Data:       nil,
		Errors:     nil,
	})

}

// ListAllCategories
// @Summary View all available categories
// @ID view-all-categories
// @Description Admin, users and unregistered users can see all the available categories
// @Tags Product Category
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/category/showall [get]
func (cr *ProductHandler) ListCategories(c *gin.Context) {
	categories, err := cr.ProductUsecase.Listallcatagory(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "List of categories",
		Data:       categories,
		Errors:     nil,
	})
}

// FindCategoryByID
// @Summary Fetch details of a specific category using category id
// @ID find-category-by-id
// @Description Users and admins can fetch details of a specific category using id
// @Tags Product Category
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/category/disply/{id} [get]
func (cr *ProductHandler) DisplayCategory(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	category, err := cr.ProductUsecase.ShowCatagory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Name",
		Data:       category,
		Errors:     nil,
	})
}

// SaveProduct
// @Summary Admin can create new product listings
// @ID create-product
// @Description Admins can create new product listings
// @Tags Product
// @Accept json
// @Produce json
// @Param new_product_details body requests.Product true "new product details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product/save [post]
func (cr *ProductHandler) SaveProduct(c *gin.Context) {
	var product requests.Product
	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newProduct, err := cr.ProductUsecase.SaveProduct(c.Request.Context(), product)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't add product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Added",
		Data:       newProduct,
		Errors:     nil,
	})

}

// UpdateProduct
// @Summary Admin can update Product details
// @ID update-Product
// @Description Admin can update Product details
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "ID of the product to be updated"
// @Param category_details body requests.Product true "Product info"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product/updateproduct/{id} [patch]
func (cr *ProductHandler) UpdateProduct(c *gin.Context) {
	var product requests.Product
	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updatedCategory, err := cr.ProductUsecase.UpdateProduct(c.Request.Context(), id, product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Updated",
		Data:       updatedCategory,
		Errors:     nil,
	})
}

// DeleteProduct
// @Summary Admin can delete a product
// @ID delete-product
// @Description Admin can delete a product
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product/delete/{product_id} [delete]
func (cr *ProductHandler) DeleteProduct(c *gin.Context) {
	parmasId := c.Param("product_id")
	id, err := strconv.Atoi(parmasId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.ProductUsecase.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't delete product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product deleted",
		Data:       nil,
		Errors:     nil,
	})

}

// ViewAllProducts
// @Summary Admins and users can see all available products
// @ID user-view-all-products
// @Description  users can ses all available products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param perPage query int false "Number of items to retrieve per page"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/product/ViewAllProducts [get]
func (cr *ProductHandler) ViewAllProducts(c *gin.Context) {

	//fetch query parameters
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "invalied pagenumber",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	perPage, err := strconv.Atoi(c.Query("perPage"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "invalied perPage",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	listProducts := requests.Pagination{
		Page:    uint(page),
		PerPage: uint(perPage),
	}

	products, err := cr.ProductUsecase.ViewAllProducts(c.Request.Context(), listProducts)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "products not found",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "List of products",
		Data:       products,
		Errors:     nil,
	})
}

// FindProductByID
// @Summary Admins and users can see products with product id
// @ID find-product-by-id
// @Description Admins and users can see products with product id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product/ViewProduct/{id} [get]
func (cr *ProductHandler) VeiwProduct(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	product, err := cr.ProductUsecase.VeiwProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       product,
		Errors:     nil,
	})
}
