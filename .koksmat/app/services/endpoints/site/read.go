/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package site

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/sitemodel"

)

func SiteRead(arg0 string) (*sitemodel.Site, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Siteread")

    return applogic.Read[database.Site, sitemodel.Site](id, applogic.MapSiteOutgoing)

}
