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
        "github.com/magicbutton/magic-meeting/services/endpoints/communicationchannel"
                    "github.com/magicbutton/magic-meeting/services/models/communicationchannelmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestCommunicationChannelcreate(t *testing.T) {
                                // transformer v1
            object := communicationchannelmodel.CommunicationChannel{}
         
            result,err := communicationchannel.CommunicationChannelCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
