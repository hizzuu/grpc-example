package model

type contextKey string

const (
	// context
	CtxJwtTokenKey     contextKey = "jwtToken"
	CtxClaimsKey       contextKey = "jwtClaims"
	CtxAuthErrorCtxKey contextKey = "authError"

	// token
	TypeName           string = "JWT"
	ClaimsKey          string = "claims"
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
