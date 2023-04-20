package sqlstore_test

import (
	"fmt"
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		"localhost",
		5432,
		"bemmanue",
		"password",
		"store",
		"disable",
	)

	println(databaseURL)

	os.Exit(m.Run())
}
