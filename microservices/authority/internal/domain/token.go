package domain

type TokenType string

type Token struct {
	IDToken      TokenType
	RefreshToken TokenType
}
