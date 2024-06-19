/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package meeting

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingmodel"

)

func MeetingCreate(item meetingmodel.Meeting) (*meetingmodel.Meeting, error) {
    log.Println("Calling Meetingcreate")

    return applogic.Create[database.Meeting, meetingmodel.Meeting](item, applogic.MapMeetingIncoming, applogic.MapMeetingOutgoing)

}
