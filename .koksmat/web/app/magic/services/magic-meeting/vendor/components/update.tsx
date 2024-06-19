/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {VendorForm} from "./form";

import {VendorItem} from "../applogic/model";
export default function UpdateVendor(props: { id: number }) {
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
      <div>{vendor && 
      <VendorForm vendor={vendor} editmode="update"/>}
     
      </div>
    );
}
