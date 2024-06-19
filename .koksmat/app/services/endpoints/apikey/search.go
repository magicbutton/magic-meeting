/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package apikey

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/apikeymodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func APIKeySearch(query string) (*Page[apikeymodel.APIKey], error) {
    log.Println("Calling APIKeysearch")

    return applogic.Search[database.APIKey, apikeymodel.APIKey]("searchindex", query, applogic.MapAPIKeyOutgoing)

}
