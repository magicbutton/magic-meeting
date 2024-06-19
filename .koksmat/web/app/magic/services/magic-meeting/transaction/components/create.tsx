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
    import {TransactionForm} from "./form";
    
    import {TransactionItem} from "../applogic/model";
    export default function CreateTransaction() {
       
        const transaction = {} as TransactionItem;
        return (
          <div>{transaction && 
          <TransactionForm transaction={transaction} editmode="create"/>}
         
          </div>
        );
    }
