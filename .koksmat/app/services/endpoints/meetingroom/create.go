/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package meetingroom

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingroommodel"

)

func MeetingRoomCreate(item meetingroommodel.MeetingRoom) (*meetingroommodel.MeetingRoom, error) {
    log.Println("Calling MeetingRoomcreate")

    return applogic.Create[database.MeetingRoom, meetingroommodel.MeetingRoom](item, applogic.MapMeetingRoomIncoming, applogic.MapMeetingRoomOutgoing)

}
