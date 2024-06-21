/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package relation

import (
	"log"

	"github.com/magicbutton/magic-meeting/applogic"
	"github.com/magicbutton/magic-meeting/database"
	"github.com/magicbutton/magic-meeting/services/models/relationmodel"
	. "github.com/magicbutton/magic-meeting/utils"
)

func RelationSearch(query string) (*Page[relationmodel.Relation], error) {
	log.Println("Calling Relationsearch")

	return applogic.Search[database.Relation, relationmodel.Relation]("searchindex", query, applogic.MapRelationOutgoing)

}
