/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package productionorder

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/productionordermodel"

)

func ProductionOrderUpdate(item productionordermodel.ProductionOrder) (*productionordermodel.ProductionOrder, error) {
    log.Println("Calling ProductionOrderupdate")

    return applogic.Update[database.ProductionOrder, productionordermodel.ProductionOrder](item.ID,item, applogic.MapProductionOrderIncoming, applogic.MapProductionOrderOutgoing)

}
