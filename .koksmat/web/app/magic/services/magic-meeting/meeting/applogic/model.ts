    
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
// Meeting
export interface MeetingItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    start : string ;
    duration : number ;
    location : string ;
    status : string ;
    exchangereference : string ;
    exchangestatus : string ;
    teamsreference : string ;
    teamsstatus : string ;

}


// Meeting
export const MeetingSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    start : z.string(), 
    duration : z.number(), 
    location : z.string(), 
    status : z.string(), 
    exchangereference : z.string().optional(), 
    exchangestatus : z.string().optional(), 
    teamsreference : z.string().optional(), 
    teamsstatus : z.string().optional(), 

});

