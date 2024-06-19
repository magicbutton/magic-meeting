/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package meeting

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/meetingmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func MeetingSearch(query string) (*Page[meetingmodel.Meeting], error) {
    log.Println("Calling Meetingsearch")

    return applogic.Search[database.Meeting, meetingmodel.Meeting]("searchindex", query, applogic.MapMeetingOutgoing)

}
