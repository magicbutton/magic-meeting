            /*
            File have been automatically created. To prevent the file from getting overwritten
            set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
            ---
            keep: false
            ---
            */
            //generator:  noma3.delete.v2
            package company
            
            import (
                "log"
                "strconv"
                "github.com/magicbutton/magic-people/applogic"
                "github.com/magicbutton/magic-people/database"
                "github.com/magicbutton/magic-people/services/models/companymodel"
            
            )
            
            func CompanyDelete(arg0 string) ( error) {
                id,_ := strconv.Atoi(arg0)
                log.Println("Calling Companydelete")
            
                return applogic.Delete[database.Company, companymodel.Company](id)
            
            }
