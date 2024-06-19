/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package building

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/buildingmodel"

)

func BuildingRead(arg0 string) (*buildingmodel.Building, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Buildingread")

    return applogic.Read[database.Building, buildingmodel.Building](id, applogic.MapBuildingOutgoing)

}
