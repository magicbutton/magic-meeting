    
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/       
"use client";
import { z } from "zod";
// spunk
// MeetingRole
export interface MeetingRoleItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    user_id : number ;
    meeting_id : number ;
    role : string ;

}


// MeetingRole
export const MeetingRoleSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    user_id : z.number(), 
    meeting_id : z.number(), 
    role : z.string(), 

});

