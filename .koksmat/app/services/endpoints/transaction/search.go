/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package transaction

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/transactionmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func TransactionSearch(query string) (*Page[transactionmodel.Transaction], error) {
    log.Println("Calling Transactionsearch")

    return applogic.Search[database.Transaction, transactionmodel.Transaction]("searchindex", query, applogic.MapTransactionOutgoing)

}
