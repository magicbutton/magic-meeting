/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package visitormodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)

func UnmarshalVisitor(data []byte) (Visitor, error) {
	var r Visitor
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Visitor) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Visitor struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Email string `json:"email"`
    Phone string `json:"phone"`
    Company string `json:"company"`
    Purpose string `json:"purpose"`
    Host_id int `json:"host_id"`
    Status string `json:"status"`

}

