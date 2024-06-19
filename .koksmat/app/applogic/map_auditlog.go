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
	"github.com/magicbutton/magic-meeting/services/models/auditlogmodel"
   
)


func MapAuditLogOutgoing(db database.AuditLog) auditlogmodel.AuditLog {
    return auditlogmodel.AuditLog{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Action : db.Action,
                User_id : db.User_id,
        Entity : db.Entity,
        Entityid : db.Entityid,

    }
}

func MapAuditLogIncoming(in auditlogmodel.AuditLog) database.AuditLog {
    return database.AuditLog{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Action : in.Action,
                User_id : in.User_id,
        Entity : in.Entity,
        Entityid : in.Entityid,
        Searchindex : in.Name,

    }
}
