/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateMapModel v1.1
package applogic
import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-meeting/database"
	"github.com/magicbutton/magic-meeting/services/models/servicecataloguemodel"
   
)


func MapServiceCatalogueOutgoing(db database.ServiceCatalogue) servicecataloguemodel.ServiceCatalogue {
    return servicecataloguemodel.ServiceCatalogue{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,

    }
}

func MapServiceCatalogueIncoming(in servicecataloguemodel.ServiceCatalogue) database.ServiceCatalogue {
    return database.ServiceCatalogue{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Searchindex : in.Name,

    }
}
