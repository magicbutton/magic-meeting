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
        "github.com/magicbutton/magic-meeting/services/endpoints/visitor"
                    "github.com/magicbutton/magic-meeting/services/models/visitormodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestVisitorcreate(t *testing.T) {
                                // transformer v1
            object := visitormodel.Visitor{}
         
            result,err := visitor.VisitorCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
