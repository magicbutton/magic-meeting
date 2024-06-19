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
import {APIKeyItem} from "../applogic/model";


/* yankiebar */

export default function ReadAPIKey(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<APIKeyItem>(
      "magic-meeting.apikey",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const apikey = readResult.data;
    return (
      <div>
          
    {apikey && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{apikey.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{apikey.description}</div>
    </div>    <div>
        <div className="font-bold" >key</div>
        <div>{apikey.key}</div>
    </div>    <div>
        <div className="font-bold" >user</div>
        <div>{apikey.user_id}</div>
    </div>    <div>
        <div className="font-bold" >validto</div>
        <div>{apikey.validto}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{apikey.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{apikey.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{apikey.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{apikey.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{apikey.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
