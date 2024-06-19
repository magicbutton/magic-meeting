/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package serviceorder

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/serviceordermodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func ServiceOrderSearch(query string) (*Page[serviceordermodel.ServiceOrder], error) {
    log.Println("Calling ServiceOrdersearch")

    return applogic.Search[database.ServiceOrder, serviceordermodel.ServiceOrder]("searchindex", query, applogic.MapServiceOrderOutgoing)

}
