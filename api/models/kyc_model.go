package models

import (
	"github.com/go-playground/validator/v10"
)

type KYC struct {
	Wallet_PK    string `db:"wallet_pk" json:"wallet_pk" validate:"required"`
	First_Name   string `db:"first_name" json:"first_name" validate:"required"`
	Last_Name    string `db:"last_name" json:"last_name" validate:"required"`
	Phone_Number string `db:"phone_num" json:"phone_num" validate:"required,e164"`
	Email        string `db:"email" json:"email" validate:"required,email"`
	Ship_Addr_A  string `db:"ship_addr_a" json:"ship_addr_a" validate:"required"`
	Ship_Addr_B  string `db:"ship_addr_b" json:"ship_addr_b"`
	Ship_City    string `db:"ship_city" json:"ship_city" validate:"required"`
	Ship_State   string `db:"ship_state" json:"ship_state" validate:"required"`
	Ship_ZIP     int32  `db:"ship_zip" json:"ship_zip" validate:"required"`
	Dob_Day      int32  `db:"dob_day" json:"dob_day" validate:"required,gte=1,lte=31"`
	Dob_Month    int32  `db:"dob_month" json:"dob_month" validate:"required,gte=1,lte=12"`
	Dob_Year     int32  `db:"dob_year" json:"dob_year" validate:"required,gte=1900"`
	Title        string `db:"title" json:"title"`
}

func (kyc *KYC) Validate() error {
	validate := validator.New()
	return validate.Struct(kyc)
}
