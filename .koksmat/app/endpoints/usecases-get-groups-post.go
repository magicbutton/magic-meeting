// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Get Groups
---
*/
package endpoints

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/swaggest/usecase"

	"github.com/magicbutton/magic-meeting/execution"
	"github.com/magicbutton/magic-meeting/schemas"
	"github.com/magicbutton/magic-meeting/utils"
)

func UsecasesGetGroupsPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *schemas.AllGroups) error {

		_, err := execution.ExecutePowerShell("john", "*", "magic-meeting", "05-usecases", "10-get-groups.ps1", "")
		if err != nil {
			return err
		}

		resultingFile := path.Join(utils.WorkDir("magic-meeting"), "all-groups.json")
		data, err := os.ReadFile(resultingFile)
		if err != nil {
			return err
		}
		resultObject := schemas.AllGroups{}
		err = json.Unmarshal(data, &resultObject)
		*output = resultObject
		return err

	})
	u.SetTitle("Get Groups")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Use Cases")
	return u
}
