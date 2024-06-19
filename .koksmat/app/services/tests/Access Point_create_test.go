    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
    //generator:  noma4.1
    package tests
    import (
        "testing"
        "github.com/magicbutton/magic-meeting/services/endpoints/accesspoint"
                    "github.com/magicbutton/magic-meeting/services/models/accesspointmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestAccess Pointcreate(t *testing.T) {
                                // transformer v1
            object := accesspointmodel.Access Point{}
         
            result,err := accesspoint.Access PointCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
