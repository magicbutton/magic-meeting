/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package meetingrole

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingrolemodel"

)

func MeetingRoleCreate(item meetingrolemodel.MeetingRole) (*meetingrolemodel.MeetingRole, error) {
    log.Println("Calling MeetingRolecreate")

    return applogic.Create[database.MeetingRole, meetingrolemodel.MeetingRole](item, applogic.MapMeetingRoleIncoming, applogic.MapMeetingRoleOutgoing)

}
