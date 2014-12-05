package devise

import (
	"github.com/jameskeane/bcrypt"
)

func ValidatePassword(password, encryptedPassword string, saltAndPepper ...string) bool {
	var salt, pepper string

	switch len(saltAndPepper) {
	case 0:
		salt = encryptedPassword[0:29]
		pepper = ""
	case 1:
		salt = saltAndPepper[0]
		pepper = ""
	default:
		salt = saltAndPepper[0]
		pepper = saltAndPepper[1]
	}

	hash, _ := bcrypt.Hash(password+pepper, salt)
	return hash == encryptedPassword
}
