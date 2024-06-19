/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package servicecategory

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicecategorymodel"

)

func ServiceCategoryRead(arg0 string) (*servicecategorymodel.ServiceCategory, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ServiceCategoryread")

    return applogic.Read[database.ServiceCategory, servicecategorymodel.ServiceCategory](id, applogic.MapServiceCategoryOutgoing)

}
