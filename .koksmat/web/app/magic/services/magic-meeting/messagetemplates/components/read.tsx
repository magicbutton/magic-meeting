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
import {MessageTemplateItem} from "../applogic/model";


/* yankiebar */

export default function ReadMessageTemplate(props: { id: number }) {
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
      <div>
          
    {messagetemplates && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{messagetemplates.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{messagetemplates.description}</div>
    </div>    <div>
        <div className="font-bold" >type</div>
        <div>{messagetemplates.type}</div>
    </div>    <div>
        <div className="font-bold" >content</div>
        <div>{messagetemplates.content}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{messagetemplates.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{messagetemplates.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{messagetemplates.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{messagetemplates.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{messagetemplates.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
