/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package meetingroom

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingroommodel"

)

func MeetingRoomRead(arg0 string) (*meetingroommodel.MeetingRoom, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling MeetingRoomread")

    return applogic.Read[database.MeetingRoom, meetingroommodel.MeetingRoom](id, applogic.MapMeetingRoomOutgoing)

}
