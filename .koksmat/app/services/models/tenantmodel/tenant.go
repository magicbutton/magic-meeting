/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package tenantmodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalTenant(data []byte) (Tenant, error) {
	var r Tenant
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Tenant) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tenant struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Email string `json:"email"`

}

