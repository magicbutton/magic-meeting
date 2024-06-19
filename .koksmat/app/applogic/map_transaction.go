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
	"github.com/magicbutton/magic-meeting/services/models/transactionmodel"
   
)


func MapTransactionOutgoing(db database.Transaction) transactionmodel.Transaction {
    return transactionmodel.Transaction{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Currency : db.Currency,
        Status : db.Status,
                Account_id : db.Account_id,

    }
}

func MapTransactionIncoming(in transactionmodel.Transaction) database.Transaction {
    return database.Transaction{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Currency : in.Currency,
        Status : in.Status,
                Account_id : in.Account_id,
        Searchindex : in.Name,

    }
}
