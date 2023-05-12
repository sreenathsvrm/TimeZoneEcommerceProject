package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) UserSignup(ctx context.Context, user urequest.Fusersign) (response.UserValue, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return response.UserValue{}, err
	}
	user.Password = string(hash)
	UserValue, err := c.userRepo.UserSignup(ctx, user)
	return UserValue, err
}

func (c *userUseCase) UserLogin(ctx context.Context, user urequest.Flogin) (string, error) {
	userData, err := c.userRepo.UserLogin(ctx, user.Email)
	var userstatus domain.Users
	if err != nil {
		return "", err
	} else if userData.ID == 0 {
		return "", fmt.Errorf("user not founf")
	}

	if user.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	fmt.Println("db", userData.Password, "user", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	if userstatus.IsBlocked {
		return "", fmt.Errorf("user is blocked")
	}
	fmt.Println("user_id on jwt generate ", userData.ID)
	claims := jwt.MapClaims{
		"id":  userData.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (c *userUseCase) OtpLogin(mobno string) (string, error) {
	id, err := c.userRepo.OtpLogin(mobno)
	if err != nil {
		return "", err
	} else if id == 0 {
		return "", errors.New("user not exist with given mobile number")
	}

	fmt.Println("user_id on otp_login", id)
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}
func (c *userUseCase)AddAdress(ctx context.Context,UserID int,address urequest.AddressReq)(domain.Address,error) {
	newAddress,err:=c.userRepo.AddAdress(ctx,UserID,address)
	return newAddress,err
}

func (c *userUseCase)UpdateAdress(ctx context.Context, UserID int, address urequest.AddressReq) (domain.Address, error)  {
	updated,err:=c.userRepo.UpdateAdress(ctx,UserID,address)
	return updated,err
}

func (c *userUseCase)VeiwAdress(ctx context.Context, UserID int) (domain.Address, error)  {
	adress,err:=c.userRepo.VeiwAdress(ctx,UserID)
	return adress,err
}