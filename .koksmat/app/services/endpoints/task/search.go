/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package task

import (
    "log"

    "github.com/magicbutton/magic-meeting/applogic"
    "github.com/magicbutton/magic-meeting/database"
    "github.com/magicbutton/magic-meeting/services/models/taskmodel"
    . "github.com/magicbutton/magic-meeting/utils"
)

func TaskSearch(query string) (*Page[taskmodel.Task], error) {
    log.Println("Calling Tasksearch")

    return applogic.Search[database.Task, taskmodel.Task]("searchindex", query, applogic.MapTaskOutgoing)

}
