/*
   File have been automatically created. To prevent the file from getting overwritten
   set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
   ---
   keep: false
   ---
*/
//generator:  noma3.delete.v2
package relation

import (
	"log"
	"strconv"

	"github.com/magicbutton/magic-meeting/applogic"
	"github.com/magicbutton/magic-meeting/database"
	"github.com/magicbutton/magic-meeting/services/models/relationmodel"
)

func RelationDelete(arg0 string) error {
	id, _ := strconv.Atoi(arg0)
	log.Println("Calling Relationdelete")

	return applogic.Delete[database.Relation, relationmodel.Relation](id)

}
