/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package system

import (
	"log"

	"github.com/magicbutton/magic-meeting/applogic"
	"github.com/magicbutton/magic-meeting/database"
	"github.com/magicbutton/magic-meeting/services/models/systemmodel"
	. "github.com/magicbutton/magic-meeting/utils"
)

func SystemSearch(query string) (*Page[systemmodel.System], error) {
	log.Println("Calling Systemsearch")

	return applogic.Search[database.System, systemmodel.System]("searchindex", query, applogic.MapSystemOutgoing)

}
