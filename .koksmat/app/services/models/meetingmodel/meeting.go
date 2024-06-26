/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package meetingmodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalMeeting(data []byte) (Meeting, error) {
	var r Meeting
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Meeting) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Meeting struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Start time.Time `json:"start"`
    Duration int `json:"duration"`
    Location string `json:"location"`
    Status string `json:"status"`
    Exchangereference string `json:"exchangereference"`
    Exchangestatus string `json:"exchangestatus"`
    Teamsreference string `json:"teamsreference"`
    Teamsstatus string `json:"teamsstatus"`

}

