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
import {AccessPointItem} from "../applogic/model";


/* yankiebar */

export default function ReadAccessPoint(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AccessPointItem>(
      "magic-meeting.accesspoint",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const accesspoint = readResult.data;
    return (
      <div>
          
    {accesspoint && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{accesspoint.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{accesspoint.description}</div>
    </div>    <div>
        <div className="font-bold" >location</div>
        <div>{accesspoint.location_id}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{accesspoint.status}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{accesspoint.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{accesspoint.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{accesspoint.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{accesspoint.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{accesspoint.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
