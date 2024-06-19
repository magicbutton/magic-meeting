    
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
// AccessPass
export interface AccessPassItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    visitor_id : number ;
    validfrom : string ;
    validto : string ;
    status : string ;

}


// AccessPass
export const AccessPassSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    visitor_id : z.number(), 
    validfrom : z.string(), 
    validto : z.string(), 
    status : z.string(), 

});

