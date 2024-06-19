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
import {AccessPassForm} from "./form";

import {AccessPassItem} from "../applogic/model";
export default function UpdateAccessPass(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AccessPassItem>(
      "magic-meeting.accesspass",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const accesspass = readResult.data;
    return (
      <div>{accesspass && 
      <AccessPassForm accesspass={accesspass} editmode="update"/>}
     
      </div>
    );
}
