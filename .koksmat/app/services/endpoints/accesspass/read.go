/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package accesspass

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/accesspassmodel"

)

func AccessPassRead(arg0 string) (*accesspassmodel.AccessPass, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling AccessPassread")

    return applogic.Read[database.AccessPass, accesspassmodel.AccessPass](id, applogic.MapAccessPassOutgoing)

}
