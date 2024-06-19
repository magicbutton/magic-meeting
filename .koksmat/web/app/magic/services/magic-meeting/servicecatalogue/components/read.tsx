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
import {ServiceCatalogueItem} from "../applogic/model";


/* yankiebar */

export default function ReadServiceCatalogue(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ServiceCatalogueItem>(
      "magic-meeting.servicecatalogue",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const servicecatalogue = readResult.data;
    return (
      <div>
          
    {servicecatalogue && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{servicecatalogue.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{servicecatalogue.description}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{servicecatalogue.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{servicecatalogue.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{servicecatalogue.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{servicecatalogue.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{servicecatalogue.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
