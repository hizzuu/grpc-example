package conf

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strconv"
)

type appConfig struct {
	Debug bool
}

type apiConfig struct {
	Addr string
}

type credentialsConfig struct {
	PrivateKey *rsa.PrivateKey
}

var (
	AppConf         appConfig
	ApiConf         apiConfig
	CredentialsConf credentialsConfig
)

func init() {
	initAppConf()
	initApiConf()
	initCredentialsConf()
}

func initAppConf() {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		AppConf.Debug = false
		return
	}
	AppConf.Debug = debug
}

func initApiConf() {
	ApiConf.Addr = "8080"
}

func initCredentialsConf() {

	block, _ := pem.Decode([]byte(os.Getenv("PRIVATE_KEY")))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse private key: %s", err))
	}
	CredentialsConf.PrivateKey = key
}
