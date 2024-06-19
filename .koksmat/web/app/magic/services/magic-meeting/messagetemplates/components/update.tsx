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
import {MessageTemplateForm} from "./form";

import {MessageTemplateItem} from "../applogic/model";
export default function UpdateMessageTemplate(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<MessageTemplateItem>(
      "magic-meeting.messagetemplates",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const messagetemplates = readResult.data;
    return (
      <div>{messagetemplates && 
      <MessageTemplateForm messagetemplates={messagetemplates} editmode="update"/>}
     
      </div>
    );
}
