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
import {ProductionOrderForm} from "./form";

import {ProductionOrderItem} from "../applogic/model";
export default function UpdateProductionOrder(props: { id: number }) {
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
      <div>{productionorder && 
      <ProductionOrderForm productionorder={productionorder} editmode="update"/>}
     
      </div>
    );
}
