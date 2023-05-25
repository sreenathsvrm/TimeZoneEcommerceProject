package requests

type Usersign struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OTPreq struct {
	Phone string `json:"phoneNumber,omitempty" validate:"required"`
}
type Otpverifier struct {
	Pin   string `json:"pin,omitempty" validate:"required"`
	Phone string `json:"phoneNumber,omitempty" validate:"required"`
}

type Pagination struct {
	Page      uint `json:"page"`
	PerPage uint `json:"page_per"`
}

type BlockUser struct {
	UserID int    `json:"user_id"`
	Reason string `json:"reason"`
}

type AddressReq struct {
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
}