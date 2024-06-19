/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package messagetemplates

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/messagetemplatesmodel"

)

func MessageTemplateUpdate(item messagetemplatesmodel.MessageTemplate) (*messagetemplatesmodel.MessageTemplate, error) {
    log.Println("Calling MessageTemplateupdate")

    return applogic.Update[database.MessageTemplate, messagetemplatesmodel.MessageTemplate](item.ID,item, applogic.MapMessageTemplateIncoming, applogic.MapMessageTemplateOutgoing)

}
