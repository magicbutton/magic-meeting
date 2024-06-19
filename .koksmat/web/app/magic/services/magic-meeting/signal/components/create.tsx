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
    import {SignalForm} from "./form";
    
    import {SignalItem} from "../applogic/model";
    export default function CreateSignal() {
       
        const signal = {} as SignalItem;
        return (
          <div>{signal && 
          <SignalForm signal={signal} editmode="create"/>}
         
          </div>
        );
    }
