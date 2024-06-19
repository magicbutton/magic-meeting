/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package meetingroom

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingroommodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func MeetingRoomSearch(query string) (*Page[meetingroommodel.MeetingRoom], error) {
    log.Println("Calling MeetingRoomsearch")

    return applogic.Search[database.MeetingRoom, meetingroommodel.MeetingRoom]("searchindex", query, applogic.MapMeetingRoomOutgoing)

}
