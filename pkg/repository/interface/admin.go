package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
)

type AdminRepository interface {
	 FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	 FindAllUser(ctx context.Context, pagination  urequest.Pagination) (users []response.UserValue, err error)
     SaveAdmin(ctx context.Context, admin domain.Admin)error
	 BlockUser(body urequest.BlockUser, AdminId int) error
	 UnblockUser(id int) error
	 FindUserbyId(ctx context.Context, userID int) (domain.Users,error)
	 ViewSalesReport(ctx context.Context) ([]response.SalesReport, error)
}
