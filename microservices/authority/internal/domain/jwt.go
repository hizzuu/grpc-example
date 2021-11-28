package domain

const (
	TypeName           string = "JWT"
	ClaimsKey          string = "jwtClaims"
	IdTokenSub         string = "id_token"
	RefreshTokenSub    string = "refresh_token"
	IdTokenExpSec      int64  = 3600
	RefreshTokenExpSec int64  = 864000
)

type JwtClaims struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
}
