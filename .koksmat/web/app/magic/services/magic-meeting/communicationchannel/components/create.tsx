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
    import {CommunicationChannelForm} from "./form";
    
    import {CommunicationChannelItem} from "../applogic/model";
    export default function CreateCommunicationChannel() {
       
        const communicationchannel = {} as CommunicationChannelItem;
        return (
          <div>{communicationchannel && 
          <CommunicationChannelForm communicationchannel={communicationchannel} editmode="create"/>}
         
          </div>
        );
    }
