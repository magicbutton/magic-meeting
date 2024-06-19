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
	"github.com/magicbutton/magic-meeting/services/models/servicemodel"
   
)


func MapServiceOutgoing(db database.Service) servicemodel.Service {
    return servicemodel.Service{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Currency : db.Currency,
                Servicecategory_id : db.Servicecategory_id,

    }
}

func MapServiceIncoming(in servicemodel.Service) database.Service {
    return database.Service{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Currency : in.Currency,
                Servicecategory_id : in.Servicecategory_id,
        Searchindex : in.Name,

    }
}
