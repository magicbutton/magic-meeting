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
    import {ServiceOrderForm} from "./form";
    
    import {ServiceOrderItem} from "../applogic/model";
    export default function CreateServiceOrder() {
       
        const serviceorder = {} as ServiceOrderItem;
        return (
          <div>{serviceorder && 
          <ServiceOrderForm serviceorder={serviceorder} editmode="create"/>}
         
          </div>
        );
    }
