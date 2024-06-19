/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package servicecategory

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicecategorymodel"

)

func ServiceCategoryCreate(item servicecategorymodel.ServiceCategory) (*servicecategorymodel.ServiceCategory, error) {
    log.Println("Calling ServiceCategorycreate")

    return applogic.Create[database.ServiceCategory, servicecategorymodel.ServiceCategory](item, applogic.MapServiceCategoryIncoming, applogic.MapServiceCategoryOutgoing)

}
