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
    import {MeetingRoomForm} from "./form";
    
    import {MeetingRoomItem} from "../applogic/model";
    export default function CreateMeetingRoom() {
       
        const meetingroom = {} as MeetingRoomItem;
        return (
          <div>{meetingroom && 
          <MeetingRoomForm meetingroom={meetingroom} editmode="create"/>}
         
          </div>
        );
    }
