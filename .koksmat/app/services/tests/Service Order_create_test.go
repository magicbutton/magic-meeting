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
        "github.com/magicbutton/magic-meeting/services/endpoints/serviceorder"
                    "github.com/magicbutton/magic-meeting/services/models/serviceordermodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestService Ordercreate(t *testing.T) {
                                // transformer v1
            object := serviceordermodel.Service Order{}
         
            result,err := serviceorder.Service OrderCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
