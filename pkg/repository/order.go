package repository

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type OrderDB struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepo {
	return &OrderDB{
		DB: DB,
	}
}

func (c *OrderDB) OrderAll(ctx context.Context, UserID, paymentMethodId int) (domain.Orders, error) {
	tx := c.DB.Begin()
	var cart domain.Cart
	findquery := `SELECT *FROM carts WHERE user_id=?`
	err := tx.Raw(findquery, UserID).Scan(&cart).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	if cart.Total_price == 0 {
		tx.Rollback()
		return domain.Orders{}, fmt.Errorf("please makesure you add those items to cart")
	}
	if cart.Total_price == 0 {
		tx.Rollback()
		return domain.Orders{}, err
	}
	// -------AddressFetch
	var address domain.Address
	findaddress := `SELECT *FROM addresses WHERE user_id=?`
	err = tx.Raw(findaddress, UserID).Scan(&address).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	if address.ID == 0 {
		tx.Rollback()
		return domain.Orders{}, err
	}

	var order domain.Orders

	insetOrder := `INSERT INTO orders (user_id,order_date,payment_method_id,shipping_address_id,order_total,order_status_id)
		VALUES($1,NOW(),$2,$3,$4,1) RETURNING *`
	err = tx.Raw(insetOrder, UserID, paymentMethodId, address.ID, cart.Total_price).Scan(&order).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	var cartItemes []urequest.CartItems
	cartDetail := `SELECT ci.product_id,ci.qty,p.prize,p.qty_in_stock  from cart_items ci join products p on ci.product_id = p.id where ci.cart_id=$1`
	err = tx.Raw(cartDetail, cart.Id).Scan(&cartItemes).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	//Add the items in the cart into the orderline
	for _, items := range cartItemes {
		if items.Qty > items.Qty_In_Stock {
			return domain.Orders{}, fmt.Errorf("out of stock")
		}
		insetOrder := `INSERT INTO order_lines (order_id,product_id,qty,price) VALUES($1,$2,$3,$4)`
		err = tx.Exec(insetOrder, order.ID, items.ProductId, items.Qty, items.Price).Error

		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	//Remove the product from the cart_items
	for _, items := range cartItemes {
		removefromCart := `DELETE FROM cart_items WHERE cart_id =$1 AND product_id=$2`
		err = tx.Exec(removefromCart, cart.Id, items.ProductId).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	for _, items := range cartItemes {
		updateQty := `UPDATE products SET qty_in_stock=products.qty_in_stock-$1 WHERE id=$2`
		err = tx.Exec(updateQty, items.Qty, items.ProductId).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	PaymentDetails := `INSERT INTO payment_details
			(orders_id,
			order_total,
			payment_method_id,
			payment_status_id,
			updated_at)
			VALUES($1,$2,$3,$4,NOW())`
	if err = tx.Exec(PaymentDetails, order.ID, order.OrderTotal, paymentMethodId, 1).Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	return order, nil

}

func (c *OrderDB) CancelOrder(ctx context.Context, orderId, userId int) error {
	tx := c.DB.Begin()

	//find the orderd product and qty and update the product with those
	var items []urequest.CartItems
	findProducts := `SELECT product_id,qty FROM order_lines WHERE order_id=?`
	err := tx.Raw(findProducts, orderId).Scan(&items).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(items) == 0 {
		return fmt.Errorf("no order found with this id")
	}
	for _, item := range items {
		updateProductItem := `UPDATE products SET qty_in_stock=qty_in_stock+$1 WHERE id=$2`
		err = tx.Exec(updateProductItem, item.Qty, item.ProductId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//Remove the items from order_lines
	remove := `DELETE FROM order_lines WHERE order_id=$1`
	err = tx.Exec(remove, orderId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//update the order status as canceled
	cancelOrder := `UPDATE orders SET order_status_id=$1 WHERE id=$2 AND user_id=$3`
	err = tx.Exec(cancelOrder, 5, orderId, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *OrderDB) Listorders(ctx context.Context) ([]response.OrderResponse, error) {
	var orders []response.OrderResponse
	Query := `SELECT o.id, o.user_id, o.order_date, o.payment_method_id, pm.payment_method, o.shipping_address_id,a.house_number,a.street,a.city,a.district,a.pincode,a.landmark,o.order_total, o.order_status_id, os.order_status, o.delivery_updated_at
	FROM orders o
	JOIN users u ON o.user_id = u.id
	JOIN payment_methods pm ON o.payment_method_id = pm.id
	JOIN addresses a ON o.shipping_address_id = a.id
	JOIN order_statuses os ON o.order_status_id = os.id`
	err := c.DB.Raw(Query).Scan(&orders).Error
    if err!=nil{
	return orders,err
    }


	// query := `SELECT id,house_number,street,city,district,pincode,landmark,
	// FROM addresses WHERE id= ?`
	// var address response.Addressrespond

	// for i, order := range orders {

	// 	if c.DB.Raw(query, order.ShippingAddressID).Scan(&address).Error != nil {
	// 		return orders, errors.New("faild to fetch addresses")
	// 	}
	// 	orders[i].Address = address
	// }
	  return orders, nil
}



func (c *OrderDB) Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error) {
	findOrder := `SELECT * FROM orders WHERE user_id=$1 AND id=$2`
	err = c.DB.Raw(findOrder, Orderid, UserId).Scan(&order).Error
	return order, err
}

func (c *OrderDB) ReturnOrder(userId, orderId int) (float64, error) {
	var orders domain.Orders
	Query := `SELECT * FROM orders WHERE user_id=$1 AND id=$2`
	err := c.DB.Raw(Query, userId, orderId).Scan(&orders).Error
	if err != nil {
		return 0, err
	}
	if orders.OrderStatusID != 3 {
		return 0, fmt.Errorf("the order is not deleverd")
	}
	returnOder := `UPDATE orders SET order_status_id=$1 WHERE id=$2`
	err = c.DB.Exec(returnOder, 6, orderId).Error
	if err != nil {
		return 0, err
	}
	return orders.OrderTotal, nil

}

//------order_management for adminside-------//

//show aall orderstatuses for  for admin
func (c *OrderDB) ListofOrderStatuses(ctx context.Context) (status []domain.OrderStatus, err error) {
	Quary := `SELECT * FROM order_statuses ORDER BY order_status DESC;`
	err = c.DB.Raw(Quary).Scan(&status).Error
	if err != nil {
		return status, errors.New("there will be some issues")
	}
	return status, err
}

//admin want to update the orderstatus
func (c *OrderDB) AdminListorders(ctx context.Context, pagination urequest.Pagination) (orders []domain.Orders, err error) {
	limit := pagination.PerPage
	offset := (pagination.Page - 1) * limit

	fmt.Println(limit, offset)
	query := `SELECT * FROM orders ORDER BY order_date  DESC LIMIT $1 OFFSET $2`
	err = c.DB.Raw(query, limit, offset).Scan(&orders).Error
	return orders, err
}

func (c *OrderDB) UpdateOrderStatus(ctx context.Context, update urequest.Update) error {
	fmt.Println("iam repo")
	fmt.Println(update.OrderId, update.StatusId)
	Quary := `UPDATE orders SET order_status_id=$1 WHERE id=$2`
	err := c.DB.Exec(Quary, update.StatusId, update.OrderId).Error
	if err != nil {
		return err
	}
	return nil
}
