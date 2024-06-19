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
	"github.com/magicbutton/magic-meeting/services/models/serviceordermodel"
   
)


func MapServiceOrderOutgoing(db database.ServiceOrder) serviceordermodel.ServiceOrder {
    return serviceordermodel.ServiceOrder{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                Deliverto_id : db.Deliverto_id,
        Status : db.Status,
                Payment_id : db.Payment_id,

    }
}

func MapServiceOrderIncoming(in serviceordermodel.ServiceOrder) database.ServiceOrder {
    return database.ServiceOrder{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                Deliverto_id : in.Deliverto_id,
        Status : in.Status,
                Payment_id : in.Payment_id,
        Searchindex : in.Name,

    }
}
