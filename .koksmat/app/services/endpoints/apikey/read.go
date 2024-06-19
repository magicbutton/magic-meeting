/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package apikey

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/apikeymodel"

)

func APIKeyRead(arg0 string) (*apikeymodel.APIKey, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling APIKeyread")

    return applogic.Read[database.APIKey, apikeymodel.APIKey](id, applogic.MapAPIKeyOutgoing)

}
