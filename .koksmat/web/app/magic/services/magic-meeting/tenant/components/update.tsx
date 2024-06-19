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
import {TenantForm} from "./form";

import {TenantItem} from "../applogic/model";
export default function UpdateTenant(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TenantItem>(
      "magic-meeting.tenant",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const tenant = readResult.data;
    return (
      <div>{tenant && 
      <TenantForm tenant={tenant} editmode="update"/>}
     
      </div>
    );
}
