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
import {TaskItem} from "../applogic/model";


/* yankiebar */

export default function ReadTask(props: { id: number }) {
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
      <div>
          
    {task && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{task.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{task.description}</div>
    </div>    <div>
        <div className="font-bold" >starttime</div>
        <div>{task.starttime}</div>
    </div>    <div>
        <div className="font-bold" >location</div>
        <div>{task.location}</div>
    </div>    <div>
        <div className="font-bold" >responsible</div>
        <div>{task.responsible_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{task.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{task.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{task.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{task.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{task.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
