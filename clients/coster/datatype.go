package coster

import "time"

type Template struct {
	Meta struct {
		Version string `json:"version"`
		Author  string `json:"author"`
	} `json:"meta"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	AccountID string `json:"account_id"`
	Category  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Division struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"division"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Messages  []struct {
		ID          string `json:"id"`
		ReferenceID string `json:"reference_id"`
		FileID      string `json:"file_id"`
		CreatedAt   string `json:"created_at"`
		Language    struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"language"`
		ApprovedAt string `json:"approved_at"`
		RejectedAt any    `json:"rejected_at"`
		Data       struct {
			Category   string `json:"category"`
			Language   string `json:"language"`
			Components []struct {
				Type    string `json:"type"`
				Format  string `json:"format,omitempty"`
				Example struct {
					HeaderHandle []string `json:"header_handle"`
				} `json:"example,omitempty"`
				Text string `json:"text,omitempty"`
			} `json:"components"`
			IsDynamicURL        bool `json:"is_dynamic_url"`
			TotalParamBody      int  `json:"total_param_body"`
			TotalDynamicURL     int  `json:"total_dynamic_url"`
			IsButtonCallback    bool `json:"is_button_callback"`
			TotalParamButton    int  `json:"total_param_button"`
			TotalParamHeader    int  `json:"total_param_header"`
			TotalButtonCallback int  `json:"total_button_callback"`
		} `json:"data"`
	} `json:"messages"`
}
