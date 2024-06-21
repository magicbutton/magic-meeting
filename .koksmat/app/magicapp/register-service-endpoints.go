/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
package magicapp

import (
	"github.com/magicbutton/magic-meeting/services"
	"github.com/nats-io/nats.go/micro"
)

func RegisterServiceEndpoints(root micro.Group) {
    root.AddEndpoint("user", micro.HandlerFunc(services.HandleUserRequests))
        root.AddEndpoint("apikey", micro.HandlerFunc(services.HandleAPIKeyRequests))
        root.AddEndpoint("communicationchannel", micro.HandlerFunc(services.HandleCommunicationChannelRequests))
        root.AddEndpoint("messagetemplates", micro.HandlerFunc(services.HandleMessageTemplateRequests))
        root.AddEndpoint("country", micro.HandlerFunc(services.HandleCountryRequests))
        root.AddEndpoint("site", micro.HandlerFunc(services.HandleSiteRequests))
        root.AddEndpoint("vendor", micro.HandlerFunc(services.HandleVendorRequests))
        root.AddEndpoint("task", micro.HandlerFunc(services.HandleTaskRequests))
        root.AddEndpoint("signal", micro.HandlerFunc(services.HandleSignalRequests))
        root.AddEndpoint("servicecatalogue", micro.HandlerFunc(services.HandleServiceCatalogueRequests))
        root.AddEndpoint("serviceorder", micro.HandlerFunc(services.HandleServiceOrderRequests))
        root.AddEndpoint("service", micro.HandlerFunc(services.HandleServiceRequests))
        root.AddEndpoint("servicecategory", micro.HandlerFunc(services.HandleServiceCategoryRequests))
        root.AddEndpoint("productionorder", micro.HandlerFunc(services.HandleProductionOrderRequests))
        root.AddEndpoint("meeting", micro.HandlerFunc(services.HandleMeetingRequests))
        root.AddEndpoint("accesspass", micro.HandlerFunc(services.HandleAccessPassRequests))
        root.AddEndpoint("accesspoint", micro.HandlerFunc(services.HandleAccessPointRequests))
        root.AddEndpoint("auditlog", micro.HandlerFunc(services.HandleAuditLogRequests))
        root.AddEndpoint("visitor", micro.HandlerFunc(services.HandleVisitorRequests))
        root.AddEndpoint("messagelog", micro.HandlerFunc(services.HandleMessageLogRequests))
        root.AddEndpoint("transaction", micro.HandlerFunc(services.HandleTransactionRequests))
        root.AddEndpoint("meetingroom", micro.HandlerFunc(services.HandleMeetingRoomRequests))
        root.AddEndpoint("building", micro.HandlerFunc(services.HandleBuildingRequests))
        root.AddEndpoint("importdata", micro.HandlerFunc(services.HandleImportDataRequests))
        root.AddEndpoint("tenant", micro.HandlerFunc(services.HandleTenantRequests))
        root.AddEndpoint("meetingrole", micro.HandlerFunc(services.HandleMeetingRoleRequests))
    }
