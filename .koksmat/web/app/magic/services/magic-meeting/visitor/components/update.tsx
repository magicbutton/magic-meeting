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
import {VisitorForm} from "./form";

import {VisitorItem} from "../applogic/model";
export default function UpdateVisitor(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<VisitorItem>(
      "magic-meeting.visitor",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const visitor = readResult.data;
    return (
      <div>{visitor && 
      <VisitorForm visitor={visitor} editmode="update"/>}
     
      </div>
    );
}
