package interactor

import "golang.org/x/crypto/bcrypt"

func genEncryptedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareHashAndPassword(encryptedPassword string, Password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(Password)); err != nil {
		return err
	}
	return nil
}
