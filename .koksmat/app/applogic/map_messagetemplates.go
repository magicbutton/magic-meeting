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
	"github.com/magicbutton/magic-meeting/services/models/messagetemplatesmodel"
   
)


func MapMessageTemplateOutgoing(db database.MessageTemplate) messagetemplatesmodel.MessageTemplate {
    return messagetemplatesmodel.MessageTemplate{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Type : db.Type,
        Content : db.Content,

    }
}

func MapMessageTemplateIncoming(in messagetemplatesmodel.MessageTemplate) database.MessageTemplate {
    return database.MessageTemplate{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Type : in.Type,
        Content : in.Content,
        Searchindex : in.Name,

    }
}
