/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package task

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/taskmodel"

)

func TaskRead(arg0 string) (*taskmodel.Task, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Taskread")

    return applogic.Read[database.Task, taskmodel.Task](id, applogic.MapTaskOutgoing)

}
