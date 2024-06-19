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
import {AccessPassItem} from "../applogic/model";


/* yankiebar */

export default function ReadAccessPass(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AccessPassItem>(
      "magic-meeting.accesspass",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const accesspass = readResult.data;
    return (
      <div>
          
    {accesspass && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{accesspass.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{accesspass.description}</div>
    </div>    <div>
        <div className="font-bold" >visitor</div>
        <div>{accesspass.visitor_id}</div>
    </div>    <div>
        <div className="font-bold" >validfrom</div>
        <div>{accesspass.validfrom}</div>
    </div>    <div>
        <div className="font-bold" >validto</div>
        <div>{accesspass.validto}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{accesspass.status}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{accesspass.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{accesspass.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{accesspass.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{accesspass.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{accesspass.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
