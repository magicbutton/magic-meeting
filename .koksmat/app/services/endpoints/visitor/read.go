/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package visitor

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/visitormodel"

)

func VisitorRead(arg0 string) (*visitormodel.Visitor, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Visitorread")

    return applogic.Read[database.Visitor, visitormodel.Visitor](id, applogic.MapVisitorOutgoing)

}
