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
import {VisitorItem} from "../applogic/model";


/* yankiebar */

export default function ReadVisitor(props: { id: number }) {
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
      <div>
          
    {visitor && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{visitor.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{visitor.description}</div>
    </div>    <div>
        <div className="font-bold" >email</div>
        <div>{visitor.email}</div>
    </div>    <div>
        <div className="font-bold" >phone</div>
        <div>{visitor.phone}</div>
    </div>    <div>
        <div className="font-bold" >company</div>
        <div>{visitor.company}</div>
    </div>    <div>
        <div className="font-bold" >purpose</div>
        <div>{visitor.purpose}</div>
    </div>    <div>
        <div className="font-bold" >host</div>
        <div>{visitor.host_id}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{visitor.status}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{visitor.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{visitor.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{visitor.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{visitor.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{visitor.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
