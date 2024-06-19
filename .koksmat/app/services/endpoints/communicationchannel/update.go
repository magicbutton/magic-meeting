/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package communicationchannel

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/communicationchannelmodel"

)

func CommunicationChannelUpdate(item communicationchannelmodel.CommunicationChannel) (*communicationchannelmodel.CommunicationChannel, error) {
    log.Println("Calling CommunicationChannelupdate")

    return applogic.Update[database.CommunicationChannel, communicationchannelmodel.CommunicationChannel](item.ID,item, applogic.MapCommunicationChannelIncoming, applogic.MapCommunicationChannelOutgoing)

}
