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
    import {MessageLogForm} from "./form";
    
    import {MessageLogItem} from "../applogic/model";
    export default function CreateMessageLog() {
       
        const messagelog = {} as MessageLogItem;
        return (
          <div>{messagelog && 
          <MessageLogForm messagelog={messagelog} editmode="create"/>}
         
          </div>
        );
    }
