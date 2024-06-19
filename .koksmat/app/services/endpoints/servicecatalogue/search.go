/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package servicecatalogue

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicecataloguemodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func ServiceCatalogueSearch(query string) (*Page[servicecataloguemodel.ServiceCatalogue], error) {
    log.Println("Calling ServiceCataloguesearch")

    return applogic.Search[database.ServiceCatalogue, servicecataloguemodel.ServiceCatalogue]("searchindex", query, applogic.MapServiceCatalogueOutgoing)

}
