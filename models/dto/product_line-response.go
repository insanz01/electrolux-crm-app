package dto

type ProductLine struct {
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

type GroupProductLine struct {
	GroupId         string        `json:"product_line_id"`
	ProductLineData []ProductLine `json:"attributes"`
}

type ProductLineResponse struct {
	ProductLine interface{} `json:"product_lines"`
}
