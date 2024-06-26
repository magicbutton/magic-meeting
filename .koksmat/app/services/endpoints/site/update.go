/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package site

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/sitemodel"

)

func SiteUpdate(item sitemodel.Site) (*sitemodel.Site, error) {
    log.Println("Calling Siteupdate")

    return applogic.Update[database.Site, sitemodel.Site](item.ID,item, applogic.MapSiteIncoming, applogic.MapSiteOutgoing)

}
