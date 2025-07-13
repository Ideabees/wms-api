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

type MockResult struct {
	err error
}

func (m *MockDB) Create(value interface{}) DBResult {
	args := m.Called(value)
	return args.Get(0).(DBResult)
}

func (r *MockResult) Error() error {
	return r.err
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
	mockResult := &MockResult{err: nil}
	mockDB.On("Create", user).Return(mockResult)

	resp, err := CreateUserWithDB(mockDB, user)

	assert.NoError(t, err)
	assert.Equal(t, user.UserId, resp.UserId)
	assert.Equal(t, user.FirstName, resp.FirstName)
	assert.Equal(t, user.Email, resp.Email)
}

func TestCreateUser_DBError(t *testing.T) {
	mockDB := new(MockDB)
	user := &dbModels.User{UserId: "123"}
	mockResult := &MockResult{err: assert.AnError}
	mockDB.On("Create", user).Return(mockResult)

	resp, err := CreateUserWithDB(mockDB, user)
	assert.Error(t, err)
	assert.Equal(t, "", resp.UserId)
}
