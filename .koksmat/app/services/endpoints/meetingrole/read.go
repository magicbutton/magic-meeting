/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package meetingrole

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingrolemodel"

)

func MeetingRoleRead(arg0 string) (*meetingrolemodel.MeetingRole, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling MeetingRoleread")

    return applogic.Read[database.MeetingRole, meetingrolemodel.MeetingRole](id, applogic.MapMeetingRoleOutgoing)

}
