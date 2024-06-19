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
import {ServiceCategoryItem} from "../applogic/model";


/* yankiebar */

export default function ReadServiceCategory(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ServiceCategoryItem>(
      "magic-meeting.servicecategory",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const servicecategory = readResult.data;
    return (
      <div>
          
    {servicecategory && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{servicecategory.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{servicecategory.description}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{servicecategory.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{servicecategory.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{servicecategory.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{servicecategory.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{servicecategory.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
