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
import {ServiceItem} from "../applogic/model";


/* yankiebar */

export default function ReadService(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ServiceItem>(
      "magic-meeting.service",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const service = readResult.data;
    return (
      <div>
          
    {service && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{service.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{service.description}</div>
    </div>    <div>
        <div className="font-bold" >price</div>
        <div>{service.price}</div>
    </div>    <div>
        <div className="font-bold" >currency</div>
        <div>{service.currency}</div>
    </div>    <div>
        <div className="font-bold" >servicecategory</div>
        <div>{service.servicecategory_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{service.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{service.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{service.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{service.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{service.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
