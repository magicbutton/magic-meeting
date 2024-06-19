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
    import {MeetingForm} from "./form";
    
    import {MeetingItem} from "../applogic/model";
    export default function CreateMeeting() {
       
        const meeting = {} as MeetingItem;
        return (
          <div>{meeting && 
          <MeetingForm meeting={meeting} editmode="create"/>}
         
          </div>
        );
    }
