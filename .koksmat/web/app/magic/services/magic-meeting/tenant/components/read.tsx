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
import {TenantItem} from "../applogic/model";


/* yankiebar */

export default function ReadTenant(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TenantItem>(
      "magic-meeting.tenant",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const tenant = readResult.data;
    return (
      <div>
          
    {tenant && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{tenant.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{tenant.description}</div>
    </div>    <div>
        <div className="font-bold" >email</div>
        <div>{tenant.email}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{tenant.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{tenant.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{tenant.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{tenant.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{tenant.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
