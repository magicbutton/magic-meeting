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
    import {AuditLogForm} from "./form";
    
    import {AuditLogItem} from "../applogic/model";
    export default function CreateAuditLog() {
       
        const auditlog = {} as AuditLogItem;
        return (
          <div>{auditlog && 
          <AuditLogForm auditlog={auditlog} editmode="create"/>}
         
          </div>
        );
    }
