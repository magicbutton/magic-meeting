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
        "github.com/magicbutton/magic-meeting/services/endpoints/messagelog"
                    "github.com/magicbutton/magic-meeting/services/models/messagelogmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestMessageLogcreate(t *testing.T) {
                                // transformer v1
            object := messagelogmodel.MessageLog{}
         
            result,err := messagelog.MessageLogCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
