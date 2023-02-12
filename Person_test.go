package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSavePerson_happyPath(t *testing.T) {
	in := &Person{
		Name:  "John",
		Phone: "123456789",
	}
	mockNFS := &mockSaver{}
	mockNFS.On("Save", mock.Anything).Return(nil).Once()
	resultErr := SavePerson(in, mockNFS)
	assert.NoError(t, resultErr)
	assert.True(t, mockNFS.AssertExpectations(t))
}

func TestSavePerson_nfsAlwaysFails(t *testing.T) {
	in := &Person{
		Name:  "John",
		Phone: "123456789",
	}
	mockNFS := &mockSaver{}
	mockNFS.On("Save", mock.Anything).Return(errors.New("save failed")).Once()
	resultErr := SavePerson(in, mockNFS)
	assert.Error(t, resultErr)
	assert.True(t, mockNFS.AssertExpectations(t))
}
