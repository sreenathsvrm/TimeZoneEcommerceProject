package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
)

type OtpUseCase interface {
	SendOTP(ctx context.Context, mobno requests.OTPreq) (string, error)
	VerifyOTP(ctx context.Context, userData requests.Otpverifier) error
}
