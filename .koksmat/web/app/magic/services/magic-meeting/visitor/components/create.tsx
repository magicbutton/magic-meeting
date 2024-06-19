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
    import {VisitorForm} from "./form";
    
    import {VisitorItem} from "../applogic/model";
    export default function CreateVisitor() {
       
        const visitor = {} as VisitorItem;
        return (
          <div>{visitor && 
          <VisitorForm visitor={visitor} editmode="create"/>}
         
          </div>
        );
    }
