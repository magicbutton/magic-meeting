    
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
// APIKey
export interface APIKeyItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    key : string ;
    user_id : number ;
    validto : string ;

}


// APIKey
export const APIKeySchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    key : z.string(), 
    user_id : z.number(), 
    validto : z.string(), 

});

