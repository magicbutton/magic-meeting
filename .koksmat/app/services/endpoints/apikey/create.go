/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package apikey

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/apikeymodel"

)

func APIKeyCreate(item apikeymodel.APIKey) (*apikeymodel.APIKey, error) {
    log.Println("Calling APIKeycreate")

    return applogic.Create[database.APIKey, apikeymodel.APIKey](item, applogic.MapAPIKeyIncoming, applogic.MapAPIKeyOutgoing)

}
