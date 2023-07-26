package models

type DeleteTableData struct {
	Id string `db:"id"`
}

type TableProperty struct {
	ID          string `db:"id"`
	TableDataID string `db:"table_data_id"`
	OrderNumber int    `db:"order_number"`
	Name        string `db:"name"`
	Key         string `db:"key"`
	Value       string `db:"value"`
	Datatype    string `db:"datatype"`
	IsMandatory bool   `db:"is_mandatory"`
	InputType   string `db:"input_type"`
}

type TableCategory struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type TableData struct {
	TableId string `db:"table_id"`
}

type Pagination struct {
	Page  int
	Limit int
}

type KeyValue struct {
	Key   string `db:"key:"`
	Value string `db:"value"`
}

type CampaignJSONParameter struct {
	Type   string `db:"type"`
	Number int    `db:"number"`
	Value  string `db:"value"`
}
