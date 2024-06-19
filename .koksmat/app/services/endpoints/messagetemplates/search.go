/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package messagetemplates

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/messagetemplatesmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func MessageTemplateSearch(query string) (*Page[messagetemplatesmodel.MessageTemplate], error) {
    log.Println("Calling MessageTemplatesearch")

    return applogic.Search[database.MessageTemplate, messagetemplatesmodel.MessageTemplate]("searchindex", query, applogic.MapMessageTemplateOutgoing)

}
