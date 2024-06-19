/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package transaction

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/transactionmodel"

)

func TransactionCreate(item transactionmodel.Transaction) (*transactionmodel.Transaction, error) {
    log.Println("Calling Transactioncreate")

    return applogic.Create[database.Transaction, transactionmodel.Transaction](item, applogic.MapTransactionIncoming, applogic.MapTransactionOutgoing)

}
