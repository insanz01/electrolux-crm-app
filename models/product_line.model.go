package models

type ProductLineProperties struct {
	ID          string  `db:"id" json:"id"`
	TableDataID string  `db:"table_data_id" json:"table_data_id"`
	OrderNumber *int    `db:"order_number" json:"order_number"`
	Name        *string `db:"name" json:"name"`
	Key         string  `db:"key" json:"key"`
	Value       string  `db:"value" json:"value"`
	Datatype    string  `db:"datatype" json:"datatype"`
	IsMandatory bool    `db:"is_mandatory" json:"is_mandatory"`
	InputType   string  `db:"input_type" json:"input_type"`
}
