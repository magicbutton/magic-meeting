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
import {MeetingForm} from "./form";

import {MeetingItem} from "../applogic/model";
export default function UpdateMeeting(props: { id: number }) {
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
      <div>{meeting && 
      <MeetingForm meeting={meeting} editmode="update"/>}
     
      </div>
    );
}
