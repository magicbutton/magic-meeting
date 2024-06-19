/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package servicecatalogue

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicecataloguemodel"

)

func ServiceCatalogueUpdate(item servicecataloguemodel.ServiceCatalogue) (*servicecataloguemodel.ServiceCatalogue, error) {
    log.Println("Calling ServiceCatalogueupdate")

    return applogic.Update[database.ServiceCatalogue, servicecataloguemodel.ServiceCatalogue](item.ID,item, applogic.MapServiceCatalogueIncoming, applogic.MapServiceCatalogueOutgoing)

}
