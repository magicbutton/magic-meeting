/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package vendor

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/vendormodel"

)

func VendorCreate(item vendormodel.Vendor) (*vendormodel.Vendor, error) {
    log.Println("Calling Vendorcreate")

    return applogic.Create[database.Vendor, vendormodel.Vendor](item, applogic.MapVendorIncoming, applogic.MapVendorOutgoing)

}
