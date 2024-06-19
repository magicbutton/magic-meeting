/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package messagelog

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/messagelogmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func MessageLogSearch(query string) (*Page[messagelogmodel.MessageLog], error) {
    log.Println("Calling MessageLogsearch")

    return applogic.Search[database.MessageLog, messagelogmodel.MessageLog]("searchindex", query, applogic.MapMessageLogOutgoing)

}
