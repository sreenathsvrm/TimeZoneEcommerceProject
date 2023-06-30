package usecase

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
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

func (c *userUseCase) UserSignup(ctx context.Context, user requests.Usersign) (response.UserValue, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return response.UserValue{}, err
	}
	user.Password = string(hash)
	UserValue, err := c.userRepo.UserSignup(ctx, user)
	return UserValue, err
}

func (c *userUseCase) UserLogin(ctx context.Context, user requests.Login) (string, error) {
	userData, err := c.userRepo.UserLogin(ctx, user.Email)
	var userstatus domain.Users
	if err != nil {
		return "", err
	} else if userData.ID == 0 {
		return "", fmt.Errorf("user not found")
	}

	if user.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	if userstatus.IsBlocked {
		return "", fmt.Errorf("user is blocked")
	}
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
func (c *userUseCase) AddAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error) {
	newAddress, err := c.userRepo.AddAdress(ctx, UserID, address)
	return newAddress, err
}

func (c *userUseCase) UpdateAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error) {
	updated, err := c.userRepo.UpdateAdress(ctx, UserID, address)
	return updated, err
}

func (c *userUseCase) VeiwAdress(ctx context.Context, UserID int) (domain.Address, error) {
	adress, err := c.userRepo.VeiwAdress(ctx, UserID)
	return adress, err
}

func (c *userUseCase) AddToWishList(ctx context.Context, wishList domain.WishList) error {

	product, err := c.userRepo.FindProduct(ctx, wishList.ProductID)
	if err != nil {
		return err
	} else if product.Id == 0 {
		return errors.New("invalid product_id")
	}

	checkWishList, err := c.userRepo.FindWishListItem(ctx, wishList.ProductID, wishList.UserID)
	if err != nil {
		return err
	} else if checkWishList.ID != 0 {
		return errors.New("product is  already exist on wishlist")
	}

	if err := c.userRepo.SaveWishListItem(ctx, wishList); err != nil {
		return err
	}

	return nil
}

func (c *userUseCase) RemoveFromWishList(ctx context.Context, wishList domain.WishList) error {

	product, err := c.userRepo.FindProduct(ctx, wishList.ProductID)
	if err != nil {
		return err
	} else if product.Id == 0 {
		return errors.New("invalid product_id")
	}

	// check the productItem already exist on wishlist for user
	wishList, err = c.userRepo.FindWishListItem(ctx, wishList.ProductID, wishList.UserID)
	if err != nil {
		return err
	} else if wishList.ID == 0 {
		return errors.New("productItem not found in wishlist")
	}

	return c.userRepo.RemoveWishListItem(ctx, wishList)
}

func (c *userUseCase) ListWishlist(ctx context.Context, userID uint) ([]response.Wishlist, error) {
	return c.userRepo.FindAllWishListItemsByUserID(ctx, userID)
}
