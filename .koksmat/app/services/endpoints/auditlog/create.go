/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package auditlog

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/auditlogmodel"

)

func AuditLogCreate(item auditlogmodel.AuditLog) (*auditlogmodel.AuditLog, error) {
    log.Println("Calling AuditLogcreate")

    return applogic.Create[database.AuditLog, auditlogmodel.AuditLog](item, applogic.MapAuditLogIncoming, applogic.MapAuditLogOutgoing)

}
