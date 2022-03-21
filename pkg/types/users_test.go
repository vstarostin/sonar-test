package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultUsers(t *testing.T) {
	expectedUsers := []User{
		{Name: "User1"},
		{Name: "User2"},
		{Name: "User3"},
		{Name: "User4"},
	}
	assert.Equal(t, expectedUsers, DefaultUsers())
}
