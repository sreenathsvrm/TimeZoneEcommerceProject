package response

import "time"

type RazorPayResponse struct {
	Email       string
	PhoneNumber string
	PaymentId   uint
	RazorpayKey string
	OrderId     interface{}
	AmountToPay float64
	Total       float64
}

type OrderResponse struct {
	ID                uint      `json:"order_ID"`
	UserID            uint      `json:"-"`
	OrderDate         time.Time `json:"order_date"`
	PaymentMethodID   uint      `json:"payment_method_id"`
	PaymentMethod     string    `json:"PaymentMethod"`
	ShippingAddressID uint      `json:"shipping_address_id"`
	House_number      string    `json:"house_number"`
	Street            string    `json:"street"`
	City              string    `json:"city"`
	District          string    `json:"district"`
	Pincode           int       `json:"pin_code"`
	Landmark          string    `json:"land_mark"`
	Discount          float64   `json:"discount"`
	OrderTotal        float64   `json:"order_total"`
	OrderStatusID     uint      `json:"order_status_id"`
	OrderStatus       string    `json:"orderStatus"`
	DeliveryUpdatedAt time.Time `json:"expected_delivery_time"`
}

type SalesReport struct {
	Id         string
	Name       string
	Payment_method string
	OrderDate      time.Time
	Order_Total    int
	Mobile      string
	HouseNumber    string
	Pincode        string

}

type AdminDashboard struct {
	CompletedOrders int     `json:"completed_orders,omitempty"`
	PendingOrders   int     `json:"pending_orders,omitempty"`
	CancelledOrders int     `json:"cancelled_orders,omitempty"`
	TotalOrders     int     `json:"total_orders,omitempty"`
	TotalOrderItems int     `json:"total_order_items,omitempty"`
	OrderValue      float64 `json:"order_value,omitempty"`
	CreditedAmount  float64 `json:"credited_amount,omitempty"`
	PendingAmount   float64 `json:"pending_amount,omitempty"`
	TotalUsers    int `json:"total_users,omitempty"`
	VerifiedUsers int `json:"verified_users,omitempty"`
	OrderedUsers  int `json:"ordered_users,omitempty"`
}
