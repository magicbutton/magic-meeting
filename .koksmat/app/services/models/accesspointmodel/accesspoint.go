/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package accesspointmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)

func UnmarshalAccessPoint(data []byte) (AccessPoint, error) {
	var r AccessPoint
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccessPoint) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AccessPoint struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Location_id int `json:"location_id"`
    Status string `json:"status"`

}

