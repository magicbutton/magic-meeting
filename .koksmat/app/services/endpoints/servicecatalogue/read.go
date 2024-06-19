/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package servicecatalogue

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicecataloguemodel"

)

func ServiceCatalogueRead(arg0 string) (*servicecataloguemodel.ServiceCatalogue, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ServiceCatalogueread")

    return applogic.Read[database.ServiceCatalogue, servicecataloguemodel.ServiceCatalogue](id, applogic.MapServiceCatalogueOutgoing)

}
