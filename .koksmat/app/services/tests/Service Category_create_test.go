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
        "github.com/magicbutton/magic-meeting/services/endpoints/servicecategory"
                    "github.com/magicbutton/magic-meeting/services/models/servicecategorymodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestService Categorycreate(t *testing.T) {
                                // transformer v1
            object := servicecategorymodel.Service Category{}
         
            result,err := servicecategory.Service CategoryCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
