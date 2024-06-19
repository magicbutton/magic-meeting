/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package visitor

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/visitormodel"

)

func VisitorUpdate(item visitormodel.Visitor) (*visitormodel.Visitor, error) {
    log.Println("Calling Visitorupdate")

    return applogic.Update[database.Visitor, visitormodel.Visitor](item.ID,item, applogic.MapVisitorIncoming, applogic.MapVisitorOutgoing)

}
