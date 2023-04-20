package teststore_test

import (
	"github.com/bemmanue/wildberries_L0/internal/model"
	"github.com/bemmanue/wildberries_L0/internal/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderRepository_Create(t *testing.T) {
	s := teststore.New()
	o := model.TestOrderJSON(t)
	assert.NoError(t, s.Order().Create(o))
	assert.NotNil(t, o)
}

func TestOrderRepository_FindAll(t *testing.T) {
	s := teststore.New()
	o1 := model.TestOrderJSON(t)

	s.Order().Create(o1)

	o2, err := s.Order().FindAll()
	assert.NoError(t, err)
	assert.NotNil(t, o2)
}
