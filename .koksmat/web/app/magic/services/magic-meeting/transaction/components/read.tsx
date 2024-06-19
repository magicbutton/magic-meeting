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
import {TransactionItem} from "../applogic/model";


/* yankiebar */

export default function ReadTransaction(props: { id: number }) {
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
      <div>
          
    {transaction && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{transaction.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{transaction.description}</div>
    </div>    <div>
        <div className="font-bold" >amount</div>
        <div>{transaction.amount}</div>
    </div>    <div>
        <div className="font-bold" >currency</div>
        <div>{transaction.currency}</div>
    </div>    <div>
        <div className="font-bold" >datetime</div>
        <div>{transaction.datetime}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{transaction.status}</div>
    </div>    <div>
        <div className="font-bold" >account</div>
        <div>{transaction.account_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{transaction.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{transaction.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{transaction.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{transaction.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{transaction.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
