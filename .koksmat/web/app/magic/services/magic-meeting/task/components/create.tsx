    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/app/koksmat/useservice";
    import { useState } from "react";
    import {TaskForm} from "./form";
    
    import {TaskItem} from "../applogic/model";
    export default function CreateTask() {
       
        const task = {} as TaskItem;
        return (
          <div>{task && 
          <TaskForm task={task} editmode="create"/>}
         
          </div>
        );
    }
