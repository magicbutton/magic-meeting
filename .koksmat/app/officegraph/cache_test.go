package officegraph

import (
	"context"
	"testing"

	"github.com/spf13/viper"
)

func TestSaveCache(t *testing.T) {
	connectionString := viper.GetString("POSTGRES_DB")
	conn, err := ConnectDB(connectionString)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	defer conn.Close()
	err = InsertIntoGraphCacheTable(context.Background(), conn, "test", "test", "test", "test", "test", "{}")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

}
