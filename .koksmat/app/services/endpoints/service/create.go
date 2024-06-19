/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package service

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicemodel"

)

func ServiceCreate(item servicemodel.Service) (*servicemodel.Service, error) {
    log.Println("Calling Servicecreate")

    return applogic.Create[database.Service, servicemodel.Service](item, applogic.MapServiceIncoming, applogic.MapServiceOutgoing)

}
