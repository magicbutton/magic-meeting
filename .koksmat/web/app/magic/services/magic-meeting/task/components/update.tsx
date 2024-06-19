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
import {TaskForm} from "./form";

import {TaskItem} from "../applogic/model";
export default function UpdateTask(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TaskItem>(
      "magic-meeting.task",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const task = readResult.data;
    return (
      <div>{task && 
      <TaskForm task={task} editmode="update"/>}
     
      </div>
    );
}
