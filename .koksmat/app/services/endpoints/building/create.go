/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package building

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/buildingmodel"

)

func BuildingCreate(item buildingmodel.Building) (*buildingmodel.Building, error) {
    log.Println("Calling Buildingcreate")

    return applogic.Create[database.Building, buildingmodel.Building](item, applogic.MapBuildingIncoming, applogic.MapBuildingOutgoing)

}
