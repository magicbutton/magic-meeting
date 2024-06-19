/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package serviceorder

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/serviceordermodel"

)

func ServiceOrderUpdate(item serviceordermodel.ServiceOrder) (*serviceordermodel.ServiceOrder, error) {
    log.Println("Calling ServiceOrderupdate")

    return applogic.Update[database.ServiceOrder, serviceordermodel.ServiceOrder](item.ID,item, applogic.MapServiceOrderIncoming, applogic.MapServiceOrderOutgoing)

}
