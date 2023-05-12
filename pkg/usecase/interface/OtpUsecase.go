package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/urequest"
)

type OtpUseCase interface {
	SendOTP(ctx context.Context, mobno urequest.OTPreq) (string, error)
	VerifyOTP(ctx context.Context, userData urequest.Otpverifier) error
}
