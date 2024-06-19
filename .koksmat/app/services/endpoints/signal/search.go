/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package signal

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/signalmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func SignalSearch(query string) (*Page[signalmodel.Signal], error) {
    log.Println("Calling Signalsearch")

    return applogic.Search[database.Signal, signalmodel.Signal]("searchindex", query, applogic.MapSignalOutgoing)

}
