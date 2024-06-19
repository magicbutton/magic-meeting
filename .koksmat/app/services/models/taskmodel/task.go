/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package taskmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)

func UnmarshalTask(data []byte) (Task, error) {
	var r Task
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Task) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Task struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Starttime time.Time `json:"starttime"`
    Location string `json:"location"`
    Responsible_id int `json:"responsible_id"`

}

