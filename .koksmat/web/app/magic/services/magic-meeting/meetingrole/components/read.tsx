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
import {MeetingRoleItem} from "../applogic/model";


/* yankiebar */

export default function ReadMeetingRole(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<MeetingRoleItem>(
      "magic-meeting.meetingrole",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const meetingrole = readResult.data;
    return (
      <div>
          
    {meetingrole && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{meetingrole.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{meetingrole.description}</div>
    </div>    <div>
        <div className="font-bold" >user</div>
        <div>{meetingrole.user_id}</div>
    </div>    <div>
        <div className="font-bold" >meeting</div>
        <div>{meetingrole.meeting_id}</div>
    </div>    <div>
        <div className="font-bold" >role</div>
        <div>{meetingrole.role}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{meetingrole.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{meetingrole.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{meetingrole.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{meetingrole.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{meetingrole.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
