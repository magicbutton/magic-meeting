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
        
        "github.com/stretchr/testify/assert"
    )
    
    func TestMessage Templatedelete(t *testing.T) {
                // noma4.1.1
        
        err := messagetemplates.Message TemplateDelete(".")
        if err != nil {
            t.Errorf("Error %s", err)
        }
        assert.True(t, true) // for additional tests
       
        
    
    }
    
