/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package accesspass

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/accesspassmodel"

)

func AccessPassUpdate(item accesspassmodel.AccessPass) (*accesspassmodel.AccessPass, error) {
    log.Println("Calling AccessPassupdate")

    return applogic.Update[database.AccessPass, accesspassmodel.AccessPass](item.ID,item, applogic.MapAccessPassIncoming, applogic.MapAccessPassOutgoing)

}
