package domain

const (
	TypeName           string = "JWT"
	UIDKey             string = "uid"
	IdTokenSub         string = "id_token"
	RefreshTokenSub    string = "refresh_token"
	IdTokenExpSec      int64  = 3600
	RefreshTokenExpSec int64  = 864000
)
