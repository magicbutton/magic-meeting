/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package tenant

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/tenantmodel"

)

func TenantCreate(item tenantmodel.Tenant) (*tenantmodel.Tenant, error) {
    log.Println("Calling Tenantcreate")

    return applogic.Create[database.Tenant, tenantmodel.Tenant](item, applogic.MapTenantIncoming, applogic.MapTenantOutgoing)

}
