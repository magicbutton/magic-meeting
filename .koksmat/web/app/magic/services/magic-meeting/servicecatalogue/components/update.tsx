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
import {ServiceCatalogueForm} from "./form";

import {ServiceCatalogueItem} from "../applogic/model";
export default function UpdateServiceCatalogue(props: { id: number }) {
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
      <div>{servicecatalogue && 
      <ServiceCatalogueForm servicecatalogue={servicecatalogue} editmode="update"/>}
     
      </div>
    );
}
