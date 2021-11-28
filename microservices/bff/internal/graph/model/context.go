package model

type contextKey string

const (
	CtxJwtTokenKey     contextKey = "jwtToken"
	CtxClaimsKey       contextKey = "jwtClaims"
	CtxAuthErrorCtxKey contextKey = "authError"
)
