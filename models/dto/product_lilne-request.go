package dto

type ProductLineInsertRequest struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}

type ProductLineUpdateRequest struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}
