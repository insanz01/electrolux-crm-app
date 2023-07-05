package middleware

type contextkey string

// AuthTokenKey is the key to get token detail from echo context
const AuthTokenKey contextkey = "git-rbi.jatismobile.com/jatis_chatcommerce/open-api-inventory/auth/auth.TokenDetail"

// TokenDetail is the field contained in token
type TokenDetail struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Roles        []string `json:"roles"`
	Title        string   `json:"title"`
	Avatar       string   `json:"avatar"`
	Mobile       string   `json:"mobile"`
	Modules      []string `json:"modules"`
	ClientID     string   `json:"client_id"`
	Divisions    []string `json:"divisions"`
	LastName     string   `json:"last_name"`
	FirstName    string   `json:"first_name"`
	Applications []string `json:"applications"`
	DisplayName  string   `json:"display_name"`
}

// CheckTokenResponse is the data structure returned from auth service in case of a valid token
type CheckTokenResponse struct {
	Status string       `json:"status"`
	Data   *TokenDetail `json:"data"`
}

// CheckTokenErrorResponse is the data structure returned from auth service in case of an invalid token or any error happen
type CheckTokenErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
