package middleware

type contextkey string

// AuthTokenKey is the key to get token detail from echo context
const AuthTokenKey contextkey = "git-rbi.jatismobile.com/jatis_chatcommerce/open-api-inventory/auth/auth.TokenDetail"

// TokenDetail is the field contained in token
type TokenDetail struct {
	AccessToken  string `json:"access_token"`
	ClientID     string `json:"client_id"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	User         struct {
		ID           string   `json:"id"`
		Name         string   `json:"name"`
		Email        string   `json:"email"`
		Phone        any      `json:"phone"`
		Roles        []string `json:"roles"`
		Title        any      `json:"title"`
		Avatar       any      `json:"avatar"`
		Mobile       any      `json:"mobile"`
		Modules      []string `json:"modules"`
		ClientID     string   `json:"client_id"`
		Divisions    []string `json:"divisions"`
		LastName     any      `json:"last_name"`
		FirstName    any      `json:"first_name"`
		Applications []string `json:"applications"`
		DisplayName  string   `json:"display_name"`
	} `json:"user"`
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
