package devise

import (
	"testing"
)

func TestValidPassword(t *testing.T) {
	encrypted := "$2a$10$GZEYTPPJZjh.tnBW0/d9U.pI4j7iyQPIzz.NQD9GvVWxPkb6nq2ZO"
	password := "123456"

	if ValidatePassword(password, encrypted) != true {
		t.Errorf("Password should be valid.")
	}
}

func TestInvalidPassword(t *testing.T) {
	encrypted := "$2a$10$vSRc95cy5AfO/OLVW2tZXe9D4fT83Mhwlpq7IIJoMmfqFKbMDL4OW"
	password := "just guess"

	if ValidatePassword(password, encrypted) == true {
		t.Errorf("Password should be invalid.")
	}
}
