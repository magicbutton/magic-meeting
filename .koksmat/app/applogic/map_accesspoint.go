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
	"github.com/magicbutton/magic-meeting/services/models/accesspointmodel"
   
)


func MapAccessPointOutgoing(db database.AccessPoint) accesspointmodel.AccessPoint {
    return accesspointmodel.AccessPoint{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                Location_id : db.Location_id,
        Status : db.Status,

    }
}

func MapAccessPointIncoming(in accesspointmodel.AccessPoint) database.AccessPoint {
    return database.AccessPoint{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                Location_id : in.Location_id,
        Status : in.Status,
        Searchindex : in.Name,

    }
}
