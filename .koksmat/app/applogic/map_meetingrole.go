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
	"github.com/magicbutton/magic-meeting/services/models/meetingrolemodel"
   
)


func MapMeetingRoleOutgoing(db database.MeetingRole) meetingrolemodel.MeetingRole {
    return meetingrolemodel.MeetingRole{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                User_id : db.User_id,
                Meeting_id : db.Meeting_id,
        Role : db.Role,

    }
}

func MapMeetingRoleIncoming(in meetingrolemodel.MeetingRole) database.MeetingRole {
    return database.MeetingRole{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                User_id : in.User_id,
                Meeting_id : in.Meeting_id,
        Role : in.Role,
        Searchindex : in.Name,

    }
}
