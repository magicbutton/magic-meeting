/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package accesspoint

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/accesspointmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func AccessPointSearch(query string) (*Page[accesspointmodel.AccessPoint], error) {
    log.Println("Calling AccessPointsearch")

    return applogic.Search[database.AccessPoint, accesspointmodel.AccessPoint]("searchindex", query, applogic.MapAccessPointOutgoing)

}
