/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//version: pølsevogn2
package database

import (
	"time"
    
	"github.com/uptrace/bun"
)

type Meeting struct {
	bun.BaseModel `bun:"table:meeting,alias:meeting"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy      string `bun:"created_by,"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedBy      string `bun:"updated_by,"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `bun:"tenant"`
    Searchindex string `bun:"searchindex"`
    Name string `bun:"name"`
    Description string `bun:"description"`
    Start time.Time `bun:"start"`
    Duration int `bun:"duration"`
    Location string `bun:"location"`
    Status string `bun:"status"`
    Exchangereference string `bun:"exchangereference"`
    Exchangestatus string `bun:"exchangestatus"`
    Teamsreference string `bun:"teamsreference"`
    Teamsstatus string `bun:"teamsstatus"`

}

