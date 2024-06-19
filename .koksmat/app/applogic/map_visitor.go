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
	"github.com/magicbutton/magic-meeting/services/models/visitormodel"
   
)


func MapVisitorOutgoing(db database.Visitor) visitormodel.Visitor {
    return visitormodel.Visitor{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Email : db.Email,
        Phone : db.Phone,
        Company : db.Company,
        Purpose : db.Purpose,
                Host_id : db.Host_id,
        Status : db.Status,

    }
}

func MapVisitorIncoming(in visitormodel.Visitor) database.Visitor {
    return database.Visitor{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Email : in.Email,
        Phone : in.Phone,
        Company : in.Company,
        Purpose : in.Purpose,
                Host_id : in.Host_id,
        Status : in.Status,
        Searchindex : in.Name,

    }
}
