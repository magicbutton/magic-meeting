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
        "github.com/magicbutton/magic-meeting/services/endpoints/servicecatalogue"
                    "github.com/magicbutton/magic-meeting/services/models/servicecataloguemodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestService Cataloguecreate(t *testing.T) {
                                // transformer v1
            object := servicecataloguemodel.Service Catalogue{}
         
            result,err := servicecatalogue.Service CatalogueCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
