/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package vendor

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/vendormodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func VendorSearch(query string) (*Page[vendormodel.Vendor], error) {
    log.Println("Calling Vendorsearch")

    return applogic.Search[database.Vendor, vendormodel.Vendor]("searchindex", query, applogic.MapVendorOutgoing)

}
