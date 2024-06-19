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
import {ServiceForm} from "./form";

import {ServiceItem} from "../applogic/model";
export default function UpdateService(props: { id: number }) {
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
      <div>{service && 
      <ServiceForm service={service} editmode="update"/>}
     
      </div>
    );
}
