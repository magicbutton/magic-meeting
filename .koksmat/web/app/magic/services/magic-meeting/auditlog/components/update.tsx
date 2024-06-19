/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {AuditLogForm} from "./form";

import {AuditLogItem} from "../applogic/model";
export default function UpdateAuditLog(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AuditLogItem>(
      "magic-meeting.auditlog",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const auditlog = readResult.data;
    return (
      <div>{auditlog && 
      <AuditLogForm auditlog={auditlog} editmode="update"/>}
     
      </div>
    );
}
