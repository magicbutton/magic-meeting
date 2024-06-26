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
	"github.com/magicbutton/magic-meeting/services/models/buildingmodel"
   
)


func MapBuildingOutgoing(db database.Building) buildingmodel.Building {
    return buildingmodel.Building{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Address : db.Address,
                Site_id : db.Site_id,

    }
}

func MapBuildingIncoming(in buildingmodel.Building) database.Building {
    return database.Building{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Address : in.Address,
                Site_id : in.Site_id,
        Searchindex : in.Name,

    }
}
