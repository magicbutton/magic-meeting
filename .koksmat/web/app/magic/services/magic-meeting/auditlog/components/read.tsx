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
import {AuditLogItem} from "../applogic/model";


/* yankiebar */

export default function ReadAuditLog(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AuditLogItem>(
      "magic-meeting.auditlog",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const auditlog = readResult.data;
    return (
      <div>
          
    {auditlog && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{auditlog.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{auditlog.description}</div>
    </div>    <div>
        <div className="font-bold" >action</div>
        <div>{auditlog.action}</div>
    </div>    <div>
        <div className="font-bold" >user</div>
        <div>{auditlog.user_id}</div>
    </div>    <div>
        <div className="font-bold" >entity</div>
        <div>{auditlog.entity}</div>
    </div>    <div>
        <div className="font-bold" >entityid</div>
        <div>{auditlog.entityid}</div>
    </div>    <div>
        <div className="font-bold" >timestamp</div>
        <div>{auditlog.timestamp}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{auditlog.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{auditlog.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{auditlog.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{auditlog.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{auditlog.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
