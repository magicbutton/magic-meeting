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
import {VendorItem} from "../applogic/model";


/* yankiebar */

export default function ReadVendor(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<VendorItem>(
      "magic-meeting.vendor",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const vendor = readResult.data;
    return (
      <div>
          
    {vendor && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{vendor.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{vendor.description}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{vendor.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{vendor.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{vendor.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{vendor.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{vendor.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
