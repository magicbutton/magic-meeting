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
import {MessageLogForm} from "./form";

import {MessageLogItem} from "../applogic/model";
export default function UpdateMessageLog(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<MessageLogItem>(
      "magic-meeting.messagelog",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const messagelog = readResult.data;
    return (
      <div>{messagelog && 
      <MessageLogForm messagelog={messagelog} editmode="update"/>}
     
      </div>
    );
}
