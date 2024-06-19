            /*
            File have been automatically created. To prevent the file from getting overwritten
            set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
            ---
            keep: false
            ---
            */
            //generator:  noma3.delete.v2
            package productionorder
            
            import (
                "log"
                "strconv"
                "github.com/magicbutton/magic-meeting/applogic"
                "github.com/magicbutton/magic-meeting/database"
                "github.com/magicbutton/magic-meeting/services/models/productionordermodel"
            
            )
            
            func ProductionOrderDelete(arg0 string) ( error) {
                id,_ := strconv.Atoi(arg0)
                log.Println("Calling ProductionOrderdelete")
            
                return applogic.Delete[database.ProductionOrder, productionordermodel.ProductionOrder](id)
            
            }
