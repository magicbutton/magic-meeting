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
import {SignalItem} from "../applogic/model";


/* yankiebar */

export default function ReadSignal(props: { id: number }) {
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
      <div>
          
    {signal && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{signal.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{signal.description}</div>
    </div>    <div>
        <div className="font-bold" >sender</div>
        <div>{signal.sender}</div>
    </div>    <div>
        <div className="font-bold" >receiver</div>
        <div>{signal.receiver}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{signal.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{signal.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{signal.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{signal.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{signal.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
