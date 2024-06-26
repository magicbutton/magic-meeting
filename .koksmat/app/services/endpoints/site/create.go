/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package site

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/sitemodel"

)

func SiteCreate(item sitemodel.Site) (*sitemodel.Site, error) {
    log.Println("Calling Sitecreate")

    return applogic.Create[database.Site, sitemodel.Site](item, applogic.MapSiteIncoming, applogic.MapSiteOutgoing)

}
