/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package meeting

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingmodel"

)

func MeetingRead(arg0 string) (*meetingmodel.Meeting, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Meetingread")

    return applogic.Read[database.Meeting, meetingmodel.Meeting](id, applogic.MapMeetingOutgoing)

}
