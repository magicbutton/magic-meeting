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
import {CommunicationChannelItem} from "../applogic/model";


/* yankiebar */

export default function ReadCommunicationChannel(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<CommunicationChannelItem>(
      "magic-meeting.communicationchannel",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const communicationchannel = readResult.data;
    return (
      <div>
          
    {communicationchannel && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{communicationchannel.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{communicationchannel.description}</div>
    </div>    <div>
        <div className="font-bold" >type</div>
        <div>{communicationchannel.type}</div>
    </div>    <div>
        <div className="font-bold" >address</div>
        <div>{communicationchannel.address}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{communicationchannel.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{communicationchannel.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{communicationchannel.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{communicationchannel.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{communicationchannel.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
