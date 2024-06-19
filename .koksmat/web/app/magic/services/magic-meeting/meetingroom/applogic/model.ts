    
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
// MeetingRoom
export interface MeetingRoomItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    email : string ;
    capacity : number ;
    features : string ;
    floor_id : number ;

}


// MeetingRoom
export const MeetingRoomSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    email : z.string(), 
    capacity : z.number(), 
    features : z.string().optional(), 
    floor_id : z.number(), 

});

