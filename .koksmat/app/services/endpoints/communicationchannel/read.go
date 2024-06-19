/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package communicationchannel

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/communicationchannelmodel"

)

func CommunicationChannelRead(arg0 string) (*communicationchannelmodel.CommunicationChannel, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling CommunicationChannelread")

    return applogic.Read[database.CommunicationChannel, communicationchannelmodel.CommunicationChannel](id, applogic.MapCommunicationChannelOutgoing)

}
