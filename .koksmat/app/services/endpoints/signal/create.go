/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package signal

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/signalmodel"

)

func SignalCreate(item signalmodel.Signal) (*signalmodel.Signal, error) {
    log.Println("Calling Signalcreate")

    return applogic.Create[database.Signal, signalmodel.Signal](item, applogic.MapSignalIncoming, applogic.MapSignalOutgoing)

}
