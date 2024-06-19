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
import {APIKeyForm} from "./form";

import {APIKeyItem} from "../applogic/model";
export default function UpdateAPIKey(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<APIKeyItem>(
      "magic-meeting.apikey",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const apikey = readResult.data;
    return (
      <div>{apikey && 
      <APIKeyForm apikey={apikey} editmode="update"/>}
     
      </div>
    );
}
