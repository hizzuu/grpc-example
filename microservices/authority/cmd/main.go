package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

var privateKey *rsa.PrivateKey

func main() {
	token := jwt.New()
	token.Set(jwt.IssuerKey, "authority")
	token.Set(jwt.SubjectKey, "1") // sub = user id
	headers := jws.NewHeaders()
	headers.Set(jws.AlgorithmKey, jwa.RS256)
	headers.Set(jws.TypeKey, "JWT")
	b, err := json.Marshal(token)
	if err != nil {
		panic(err)
	}
	signedToken, err := jws.Sign(b, jwa.RS256, privateKey, jws.WithHeaders(headers))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(signedToken))
}

func init() {
	block, _ := pem.Decode([]byte(os.Getenv("PRIVATE_KEY")))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse private key: %s", err))
	}
	privateKey = key
}
