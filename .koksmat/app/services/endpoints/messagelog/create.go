/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package messagelog

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/messagelogmodel"

)

func MessageLogCreate(item messagelogmodel.MessageLog) (*messagelogmodel.MessageLog, error) {
    log.Println("Calling MessageLogcreate")

    return applogic.Create[database.MessageLog, messagelogmodel.MessageLog](item, applogic.MapMessageLogIncoming, applogic.MapMessageLogOutgoing)

}
