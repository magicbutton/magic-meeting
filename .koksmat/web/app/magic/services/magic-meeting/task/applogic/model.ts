    
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
// Task
export interface TaskItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    starttime : string ;
    location : string ;
    responsible_id : number ;

}


// Task
export const TaskSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    starttime : z.string(), 
    location : z.string(), 
    responsible_id : z.number(), 

});

