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
import {ServiceOrderItem} from "../applogic/model";


/* yankiebar */

export default function ReadServiceOrder(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ServiceOrderItem>(
      "magic-meeting.serviceorder",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const serviceorder = readResult.data;
    return (
      <div>
          
    {serviceorder && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{serviceorder.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{serviceorder.description}</div>
    </div>    <div>
        <div className="font-bold" >deliverydate</div>
        <div>{serviceorder.deliverydate}</div>
    </div>    <div>
        <div className="font-bold" >deliverto</div>
        <div>{serviceorder.deliverto_id}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{serviceorder.status}</div>
    </div>    <div>
        <div className="font-bold" >payment</div>
        <div>{serviceorder.payment_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{serviceorder.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{serviceorder.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{serviceorder.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{serviceorder.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{serviceorder.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
