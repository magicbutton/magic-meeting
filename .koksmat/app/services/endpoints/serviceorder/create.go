/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package serviceorder

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/serviceordermodel"

)

func ServiceOrderCreate(item serviceordermodel.ServiceOrder) (*serviceordermodel.ServiceOrder, error) {
    log.Println("Calling ServiceOrdercreate")

    return applogic.Create[database.ServiceOrder, serviceordermodel.ServiceOrder](item, applogic.MapServiceOrderIncoming, applogic.MapServiceOrderOutgoing)

}
