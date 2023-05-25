package repository

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	"errors"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) UserSignup(ctx context.Context, user requests.Usersign) (userValue response.UserValue, err error) {
	// var userValue response.UserValue
	insertQuery := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4) 
					RETURNING id,name,email,mobile`
	err = c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userValue).Error

	return userValue, err
}

func (c *userDatabase) UserLogin(ctx context.Context, Email string) (domain.Users, error) {
	var userData domain.Users
	err := c.DB.Raw("SELECT * FROM users WHERE email=?", Email).Scan(&userData).Error
	return userData, err
}

func (c *userDatabase) OtpLogin(mbnum string) (int, error) {
	var id int
	query := "SELECT id FROM users WHERE mobile=?"
	err := c.DB.Raw(query, mbnum).Scan(&id).Error
	return id, err
}

func (c *userDatabase) AddAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error) {
	var existaddress, newAddress domain.Address

	findaddressbyUser := "SELECT *FROM addresses WHERE user_id=?"

	err := c.DB.Raw(findaddressbyUser, UserID).Scan(&existaddress).Error
	if err != nil {
		return domain.Address{}, err
	}

	if existaddress.ID == 0 || existaddress.UserID == 0 {
		AddAddressQuery := `	INSERT INTO addresses(
			user_id, house_number, street, city, district, pincode, landmark) 
			VALUES($1,$2,$3,$4,$5,$6, $7) RETURNING *`

		err := c.DB.Raw(AddAddressQuery, UserID, address.HouseNumber, address.Street, address.City, address.District, address.Pincode, address.Landmark).Scan(&newAddress).Error
		return newAddress, err
	} else {
		//	address is already there, Edit it
		EditAddressQuery := `	UPDATE addresses SET
								house_number = $1, street = $2, city = $3, district = $4, pincode = $5, landmark = $6
								WHERE user_id = $7
								RETURNING *`
		err := c.DB.Raw(EditAddressQuery, address.HouseNumber, address.Street, address.City, address.District, address.Pincode, address.Landmark, UserID).Scan(&newAddress).Error

		return newAddress, err
	}
}

func (c *userDatabase) UpdateAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error) {
	var updated domain.Address

	updateQuery := `UPDATE addresses SET
								house_number = $1, street = $2, city = $3, district = $4, pincode = $5, landmark = $6
								WHERE user_id = $7
								RETURNING *`
	err := c.DB.Raw(updateQuery, address.HouseNumber, address.Street, address.City, address.District, address.Pincode, address.Landmark, UserID).Scan(&updated).Error
	return updated, err

}

func (c *userDatabase) VeiwAdress(ctx context.Context, UserID int) (domain.Address, error) {
	var Adress domain.Address

	quary := `SELECT *FROM addresses WHERE user_id=?`

	err := c.DB.Raw(quary, UserID).Scan(&Adress).Error
	return Adress, err
}

func (c *userDatabase) FindWishListItem(ctx context.Context, productID, userID uint) (domain.WishList, error) {

	var wishList domain.WishList
	query := `SELECT * FROM wish_lists WHERE user_id=? AND product_id=?`
	if c.DB.Raw(query, userID, productID).Scan(&wishList).Error != nil {
		return wishList, errors.New("faild to find wishlist item")
	}
	return wishList, nil
}

func (c *userDatabase) FindAllWishListItemsByUserID(ctx context.Context, userID uint) ([]response.Wishlist, error) {

	var wishLists []response.Wishlist

	favourite := ` SELECT *
	FROM products p
	JOIN wish_lists w ON w.product_id = p.id
	WHERE w.user_id = ?`

	if c.DB.Raw(favourite, userID).Scan(&wishLists).Error != nil {
		return wishLists, errors.New("faild to get wish_list items")
	}
	return wishLists, nil
}

func (c *userDatabase) SaveWishListItem(ctx context.Context, wishList domain.WishList) error {

	query := `INSERT INTO wish_lists (user_id, product_id) VALUES ($1, $2)`

	if c.DB.Raw(query, wishList.UserID, wishList.ProductID).Scan(&wishList).Error != nil {
		return errors.New("faild to insert a product into whishlist")
	}
	return nil
}

func (c *userDatabase) RemoveWishListItem(ctx context.Context, wishList domain.WishList) error {

	query := `DELETE FROM wish_lists WHERE id=?`
	if c.DB.Raw(query, wishList.ID).Scan(&wishList).Error != nil {
		return errors.New("faild to delete product")
	}
	return nil
}

func (c *userDatabase) FindProduct(ctx context.Context, id uint) (response.Product, error) {
	var product response.Product
	query := `SELECT p.id,p.product_name as name,p.description,p.brand,p.prize,p.category_id,p.qty_in_stock,c.category_name,p.created_at,p.updated_at FROM products p 
		JOIN categories c ON p.category_id=c.id WHERE p.id=$1`
	err := c.DB.Raw(query, id).Scan(&product).Error
	return product, err
}
