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
    import {APIKeyForm} from "./form";
    
    import {APIKeyItem} from "../applogic/model";
    export default function CreateAPIKey() {
       
        const apikey = {} as APIKeyItem;
        return (
          <div>{apikey && 
          <APIKeyForm apikey={apikey} editmode="create"/>}
         
          </div>
        );
    }
