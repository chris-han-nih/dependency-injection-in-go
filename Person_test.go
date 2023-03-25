package main

import (
	"errors"
	"fmt"
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
	mockNFS.On("Save", mock.Anything).Return(1, nil).Once()
	_, resultErr := SavePerson(in, mockNFS)
	assert.NoError(t, resultErr)
	assert.True(t, mockNFS.AssertExpectations(t))
}

type mockSaver struct {
	mock.Mock
}

func (m *mockSaver) Save(data []byte) (int, error) {
	outputs := m.Mock.Called(data)

	fmt.Println("outputs: ", outputs)

	return outputs.Int(0), outputs.Error(1)
}

func TestSavePerson_nfsAlwaysFails(t *testing.T) {
	in := &Person{
		Name:  "John",
		Phone: "123456789",
	}
	mockNFS := &mockSaver{}
	mockNFS.On("Save", mock.Anything).Return(0, errors.New("save failed")).Once()
	_, resultErr := SavePerson(in, mockNFS)
	assert.Error(t, resultErr)
	assert.True(t, mockNFS.AssertExpectations(t))
}
