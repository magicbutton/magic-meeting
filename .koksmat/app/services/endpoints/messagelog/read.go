/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package messagelog

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/messagelogmodel"

)

func MessageLogRead(arg0 string) (*messagelogmodel.MessageLog, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling MessageLogread")

    return applogic.Read[database.MessageLog, messagelogmodel.MessageLog](id, applogic.MapMessageLogOutgoing)

}
