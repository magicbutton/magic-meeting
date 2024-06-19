/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package task

import (
    "log"
   
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/taskmodel"

)

func TaskCreate(item taskmodel.Task) (*taskmodel.Task, error) {
    log.Println("Calling Taskcreate")

    return applogic.Create[database.Task, taskmodel.Task](item, applogic.MapTaskIncoming, applogic.MapTaskOutgoing)

}
