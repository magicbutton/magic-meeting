/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v3.api
package magicmeeting
import (

	"time"
    // "github.com/magicbutton/magic-meeting/database/databasetypes"
)


type APIKey struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Key string `json:"key"`
    User_id int `json:"user_id"`
    Validto time.Time `json:"validto"`

}

