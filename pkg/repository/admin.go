package repository

import (
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type AdminDB struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDB{
		DB: DB,
	}
}

func (c *AdminDB) FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	if c.DB.Raw("SELECT * FROM admins WHERE email=? OR admin_name=?", admin.Email, admin.AdminName).Scan(&admin).Error != nil {
		return admin, errors.New("faild to find admin")
	}
	return admin, nil
}

func (c *AdminDB) SaveAdmin(ctx context.Context, admin domain.Admin) error {
	Query := `	INSERT INTO admins(admin_name, email, password)
 VALUES($1, $2, $3) RETURNING *;`

	if c.DB.Exec(Query, admin.AdminName, admin.Email, admin.Password).Error != nil {
		return errors.New("failed to create admin")
	}
	return nil
}

func (c *AdminDB) FindAllUser(ctx context.Context, pagination urequest.Pagination) (users []response.UserValue, err error) {

	limit := pagination.Page
	offset := (pagination.PerPage - 1) * limit

	query := `SELECT * FROM users ORDER BY name DESC LIMIT $1 OFFSET $2`

	err = c.DB.Raw(query, limit, offset).Scan(&users).Error

	return users, err
}

func (c *AdminDB) BlockUser(body urequest.BlockUser, AdminId int) error {
	// Start a transaction
	tx := c.DB.Begin()
	//Check if the user is there
	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", body.UserID).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such user")
	}

	// Execute the first SQL command (UPDATE)
	if err := tx.Exec("UPDATE users SET is_blocked = true WHERE id = ?", body.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Execute the second SQL command (INSERT)
	if err := tx.Exec("INSERT INTO user_statuses  (users_id, reason_for_blocking, blocked_at, blocked_by) VALUES (?, ?, NOW(), ?)", body.UserID, body.Reason, AdminId).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	// If all commands were executed successfully, return nil
	return nil

}
func (c *AdminDB) UnblockUser(id int) error {
	tx := c.DB.Begin()

	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND is_blocked=true)", id).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such user to unblock")
	}
	if err := tx.Exec("UPDATE users SET is_blocked = false WHERE id=$1", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	query := "UPDATE user_statuses SET reason_for_blocking=$1,blocked_at=NULL,blocked_by=$2 WHERE users_id=$3"
	if err := tx.Exec(query, "", 0, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *AdminDB) FindUserbyId(ctx context.Context, userID int) (domain.Users, error) {
	var user domain.Users

	FindbyID := `SELECT *FROM users WHERE id= $1;`

	err := c.DB.Raw(FindbyID, userID).Scan(&user).Error

	if user.ID == 0 {
		return domain.Users{}, fmt.Errorf("no user found")
	}
	return user, err
}
