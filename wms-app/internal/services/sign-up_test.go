package services

import (
	"testing"
	"wms-app/internal/models/dbModels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDB simulates the DB Create method
// You can expand this for more DB operations

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(user *dbModels.User) *MockResult {
	args := m.Called(user)
	return args.Get(0).(*MockResult)
}

type MockResult struct {
	Error error
}

func TestCreateUser_Success(t *testing.T) {
	mockDB := new(MockDB)
	user := &dbModels.User{
		UserId:       "123",
		FirstName:    "John",
		MiddleName:   "A.",
		LastName:     "Doe",
		Email:        "john@example.com",
		MobileNumber: "1234567890",
	}
	mockResult := &MockResult{Error: nil}
	mockDB.On("Create", user).Return(mockResult)

	// Replace config.DB with mockDB if possible, or refactor CreateUser to accept DB as param
	resp, err := CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, user.UserId, resp.UserId)
	assert.Equal(t, user.FirstName, resp.FirstName)
	assert.Equal(t, user.Email, resp.Email)
}

func TestCreateUser_DBError(t *testing.T) {
	mockDB := new(MockDB)
	user := &dbModels.User{UserId: "123"}
	mockResult := &MockResult{Error: assert.AnError}
	mockDB.On("Create", user).Return(mockResult)

	resp, err := CreateUser(user)
	assert.Error(t, err)
	assert.Equal(t, "", resp.UserId)
}
