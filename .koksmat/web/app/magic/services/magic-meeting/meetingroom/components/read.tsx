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
import {MeetingRoomItem} from "../applogic/model";


/* yankiebar */

export default function ReadMeetingRoom(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<MeetingRoomItem>(
      "magic-meeting.meetingroom",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const meetingroom = readResult.data;
    return (
      <div>
          
    {meetingroom && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{meetingroom.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{meetingroom.description}</div>
    </div>    <div>
        <div className="font-bold" >email</div>
        <div>{meetingroom.email}</div>
    </div>    <div>
        <div className="font-bold" >capacity</div>
        <div>{meetingroom.capacity}</div>
    </div>    <div>
        <div className="font-bold" >features</div>
        <div>{meetingroom.features}</div>
    </div>    <div>
        <div className="font-bold" >floor</div>
        <div>{meetingroom.floor_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{meetingroom.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{meetingroom.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{meetingroom.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{meetingroom.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{meetingroom.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
