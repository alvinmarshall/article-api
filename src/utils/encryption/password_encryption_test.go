package encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	asserts := assert.New(t)
	type args struct {
		password string
	}
	expected, err := HashPassword(args{"any password"}.password)
	asserts.Nil(err)
	asserts.NotNil(expected)
}

func TestVerifyHashPassword(t *testing.T) {
	asserts := assert.New(t)
	type args struct {
		hashPassword string
		password     string
	}
	hashPassword, _ := HashPassword("any password")
	input := args{string(hashPassword), "any password"}
	expected, err := VerifyHashPassword(input.hashPassword, input.password)
	asserts.Nil(err)
	asserts.Equal(expected, true)
}
