/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package service

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicemodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func ServiceSearch(query string) (*Page[servicemodel.Service], error) {
    log.Println("Calling Servicesearch")

    return applogic.Search[database.Service, servicemodel.Service]("searchindex", query, applogic.MapServiceOutgoing)

}
