package handler

import (
	"ecommerce/pkg/api/utilhandler"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"

	"ecommerce/pkg/domain"
	services "ecommerce/pkg/usecase/interface"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUsecase services.AdminUsecase
}

func NewAdminHandler(Adusecase services.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		AdminUsecase: Adusecase,
	}
}

// @Summary SaveAdmin
// @ID SaveAdmin
// @Description Save admin with details.
// @Tags Admin
// @Accept json
// @Produce json
// @Param   inputs   body     domain.Admin{}   true  "Input Field"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/signup [post]
func (cr *AdminHandler) SaveAdmin(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	// defer cancel()
	var admin domain.Admin
	err := c.Bind(&admin)
	fmt.Println(admin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.AdminUsecase.SaveAdmin(c.Request.Context(), admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "signup Successfully",
		Errors:     nil,
	})

}

//@Summary LoginAdmin
//@Id LoginAdmin
//@Discription login with admin details.
//@tags Admin
// @Accept json
// @Produce json
// @Param   input   body     domain.Admin{}   true  "Input Field"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/login [post]
func (cr *AdminHandler) LoginAdmin(c *gin.Context) {
	var admin domain.Admin
	err := c.Bind(&admin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ss, err := cr.AdminUsecase.LoginAdmin(c.Request.Context(), admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuth", ss, 3600*24*1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "logined successfuly",
		Data:       nil,
		Errors:     nil,
	})

}

// AdminLogout
// @Summary Adminlogout
// @ID AdminLogout
// @Description Logout as a user exit from the ecommerce site
// @Tags Admin
// @Success 200 "success"
// @Failure 400 "failed"
// @Router /admin/logout [post]
func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "logout successfully",
		Data:       nil,
		Errors:     nil,
	})
}

//@Summary FindAllUsers
//@Id FindAllUsers
//@Discription list of users.
//@tags Admin
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param perPage query int false "Number of items to retrieve per page"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/findall [get]
func (cr *AdminHandler) FindAllUser(c *gin.Context) {

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

	list := requests.Pagination{
		Page:    uint(page),
		PerPage: uint(perPage),
	}
	users, err := cr.AdminUsecase.FindAllUser(c.Request.Context(), list)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "users not found",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "List of users",
		Data:       users,
		Errors:     nil,
	})

}

// BlockUser
// @Summary Admin can block a user
// @ID block-user
// @Description Admin can block a  user
// @Tags Admin
// @Accept json
// @Produce json
// @Param input body requests.BlockUser{} true "inputs"
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/block [patch]
func (cr *AdminHandler) BlockUser(c *gin.Context) {
	var body requests.BlockUser
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminId, err := utilhandler.GetAdminIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find AdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.AdminUsecase.BlockUser(body, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Block",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "User Blocked",
		Data:       nil,
		Errors:     nil,
	})
}

// UnblockUser
// @Summary Admin can unblock a blocked user
// @ID unblock-user
// @Description Admin can unblock a blocked user
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be unblocked"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/unblock/{user_id} [patch]
func (cr *AdminHandler) UnblockUser(c *gin.Context) {
	paramsId := c.Param("user_id")
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
	err = cr.AdminUsecase.UnblockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant unblock user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user unblocked",
		Data:       nil,
		Errors:     nil,
	})
}

// FindUserByID
// @Summary Admin can fetch a specific user details using user id
// @ID find-user-by-id
// @Description Admin can fetch a specific user details using user id
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be fetched"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/finduser/{user_id} [get]
func (cr *AdminHandler) FindUserByID(c *gin.Context) {
	paramsID := c.Param("user_id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse user id", Data: nil, Errors: err.Error()})
		return
	}
	user, err := cr.AdminUsecase.FindUserbyId(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed fetch user", Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200, Message: "Successfully fetched user details", Data: user, Errors: nil,
	})

}

// ViewSalesReport
// @Summary Admin can view sales report
// @ID view-sales-report
// @Description Admin can view the sales report
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/salesreport [get]
func (cr *AdminHandler) ViewSalesReport(ctx *gin.Context) {
	sales, err := cr.AdminUsecase.ViewSalesReport(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Sales report",
		Data:       sales,
		Errors:     nil,
	})

}

// DownloadSalesReport
// @Summary Admin can download sales report
// @ID download-sales-report
// @Description Admin can download sales report in .csv format
// @Tags Admin
// @Accept json
// @Produce json
// @Failure 400 {object} response.Response
// @Router /admin/salesreport/download [get]
func (cr *AdminHandler) DownloadSalesReport(ctx *gin.Context) {
	sales, err := cr.AdminUsecase.ViewSalesReport(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// Set headers so browser will download the file
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment;filename=sales.csv")

	// Create a CSV writer using our response writer as our io.Writer
	wr := csv.NewWriter(ctx.Writer)

	// Write CSV header row
	headers := []string{"Id", "Name", "Payment_method", "OrderDate", "Order_Total", "Mobile", "HouseNumber", "Pincode"}
	if err := wr.Write(headers); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Write data rows
	for _, sale := range sales {
		row := []string{sale.Id, sale.Name, sale.Payment_method, sale.OrderDate.Format("2006-01-02 15:04:05"), strconv.Itoa(sale.Order_Total),
			sale.Mobile, sale.HouseNumber, sale.Pincode,
		}
		if err := wr.Write(row); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	// Flush the writer's buffer to ensure all data is written to the client
	wr.Flush()
	if err := wr.Error(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}
