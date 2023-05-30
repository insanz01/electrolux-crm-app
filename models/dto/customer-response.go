package dto

// type CustomerLocation struct {
// 	City    string `json:"city"`
// 	State   string `json:"state"`
// 	Country string `json:"country"`
// }

// type CustomerResponse struct {
// 	ServiceOrderNo    string           `json:"service_order_no"`
// 	MobilePhoneNumber string           `json:"mobile_phone_number"`
// 	ProductLine       string           `json:"product_line"`
// 	DateOfPurchase    string           `json:"date_of_purchase"`
// 	Location          CustomerLocation `json:"location"`
// 	ServiceType       string           `json:"service_type"`
// 	ModelCode         string           `json:"model_code"`
// }

type Customer struct {
	ID          string  `json:"id"`
	TableDataID string  `json:"table_data_id"`
	OrderNumber *int    `json:"order_number"`
	Name        *string `json:"name"`
	Key         string  `json:"key"`
	Value       string  `json:"value"`
	Datatype    string  `json:"datatype"`
	IsMandatory bool    `json:"is_mandatory"`
	InputType   string  `json:"input_type"`
}

type GroupCustomer struct {
	GroupId      string     `json:"customer_id"`
	CustomerData []Customer `json:"attributes"`
}

type CustomerResponse struct {
	Customer interface{} `json:"customers"`
}
