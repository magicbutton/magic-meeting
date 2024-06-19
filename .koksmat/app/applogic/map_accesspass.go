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
	"github.com/magicbutton/magic-meeting/services/models/accesspassmodel"
   
)


func MapAccessPassOutgoing(db database.AccessPass) accesspassmodel.AccessPass {
    return accesspassmodel.AccessPass{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                Visitor_id : db.Visitor_id,
        Status : db.Status,

    }
}

func MapAccessPassIncoming(in accesspassmodel.AccessPass) database.AccessPass {
    return database.AccessPass{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                Visitor_id : in.Visitor_id,
        Status : in.Status,
        Searchindex : in.Name,

    }
}
