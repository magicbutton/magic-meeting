/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package serviceorder

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/serviceordermodel"

)

func ServiceOrderRead(arg0 string) (*serviceordermodel.ServiceOrder, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ServiceOrderread")

    return applogic.Read[database.ServiceOrder, serviceordermodel.ServiceOrder](id, applogic.MapServiceOrderOutgoing)

}
