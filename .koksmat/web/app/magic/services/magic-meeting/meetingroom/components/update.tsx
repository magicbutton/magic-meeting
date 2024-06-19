/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {MeetingRoomForm} from "./form";

import {MeetingRoomItem} from "../applogic/model";
export default function UpdateMeetingRoom(props: { id: number }) {
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
      <div>{meetingroom && 
      <MeetingRoomForm meetingroom={meetingroom} editmode="update"/>}
     
      </div>
    );
}
