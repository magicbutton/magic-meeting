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
import {MeetingRoleForm} from "./form";

import {MeetingRoleItem} from "../applogic/model";
export default function UpdateMeetingRole(props: { id: number }) {
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
      <div>{meetingrole && 
      <MeetingRoleForm meetingrole={meetingrole} editmode="update"/>}
     
      </div>
    );
}
