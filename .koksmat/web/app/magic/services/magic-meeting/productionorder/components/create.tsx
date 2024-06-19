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
    import {ProductionOrderForm} from "./form";
    
    import {ProductionOrderItem} from "../applogic/model";
    export default function CreateProductionOrder() {
       
        const productionorder = {} as ProductionOrderItem;
        return (
          <div>{productionorder && 
          <ProductionOrderForm productionorder={productionorder} editmode="create"/>}
         
          </div>
        );
    }
