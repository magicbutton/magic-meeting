package officegraph

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

func UploadGraphData() {
}

const batchSize = 10000

type Record struct {
	// Define the structure based on the view's output
	// Example fields:

	Data json.RawMessage `json:"data"`
}

// ConnectDB initializes a connection to the database using a connection string
func ConnectDB(connStr string) (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}

// InsertIntoImportTable inserts JSON data into the import table
func InsertIntoGraphCacheTable(ctx context.Context, db *sql.DB, batchName, name string, actor string, token string, url string, jsonData string) error {
	query := `
        INSERT INTO public.graphcache (
	id, created_at, created_by, updated_at, updated_by, deleted_at, tenant, searchindex, name, description, upn, token, url, data
          
        ) VALUES (
            DEFAULT,DEFAULT,$1, DEFAULT,$1, DEFAULT,
						'', -- tenant
						 '', -- searchindex						 
						 $2, -- name
						 
						 '',  -- description
						 $2,  -- upn
						 $3, -- token
						 $4, -- url
						 $5  -- data
					
        )`

	_, err := db.ExecContext(ctx, query, actor, name, token, url, jsonData)
	if err != nil {
		return fmt.Errorf("failed to insert into graphcache table: %w", err)
	}
	return nil
}
