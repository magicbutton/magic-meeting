    
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
// ServiceOrder
export interface ServiceOrderItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    deliverydate : string ;
    deliverto_id : number ;
    status : string ;
    payment_id : number ;

}


// ServiceOrder
export const ServiceOrderSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    deliverydate : z.string().optional(), 
    deliverto_id : z.number().optional(), 
    status : z.string(), 
    payment_id : z.number().optional(), 

});

