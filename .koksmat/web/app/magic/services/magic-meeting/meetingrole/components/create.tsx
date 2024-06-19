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
    import {MeetingRoleForm} from "./form";
    
    import {MeetingRoleItem} from "../applogic/model";
    export default function CreateMeetingRole() {
       
        const meetingrole = {} as MeetingRoleItem;
        return (
          <div>{meetingrole && 
          <MeetingRoleForm meetingrole={meetingrole} editmode="create"/>}
         
          </div>
        );
    }
