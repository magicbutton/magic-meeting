/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package building

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/buildingmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func BuildingSearch(query string) (*Page[buildingmodel.Building], error) {
    log.Println("Calling Buildingsearch")

    return applogic.Search[database.Building, buildingmodel.Building]("searchindex", query, applogic.MapBuildingOutgoing)

}
