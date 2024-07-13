/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3
package app

// noma2
import (
	"encoding/json"
	"fmt"

	"github.com/magicbutton/magic-meeting/utils"
)

func Microsoft(args []string) (*SelectResponse, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("Expected arguments")
	}
	jwt := args[1]
	if jwt == "" {
		return nil, fmt.Errorf("Expected JWT")
	}
	claims, err := utils.DecodeAndValidateMicrosoftJWT(jwt)
	if err != nil {
		return nil, err
	}
	upn := claims["upn"].(string)

	rows := fmt.Sprintf(`{"hello": "%s"}`, upn)

	result := SelectResponse{
		Result: json.RawMessage(rows),
	}

	return &result, nil
}
