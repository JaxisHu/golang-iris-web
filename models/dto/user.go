package dto

import (
	. "IRIS_WEB/errors"
	"IRIS_WEB/utility/validator"
	"github.com/kataras/iris"
	"time"
)

//  --------------- DTO For Params ---------------
type UserParamDTO struct {
	UserId int `form:"user_id" json:"user_id" validate:"min=1"`
	//UserEmail string `form:"user_email" json:"user_email" validate:"email"`
	//UserPhone string `form:"user_phone" json:"user_phone" validate:"phone"`
}

// /path?user_id=1
func (u *UserParamDTO) Bind(ctx iris.Context) error {
	if err := ctx.ReadForm(u); err != nil {
		return ParamError("invalid form format")
	}

	if err, errMsg := validator.Check(u); err != nil {
		return ParamError(err, errMsg)
	}

	return nil
}

//  --------------- DTO ---------------
type UserDTO struct {
	ID                 int       `gorm:"column:id" json:"id"`
	UserName           string    `gorm:"column:user_name" json:"user_name"`
	AuthKey            string    `gorm:"column:auth_key" json:"auth_key"`
	PasswordHash       string    `gorm:"column:password_hash" json:"password_hash"`
	PasswordResetToken string    `gorm:"column:password_reset_token" json:"password_reset_token"`
	Email              string    `gorm:"column:email" json:"email"`
	Status             int       `gorm:"column:status" json:"status"`
	VerificationToken  string    `gorm:"column:verification_token" json:"verification_token"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (UserDTO) TableName() string {
	return "user"
}
