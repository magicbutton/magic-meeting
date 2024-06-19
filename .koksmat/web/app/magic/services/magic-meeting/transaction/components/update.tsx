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
import {TransactionForm} from "./form";

import {TransactionItem} from "../applogic/model";
export default function UpdateTransaction(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TransactionItem>(
      "magic-meeting.transaction",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const transaction = readResult.data;
    return (
      <div>{transaction && 
      <TransactionForm transaction={transaction} editmode="update"/>}
     
      </div>
    );
}
