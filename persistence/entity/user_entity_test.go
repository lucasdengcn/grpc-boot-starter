package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserEntity(t *testing.T) {
	ue := User{}
	assert.Nil(t, ue.BirthDay)
	assert.Equal(t, uint(0), ue.ID)
	assert.Empty(t, ue.Name)
	assert.Empty(t, ue.Password)
	assert.Empty(t, ue.Gender)
	//
	assert.NotNil(t, ue.CreatedAt)
	assert.NotNil(t, ue.UpdatedAt)
}
