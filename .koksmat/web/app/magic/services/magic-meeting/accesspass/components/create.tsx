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
    import {AccessPassForm} from "./form";
    
    import {AccessPassItem} from "../applogic/model";
    export default function CreateAccessPass() {
       
        const accesspass = {} as AccessPassItem;
        return (
          <div>{accesspass && 
          <AccessPassForm accesspass={accesspass} editmode="create"/>}
         
          </div>
        );
    }
