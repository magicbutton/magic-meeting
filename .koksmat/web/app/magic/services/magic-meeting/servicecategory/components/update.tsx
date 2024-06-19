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
import {ServiceCategoryForm} from "./form";

import {ServiceCategoryItem} from "../applogic/model";
export default function UpdateServiceCategory(props: { id: number }) {
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
      <div>{servicecategory && 
      <ServiceCategoryForm servicecategory={servicecategory} editmode="update"/>}
     
      </div>
    );
}
