/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package productionorder

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/productionordermodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func ProductionOrderSearch(query string) (*Page[productionordermodel.ProductionOrder], error) {
    log.Println("Calling ProductionOrdersearch")

    return applogic.Search[database.ProductionOrder, productionordermodel.ProductionOrder]("searchindex", query, applogic.MapProductionOrderOutgoing)

}
