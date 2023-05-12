package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/config"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

type Orderusecase struct {
	cartRepo  interfaces.CartRepo
	orderRepo interfaces.OrderRepo
}

func NewOrderUseCase(orderRepo interfaces.OrderRepo, cartRepo interfaces.CartRepo) services.Orderusecase {
	return &Orderusecase{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (c *Orderusecase) PlaceOrder(ctx context.Context, UserID, paymentMethodId int) (domain.Orders, error) {
	order, err := c.orderRepo.OrderAll(ctx, UserID, paymentMethodId)
	return order, err
}
func (c *Orderusecase) Razorpay(ctx context.Context, UserID, paymentMethodId int) (response.RazorPayResponse, error) {

	cart, err := c.cartRepo.FindCartByUserID(ctx, UserID)
	if err != nil {
		return response.RazorPayResponse{}, err
	}
	if cart.Total_price == 0 {
		return response.RazorPayResponse{}, fmt.Errorf("there is no products in your list")
	}

	razorpayKey := config.Getconfig().RAZOR_PAY_KEY
	razorpaySecret := config.Getconfig().RAZOR_PAY_SECRET

	client := razorpay.NewClient(razorpayKey, razorpaySecret)

	razorPayAmount := cart.Total_price * 100

	data := map[string]interface{}{
		"amount":   razorPayAmount,
		"currency": "INR",
		"receipt":  "reciept_id",
	}
	// create an order on razor pay
	order, err := client.Order.Create(data, nil)

	if err != nil {
		return response.RazorPayResponse{}, fmt.Errorf("faild to create razorpay order")
	}

	return response.RazorPayResponse{
		Email:       "",
		PhoneNumber: "",
		RazorpayKey: razorpayKey,
		PaymentId:   uint(paymentMethodId),
		OrderId:     order["id"],
		Total:       razorPayAmount,
		AmountToPay: cart.Total_price,
	}, nil
}

func (c *Orderusecase) VerifyRazorPay(ctx context.Context, body urequest.RazorPayRequest) error {
	razorpayKey := config.Getconfig().RAZOR_PAY_KEY
	razorPaySecret := config.Getconfig().RAZOR_PAY_SECRET

	//varify signature
	data := body.RazorPayOrderId + "|" + body.RazorPayPaymentId
	h := hmac.New(sha256.New, []byte(razorPaySecret))
	_, err := h.Write([]byte(data))
	if err != nil {
		return errors.New("faild to veify signature")
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(body.Razorpay_signature)) != 1 {
		return errors.New("razorpay signature not match")
	}

	// then vefiy payment
	client := razorpay.NewClient(razorpayKey, razorPaySecret)

	// fetch payment and vefify
	payment, err := client.Payment.Fetch(body.RazorPayPaymentId, nil, nil)

	if err != nil {
		return err
	}

	// check payment status
	if payment["status"] != "captured" {
		return errors.New("faild to verify razorpay payment")
	}

	return nil
}

func (c *Orderusecase) CancelOrder(ctx context.Context, orderId, userId int) error {
	err := c.orderRepo.CancelOrder(ctx, orderId, userId)
	return err
}

func (c *Orderusecase) Listorders(ctx context.Context, userid int) ([]domain.Orders, error) {
	var orders []domain.Orders
	orders, err := c.orderRepo.Listorders(ctx, userid)
	return orders, err
}

func (c *Orderusecase) Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error) {
	order, err = c.orderRepo.Listorder(ctx, Orderid, UserId)
	return order, err
}

func (c *Orderusecase) ReturnOrder(userId, orderId int) (float64, error) {
	total, err := c.orderRepo.ReturnOrder(userId, orderId)
	return total, err
}
