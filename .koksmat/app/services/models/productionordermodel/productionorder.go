/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package productionordermodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)

func UnmarshalProductionOrder(data []byte) (ProductionOrder, error) {
	var r ProductionOrder
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProductionOrder) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ProductionOrder struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Deliverydate time.Time `json:"deliverydate"`
    Deliverto_id int `json:"deliverto_id"`
    Status string `json:"status"`
    Payment_id int `json:"payment_id"`

}

