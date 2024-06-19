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
import {AccessPointForm} from "./form";

import {AccessPointItem} from "../applogic/model";
export default function UpdateAccessPoint(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AccessPointItem>(
      "magic-meeting.accesspoint",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const accesspoint = readResult.data;
    return (
      <div>{accesspoint && 
      <AccessPointForm accesspoint={accesspoint} editmode="update"/>}
     
      </div>
    );
}
