/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package meeting

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingmodel"

)

func MeetingUpdate(item meetingmodel.Meeting) (*meetingmodel.Meeting, error) {
    log.Println("Calling Meetingupdate")

    return applogic.Update[database.Meeting, meetingmodel.Meeting](item.ID,item, applogic.MapMeetingIncoming, applogic.MapMeetingOutgoing)

}
