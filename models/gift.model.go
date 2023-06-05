package models

import "time"

type GiftClaimProperty struct {
}

type GiftClaim struct {
	Id              int64             `json:"id" db:"id"`
	CustomerId      int64             `json:"customer_id" db:"customer_id"`
	ClaimUniqueCode string            `json:"claim_unique_code" db:"claim_unique_code"`
	Property        GiftClaimProperty `json:"property" db:"propery"`
	ClaimDate       *time.Time        `json:"claim_date" db:"claim_date"`
	Status          string            `json:"status" db:"status"`
}

type GiftProperties struct {
	ID          string  `db:"id" json:"id"`
	TableDataID string  `db:"table_data_id" json:"table_data_id"`
	OrderNumber *int    `db:"order_number" json:"order_number"`
	Name        *string `db:"name" json:"name"`
	Key         string  `db:"key" json:"key"`
	Value       string  `db:"value" json:"value"`
	Datatype    string  `db:"datatype" json:"datatype"`
	IsMandatory bool    `db:"is_mandatory" json:"is_mandatory"`
	InputType   string  `db:"input_type" json:"input_type"`
	UpdatedAt   string  `db:"updated_at" json:"updated_at"`
}

type Gifts struct {
	Gift []GiftProperties `json:"gift_claim"`
}
