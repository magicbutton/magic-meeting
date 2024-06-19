/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package meetingroommodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)

func UnmarshalMeetingRoom(data []byte) (MeetingRoom, error) {
	var r MeetingRoom
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MeetingRoom) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MeetingRoom struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Email string `json:"email"`
    Capacity int `json:"capacity"`
    Features string `json:"features"`
    Floor_id int `json:"floor_id"`

}

