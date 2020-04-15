package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	user := &User{
		Id:        "asdf",
		FirstName: "",
	}
	assert.Equal(t, user.IsValid(), InvalidUserError("first_name", user.FirstName))
}
