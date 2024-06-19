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
	"github.com/magicbutton/magic-meeting/services/models/meetingroommodel"
   
)


func MapMeetingRoomOutgoing(db database.MeetingRoom) meetingroommodel.MeetingRoom {
    return meetingroommodel.MeetingRoom{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Email : db.Email,
        Features : db.Features,
                Floor_id : db.Floor_id,

    }
}

func MapMeetingRoomIncoming(in meetingroommodel.MeetingRoom) database.MeetingRoom {
    return database.MeetingRoom{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Email : in.Email,
        Features : in.Features,
                Floor_id : in.Floor_id,
        Searchindex : in.Name,

    }
}
