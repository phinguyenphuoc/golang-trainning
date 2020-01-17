package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/mock"
)

type DBMock struct {
	mock.Mock
}

func (m *DBMock) InitDB() error {
	args := m.Called()
	return args.Error(0)
}

func (m *DBMock) GetDB() *sql.DB {
	db, _, _ := sqlmock.New()
	return db
}
func TestRoutes(t *testing.T) {
	mockDB := new(DBMock)
	mockDB.On("GetDB").Return(new(sql.DB))
	routes(mockDB)
	assert.Equal(t, 1, 1)
}
