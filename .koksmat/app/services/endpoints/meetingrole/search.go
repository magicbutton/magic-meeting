/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package meetingrole

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingrolemodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func MeetingRoleSearch(query string) (*Page[meetingrolemodel.MeetingRole], error) {
    log.Println("Calling MeetingRolesearch")

    return applogic.Search[database.MeetingRole, meetingrolemodel.MeetingRole]("searchindex", query, applogic.MapMeetingRoleOutgoing)

}
