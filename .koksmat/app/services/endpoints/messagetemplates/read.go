/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package messagetemplates

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/messagetemplatesmodel"

)

func MessageTemplateRead(arg0 string) (*messagetemplatesmodel.MessageTemplate, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling MessageTemplateread")

    return applogic.Read[database.MessageTemplate, messagetemplatesmodel.MessageTemplate](id, applogic.MapMessageTemplateOutgoing)

}
