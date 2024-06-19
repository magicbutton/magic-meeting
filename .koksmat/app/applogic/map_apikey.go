/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateMapModel v1.1
package applogic
import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-meeting/database"
	"github.com/magicbutton/magic-meeting/services/models/apikeymodel"
   
)


func MapAPIKeyOutgoing(db database.APIKey) apikeymodel.APIKey {
    return apikeymodel.APIKey{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Key : db.Key,
                User_id : db.User_id,

    }
}

func MapAPIKeyIncoming(in apikeymodel.APIKey) database.APIKey {
    return database.APIKey{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Key : in.Key,
                User_id : in.User_id,
        Searchindex : in.Name,

    }
}
