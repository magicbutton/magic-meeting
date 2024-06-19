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
	"github.com/magicbutton/magic-meeting/services/models/meetingmodel"
   
)


func MapMeetingOutgoing(db database.Meeting) meetingmodel.Meeting {
    return meetingmodel.Meeting{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Location : db.Location,
        Status : db.Status,
        Exchangereference : db.Exchangereference,
        Exchangestatus : db.Exchangestatus,
        Teamsreference : db.Teamsreference,
        Teamsstatus : db.Teamsstatus,

    }
}

func MapMeetingIncoming(in meetingmodel.Meeting) database.Meeting {
    return database.Meeting{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Location : in.Location,
        Status : in.Status,
        Exchangereference : in.Exchangereference,
        Exchangestatus : in.Exchangestatus,
        Teamsreference : in.Teamsreference,
        Teamsstatus : in.Teamsstatus,
        Searchindex : in.Name,

    }
}
