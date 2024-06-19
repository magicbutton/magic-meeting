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
import {ProductionOrderItem} from "../applogic/model";


/* yankiebar */

export default function ReadProductionOrder(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ProductionOrderItem>(
      "magic-meeting.productionorder",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const productionorder = readResult.data;
    return (
      <div>
          
    {productionorder && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{productionorder.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{productionorder.description}</div>
    </div>    <div>
        <div className="font-bold" >deliverydate</div>
        <div>{productionorder.deliverydate}</div>
    </div>    <div>
        <div className="font-bold" >deliverto</div>
        <div>{productionorder.deliverto_id}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{productionorder.status}</div>
    </div>    <div>
        <div className="font-bold" >payment</div>
        <div>{productionorder.payment_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{productionorder.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{productionorder.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{productionorder.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{productionorder.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{productionorder.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
