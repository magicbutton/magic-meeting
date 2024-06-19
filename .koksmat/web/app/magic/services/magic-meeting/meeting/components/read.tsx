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
import {MeetingItem} from "../applogic/model";


/* yankiebar */

export default function ReadMeeting(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<MeetingItem>(
      "magic-meeting.meeting",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const meeting = readResult.data;
    return (
      <div>
          
    {meeting && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{meeting.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{meeting.description}</div>
    </div>    <div>
        <div className="font-bold" >start</div>
        <div>{meeting.start}</div>
    </div>    <div>
        <div className="font-bold" >duration</div>
        <div>{meeting.duration}</div>
    </div>    <div>
        <div className="font-bold" >location</div>
        <div>{meeting.location}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{meeting.status}</div>
    </div>    <div>
        <div className="font-bold" >exchangereference</div>
        <div>{meeting.exchangereference}</div>
    </div>    <div>
        <div className="font-bold" >exchangestatus</div>
        <div>{meeting.exchangestatus}</div>
    </div>    <div>
        <div className="font-bold" >teamsreference</div>
        <div>{meeting.teamsreference}</div>
    </div>    <div>
        <div className="font-bold" >teamsstatus</div>
        <div>{meeting.teamsstatus}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{meeting.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{meeting.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{meeting.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{meeting.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{meeting.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
