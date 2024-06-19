/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package meetingrole

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingrolemodel"

)

func MeetingRoleUpdate(item meetingrolemodel.MeetingRole) (*meetingrolemodel.MeetingRole, error) {
    log.Println("Calling MeetingRoleupdate")

    return applogic.Update[database.MeetingRole, meetingrolemodel.MeetingRole](item.ID,item, applogic.MapMeetingRoleIncoming, applogic.MapMeetingRoleOutgoing)

}
