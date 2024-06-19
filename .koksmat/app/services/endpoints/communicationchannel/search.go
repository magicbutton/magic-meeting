/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package communicationchannel

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/communicationchannelmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func CommunicationChannelSearch(query string) (*Page[communicationchannelmodel.CommunicationChannel], error) {
    log.Println("Calling CommunicationChannelsearch")

    return applogic.Search[database.CommunicationChannel, communicationchannelmodel.CommunicationChannel]("searchindex", query, applogic.MapCommunicationChannelOutgoing)

}
