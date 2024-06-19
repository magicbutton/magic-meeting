/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package vendor

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/vendormodel"

)

func VendorRead(arg0 string) (*vendormodel.Vendor, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Vendorread")

    return applogic.Read[database.Vendor, vendormodel.Vendor](id, applogic.MapVendorOutgoing)

}
