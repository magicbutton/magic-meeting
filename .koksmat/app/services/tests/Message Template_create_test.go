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
        "github.com/magicbutton/magic-meeting/services/endpoints/messagetemplates"
                    "github.com/magicbutton/magic-meeting/services/models/messagetemplatesmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestMessage Templatecreate(t *testing.T) {
                                // transformer v1
            object := messagetemplatesmodel.Message Template{}
         
            result,err := messagetemplates.Message TemplateCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
