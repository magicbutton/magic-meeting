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
import {MessageLogItem} from "../applogic/model";


/* yankiebar */

export default function ReadMessageLog(props: { id: number }) {
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
      <div>
          
    {messagelog && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{messagelog.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{messagelog.description}</div>
    </div>    <div>
        <div className="font-bold" >sender</div>
        <div>{messagelog.sender}</div>
    </div>    <div>
        <div className="font-bold" >receiver</div>
        <div>{messagelog.receiver}</div>
    </div>    <div>
        <div className="font-bold" >message</div>
        <div>{messagelog.message}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{messagelog.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{messagelog.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{messagelog.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{messagelog.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{messagelog.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
