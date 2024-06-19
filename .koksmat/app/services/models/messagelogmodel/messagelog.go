/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package messagelogmodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalMessageLog(data []byte) (MessageLog, error) {
	var r MessageLog
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessageLog) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageLog struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Sender string `json:"sender"`
    Receiver string `json:"receiver"`
    Message string `json:"message"`

}

