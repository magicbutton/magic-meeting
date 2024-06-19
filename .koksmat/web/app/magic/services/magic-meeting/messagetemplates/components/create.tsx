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
    import {MessageTemplateForm} from "./form";
    
    import {MessageTemplateItem} from "../applogic/model";
    export default function CreateMessageTemplate() {
       
        const messagetemplates = {} as MessageTemplateItem;
        return (
          <div>{messagetemplates && 
          <MessageTemplateForm messagetemplates={messagetemplates} editmode="create"/>}
         
          </div>
        );
    }
