    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/app/koksmat/useservice";
    import { useState } from "react";
    import {VendorForm} from "./form";
    
    import {VendorItem} from "../applogic/model";
    export default function CreateVendor() {
       
        const vendor = {} as VendorItem;
        return (
          <div>{vendor && 
          <VendorForm vendor={vendor} editmode="create"/>}
         
          </div>
        );
    }
