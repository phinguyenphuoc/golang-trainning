package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInitDBFail(t *testing.T) {
	dbi := CreateDB()
	os.Setenv("MYSQL_DATABASE", "oracle")
	err := dbi.InitDB()
	assert.NotEqual(t, nil, err)

}

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
