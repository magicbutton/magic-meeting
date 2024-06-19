    
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
// Visitor
export interface VisitorItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    email : string ;
    phone : string ;
    company : string ;
    purpose : string ;
    host_id : number ;
    status : string ;

}


// Visitor
export const VisitorSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    email : z.string(), 
    phone : z.string(), 
    company : z.string(), 
    purpose : z.string(), 
    host_id : z.number(), 
    status : z.string(), 

});

