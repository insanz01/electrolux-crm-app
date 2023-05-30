package dto

type CustomerById struct {
	Id string `param:"id"`
}

type CustomerRequest struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Key         string  `json:"key"`
	Value       string  `json:"value"`
	Datatype    string  `json:"datatype"`
	IsMandatory bool    `json:"is_mandatory"`
	InputType   string  `json:"input_type"`
}

type CustomerUpdateRequest struct {
	Customers []CustomerRequest `json:"customer"`
}
