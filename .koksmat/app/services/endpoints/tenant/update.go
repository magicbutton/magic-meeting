/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package tenant

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/tenantmodel"

)

func TenantUpdate(item tenantmodel.Tenant) (*tenantmodel.Tenant, error) {
    log.Println("Calling Tenantupdate")

    return applogic.Update[database.Tenant, tenantmodel.Tenant](item.ID,item, applogic.MapTenantIncoming, applogic.MapTenantOutgoing)

}
