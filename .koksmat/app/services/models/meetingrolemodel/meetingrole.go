/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package meetingrolemodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)

func UnmarshalMeetingRole(data []byte) (MeetingRole, error) {
	var r MeetingRole
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MeetingRole) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MeetingRole struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    User_id int `json:"user_id"`
    Meeting_id int `json:"meeting_id"`
    Role string `json:"role"`

}

