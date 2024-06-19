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
import {SignalForm} from "./form";

import {SignalItem} from "../applogic/model";
export default function UpdateSignal(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<SignalItem>(
      "magic-meeting.signal",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const signal = readResult.data;
    return (
      <div>{signal && 
      <SignalForm signal={signal} editmode="update"/>}
     
      </div>
    );
}
