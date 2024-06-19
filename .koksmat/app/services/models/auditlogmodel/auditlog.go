/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package auditlogmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)

func UnmarshalAuditLog(data []byte) (AuditLog, error) {
	var r AuditLog
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuditLog) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuditLog struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Action string `json:"action"`
    User_id int `json:"user_id"`
    Entity string `json:"entity"`
    Entityid string `json:"entityid"`
    Timestamp time.Time `json:"timestamp"`

}

