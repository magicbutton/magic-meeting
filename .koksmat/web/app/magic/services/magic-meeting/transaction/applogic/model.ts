    
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
// Transaction
export interface TransactionItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    amount : number ;
    currency : string ;
    datetime : string ;
    status : string ;
    account_id : number ;

}


// Transaction
export const TransactionSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    amount : z.number(), 
    currency : z.string(), 
    datetime : z.string(), 
    status : z.string(), 
    account_id : z.number(), 

});

