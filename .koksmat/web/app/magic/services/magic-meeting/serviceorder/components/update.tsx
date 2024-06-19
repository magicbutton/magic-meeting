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
import {ServiceOrderForm} from "./form";

import {ServiceOrderItem} from "../applogic/model";
export default function UpdateServiceOrder(props: { id: number }) {
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
      <div>{serviceorder && 
      <ServiceOrderForm serviceorder={serviceorder} editmode="update"/>}
     
      </div>
    );
}
