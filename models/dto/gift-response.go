package dto

type GiftClaim struct {
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

type GroupGiftClaim struct {
	GroupId       string      `json:"gift_claim_id"`
	GiftClaimData []GiftClaim `json:"attributes"`
}

type GiftClaimResponse struct {
	GiftClaim interface{} `json:"gift_claims"`
}
