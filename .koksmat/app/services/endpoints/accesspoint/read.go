/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package accesspoint

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/accesspointmodel"

)

func AccessPointRead(arg0 string) (*accesspointmodel.AccessPoint, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling AccessPointread")

    return applogic.Read[database.AccessPoint, accesspointmodel.AccessPoint](id, applogic.MapAccessPointOutgoing)

}
