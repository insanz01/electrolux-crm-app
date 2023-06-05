package models

import "time"

type CustomerProperty struct {
}

type CustomerLocation struct {
	City    string `db:"city" json:"city"`
	State   string `db:"state" json:"state"`
	Country string `db:"country" json:"country"`
}

type CustomerInsert struct {
	Id               int              `json:"id" db:"id"`
	ServiceOrderNo   string           `json:"service_order_no" db:"service_order_no"`
	ProductLine      string           `json:"product_line" db:"product_line"`
	PhoneNumber      string           `json:"phone_number" db:"phone_number"`
	DateOfPurchase   string           `json:"date_of_purchase" db:"date_of_purchase"`
	CustomerLocation CustomerLocation `json:"location"`
	ServiceType      string           `json:"service_type" db:"service_type"`
	ModelCode        string           `json:"model_code" db:"model_code"`
	Property         CustomerProperty `json:"property" db:"property"`
}

type Customer struct {
	ServiceOrderNo    string           `db:"service_order_no" json:"service_order_no"`
	MobilePhoneNumber string           `db:"mobile_phone_number" json:"mobile_phone_number"`
	ProductLine       string           `db:"product_line" json:"product_line"`
	DateOfPurchase    string           `db:"date_of_purchase" json:"date_of_purchase"`
	Location          CustomerLocation `db:"location" json:"location"`
	ServiceType       string           `db:"service_type" json:"service_type"`
	ModelCode         string           `db:"model_code" json:"model_code"`
}

type CustomerProperties struct {
	ID          string     `db:"id" json:"id"`
	TableDataID string     `db:"table_data_id" json:"table_data_id"`
	OrderNumber *int       `db:"order_number" json:"order_number"`
	Name        *string    `db:"name" json:"name"`
	Key         string     `db:"key" json:"key"`
	Value       string     `db:"value" json:"value"`
	Datatype    string     `db:"datatype" json:"datatype"`
	IsMandatory bool       `db:"is_mandatory" json:"is_mandatory"`
	InputType   string     `db:"input_type" json:"input_type"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

type Customers struct {
	Customer []CustomerProperties `json:"customer"`
}
