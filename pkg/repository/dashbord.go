package repository

import (
	"ecommerce/pkg/commonhelp/response"

	"golang.org/x/net/context"
)

func (c *AdminDB) ViewSalesReport(ctx context.Context) ([]response.SalesReport, error) {
	var report []response.SalesReport
	FetchReports := `SELECT u.id, u.name, pt.payment_method AS payment_method, o.order_date, o.order_total,u.mobile, a.house_number,a.pincode           
	FROM orders o
	JOIN users u ON u.id = o.user_id
	JOIN payment_methods pt ON o.payment_method_id = pt.id
	JOIN addresses a ON a.id = o.shipping_address_id
	WHERE o.order_status_id = 1;`
	err := c.DB.Raw(FetchReports).Scan(&report).Error
	return report, err
}


