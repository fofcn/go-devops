package security

import "golang.org/x/crypto/bcrypt"

type PasswordEncoder interface {
	Encode(raw string) string
	Matches(raw string, encoded string) bool
}

type bcryptpasswordencoder struct {
}

func NewBcryptPasswordEncoder() PasswordEncoder {
	return &bcryptpasswordencoder{}
}

func (bcryptpasswordencoder) Encode(rawPass string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (bcryptpasswordencoder) Matches(raw string, encoded string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(raw), []byte(encoded))
	return err != nil
}
