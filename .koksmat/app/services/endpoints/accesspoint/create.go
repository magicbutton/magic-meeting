/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package accesspoint

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/accesspointmodel"

)

func AccessPointCreate(item accesspointmodel.AccessPoint) (*accesspointmodel.AccessPoint, error) {
    log.Println("Calling AccessPointcreate")

    return applogic.Create[database.AccessPoint, accesspointmodel.AccessPoint](item, applogic.MapAccessPointIncoming, applogic.MapAccessPointOutgoing)

}
