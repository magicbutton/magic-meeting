/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package servicecategory

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicecategorymodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func ServiceCategorySearch(query string) (*Page[servicecategorymodel.ServiceCategory], error) {
    log.Println("Calling ServiceCategorysearch")

    return applogic.Search[database.ServiceCategory, servicecategorymodel.ServiceCategory]("searchindex", query, applogic.MapServiceCategoryOutgoing)

}
