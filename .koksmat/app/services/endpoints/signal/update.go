/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package signal

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/signalmodel"

)

func SignalUpdate(item signalmodel.Signal) (*signalmodel.Signal, error) {
    log.Println("Calling Signalupdate")

    return applogic.Update[database.Signal, signalmodel.Signal](item.ID,item, applogic.MapSignalIncoming, applogic.MapSignalOutgoing)

}
