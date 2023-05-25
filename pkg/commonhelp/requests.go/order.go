package requests

type RazorPayRequest struct {
	RazorPayPaymentId  string
	RazorPayOrderId    string
	Razorpay_signature string
}


type Update struct{
	OrderId int   `json:"order_id" binding:"required"`
	StatusId int   `json:"status_id" binding:"required"`
}