/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package transactionmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)

func UnmarshalTransaction(data []byte) (Transaction, error) {
	var r Transaction
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Transaction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Transaction struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Amount int `json:"amount"`
    Currency string `json:"currency"`
    Datetime time.Time `json:"datetime"`
    Status string `json:"status"`
    Account_id int `json:"account_id"`

}

