package usecase

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"
	"errors"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase struct {
	AdminRepo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUsecase {
	return &AdminUsecase{
		AdminRepo: repo,
	}
}

func (c *AdminUsecase) SaveAdmin(ctx context.Context, admin domain.Admin) error {
	
	if admin, err := c.AdminRepo.FindAdmin(ctx, admin); err != nil {
		return err
	} else if admin.ID != 0 {
		return errors.New(" already exist with the same details")
	}
	// generate a hashed password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)

	if err != nil {
		return errors.New("faild to generate hashed password for admin")
	}
	// set the hashed password
	admin.Password = string(hashPass)
	
	return c.AdminRepo.SaveAdmin(ctx, admin)
}

func (c *AdminUsecase) LoginAdmin(ctx context.Context, admin domain.Admin) (string, error) {
	DBadmin, err := c.AdminRepo.FindAdmin(ctx, admin)

	if err != nil {

		return "", err

	} else if DBadmin.ID == 0 {
		return "", errors.New("this id not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(DBadmin.Password), []byte(admin.Password)) != nil {
		return "", errors.New("incorrect password")
	}

	claims := &jwt.MapClaims{
		"id":  admin.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (c *AdminUsecase) FindAllUser(ctx context.Context, pagination urequest.Pagination) (users []response.UserValue, err error) {

	users, err = c.AdminRepo.FindAllUser(ctx, pagination)

	if err != nil {
		return nil, err
	}
	var respond []response.UserValue
	copier.Copy(&respond, &users)
	return respond, nil
}


func (c *AdminUsecase)  BlockUser(body urequest.BlockUser, adminId int) error {
	err := c.AdminRepo.BlockUser(body, adminId)
	return err
}

func (c *AdminUsecase)  UnblockUser(id int) error {
	 err := c.AdminRepo.UnblockUser(id)
	return err
}

func (c *AdminUsecase) FindUserbyId(ctx context.Context, userID int) (domain.Users,error)  {
	user, err := c.AdminRepo.FindUserbyId(ctx, userID)
	return user, err
}