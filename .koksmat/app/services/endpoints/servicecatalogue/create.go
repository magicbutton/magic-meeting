/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package servicecatalogue

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicecataloguemodel"

)

func ServiceCatalogueCreate(item servicecataloguemodel.ServiceCatalogue) (*servicecataloguemodel.ServiceCatalogue, error) {
    log.Println("Calling ServiceCataloguecreate")

    return applogic.Create[database.ServiceCatalogue, servicecataloguemodel.ServiceCatalogue](item, applogic.MapServiceCatalogueIncoming, applogic.MapServiceCatalogueOutgoing)

}
