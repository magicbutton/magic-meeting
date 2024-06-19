/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package user

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/usermodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func UserSearch(query string) (*Page[usermodel.User], error) {
    log.Println("Calling Usersearch")

    return applogic.Search[database.User, usermodel.User]("searchindex", query, applogic.MapUserOutgoing)

}
