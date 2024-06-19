/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package tenant

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/tenantmodel"

)

func TenantRead(arg0 string) (*tenantmodel.Tenant, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Tenantread")

    return applogic.Read[database.Tenant, tenantmodel.Tenant](id, applogic.MapTenantOutgoing)

}
