/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package auditlog

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/auditlogmodel"

)

func AuditLogRead(arg0 string) (*auditlogmodel.AuditLog, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling AuditLogread")

    return applogic.Read[database.AuditLog, auditlogmodel.AuditLog](id, applogic.MapAuditLogOutgoing)

}
