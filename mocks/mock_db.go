package mocks

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

// MockDB is a mock struct for GORM's DB methods
type MockDB struct {
	mock.Mock
}

// Mock the First method of GORM to simulate database queries
func (m *MockDB) First(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

// Mock the Save method of GORM
func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

// Mock the Create method of GORM
func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

// Implement any other GORM methods you need for testing
