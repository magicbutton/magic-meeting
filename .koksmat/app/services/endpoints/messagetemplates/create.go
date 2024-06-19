/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package messagetemplates

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/messagetemplatesmodel"

)

func MessageTemplateCreate(item messagetemplatesmodel.MessageTemplate) (*messagetemplatesmodel.MessageTemplate, error) {
    log.Println("Calling MessageTemplatecreate")

    return applogic.Create[database.MessageTemplate, messagetemplatesmodel.MessageTemplate](item, applogic.MapMessageTemplateIncoming, applogic.MapMessageTemplateOutgoing)

}
