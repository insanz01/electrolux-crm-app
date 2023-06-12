package dto

type GiftClaimById struct {
	Id string `param:"id"`
}

type GiftClaimRequest struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Key         string  `json:"key"`
	Value       string  `json:"value"`
	Datatype    string  `json:"datatype"`
	IsMandatory bool    `json:"is_mandatory"`
	InputType   string  `json:"input_type"`
}

type GiftClaimUpdateRequest struct {
	GiftClaims []GiftClaimRequest `json:"gift_claim"`
}

type GiftClaimProperties struct {
	Properties []*string         `json:"properties"`
	Filters    []*CustomerFilter `json:"filters"`
}

type SearchGiftProperties struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// type SearchGift struct {
// 	IDClaim     string `json:"id_claim"`
// 	PhoneNumber string `json:"phone_number"`
// }

type SearchGift struct {
	KeyValue []*SearchGiftProperties `json:"properties"`
}
