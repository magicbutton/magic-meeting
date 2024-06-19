/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package visitor

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/visitormodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func VisitorSearch(query string) (*Page[visitormodel.Visitor], error) {
    log.Println("Calling Visitorsearch")

    return applogic.Search[database.Visitor, visitormodel.Visitor]("searchindex", query, applogic.MapVisitorOutgoing)

}
