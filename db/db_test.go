package db

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInitDBSuccess(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.Equal(t, nil, err)
	dbi := CreateDB()
	err = dbi.InitDB()
	assert.Equal(t, nil, err)
}

func TestGetDB(t *testing.T) {
	dbnull := CreateDB()
	err := dbnull.InitDB()
	db := dbnull.GetDB()
	assert.NotEqual(t, nil, dbnull)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, db)
}
