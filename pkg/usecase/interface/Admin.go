package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
)

type AdminUsecase interface {
	SaveAdmin(ctx context.Context, admin domain.Admin) error
	LoginAdmin(ctx context.Context, admin domain.Admin) (string, error)
	FindAllUser(ctx context.Context, pagination urequest.Pagination) (users []response.UserValue, err error)
	BlockUser(body urequest.BlockUser, AdminId int) error
	UnblockUser(id int) error
	FindUserbyId(ctx context.Context, userID int) (domain.Users,error)
	ViewSalesReport(ctx context.Context) ([]response.SalesReport, error)
}
