/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package service

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/servicemodel"

)

func ServiceRead(arg0 string) (*servicemodel.Service, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Serviceread")

    return applogic.Read[database.Service, servicemodel.Service](id, applogic.MapServiceOutgoing)

}
