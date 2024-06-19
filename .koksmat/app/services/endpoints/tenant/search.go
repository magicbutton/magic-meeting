/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package tenant

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/tenantmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func TenantSearch(query string) (*Page[tenantmodel.Tenant], error) {
    log.Println("Calling Tenantsearch")

    return applogic.Search[database.Tenant, tenantmodel.Tenant]("searchindex", query, applogic.MapTenantOutgoing)

}
