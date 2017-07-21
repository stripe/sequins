// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package databasemigrationserviceiface provides an interface to enable mocking the AWS Database Migration Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package databasemigrationserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
)

// DatabaseMigrationServiceAPI provides an interface to enable mocking the
// databasemigrationservice.DatabaseMigrationService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Database Migration Service.
//    func myFunc(svc databasemigrationserviceiface.DatabaseMigrationServiceAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := databasemigrationservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDatabaseMigrationServiceClient struct {
//        databasemigrationserviceiface.DatabaseMigrationServiceAPI
//    }
//    func (m *mockDatabaseMigrationServiceClient) AddTagsToResource(input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDatabaseMigrationServiceClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DatabaseMigrationServiceAPI interface {
	AddTagsToResource(*databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error)
	AddTagsToResourceWithContext(aws.Context, *databasemigrationservice.AddTagsToResourceInput, ...request.Option) (*databasemigrationservice.AddTagsToResourceOutput, error)
	AddTagsToResourceRequest(*databasemigrationservice.AddTagsToResourceInput) (*request.Request, *databasemigrationservice.AddTagsToResourceOutput)

	CreateEndpoint(*databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error)
	CreateEndpointWithContext(aws.Context, *databasemigrationservice.CreateEndpointInput, ...request.Option) (*databasemigrationservice.CreateEndpointOutput, error)
	CreateEndpointRequest(*databasemigrationservice.CreateEndpointInput) (*request.Request, *databasemigrationservice.CreateEndpointOutput)

	CreateEventSubscription(*databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionWithContext(aws.Context, *databasemigrationservice.CreateEventSubscriptionInput, ...request.Option) (*databasemigrationservice.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionRequest(*databasemigrationservice.CreateEventSubscriptionInput) (*request.Request, *databasemigrationservice.CreateEventSubscriptionOutput)

	CreateReplicationInstance(*databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error)
	CreateReplicationInstanceWithContext(aws.Context, *databasemigrationservice.CreateReplicationInstanceInput, ...request.Option) (*databasemigrationservice.CreateReplicationInstanceOutput, error)
	CreateReplicationInstanceRequest(*databasemigrationservice.CreateReplicationInstanceInput) (*request.Request, *databasemigrationservice.CreateReplicationInstanceOutput)

	CreateReplicationSubnetGroup(*databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error)
	CreateReplicationSubnetGroupWithContext(aws.Context, *databasemigrationservice.CreateReplicationSubnetGroupInput, ...request.Option) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error)
	CreateReplicationSubnetGroupRequest(*databasemigrationservice.CreateReplicationSubnetGroupInput) (*request.Request, *databasemigrationservice.CreateReplicationSubnetGroupOutput)

	CreateReplicationTask(*databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error)
	CreateReplicationTaskWithContext(aws.Context, *databasemigrationservice.CreateReplicationTaskInput, ...request.Option) (*databasemigrationservice.CreateReplicationTaskOutput, error)
	CreateReplicationTaskRequest(*databasemigrationservice.CreateReplicationTaskInput) (*request.Request, *databasemigrationservice.CreateReplicationTaskOutput)

	DeleteCertificate(*databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error)
	DeleteCertificateWithContext(aws.Context, *databasemigrationservice.DeleteCertificateInput, ...request.Option) (*databasemigrationservice.DeleteCertificateOutput, error)
	DeleteCertificateRequest(*databasemigrationservice.DeleteCertificateInput) (*request.Request, *databasemigrationservice.DeleteCertificateOutput)

	DeleteEndpoint(*databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error)
	DeleteEndpointWithContext(aws.Context, *databasemigrationservice.DeleteEndpointInput, ...request.Option) (*databasemigrationservice.DeleteEndpointOutput, error)
	DeleteEndpointRequest(*databasemigrationservice.DeleteEndpointInput) (*request.Request, *databasemigrationservice.DeleteEndpointOutput)

	DeleteEventSubscription(*databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionWithContext(aws.Context, *databasemigrationservice.DeleteEventSubscriptionInput, ...request.Option) (*databasemigrationservice.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionRequest(*databasemigrationservice.DeleteEventSubscriptionInput) (*request.Request, *databasemigrationservice.DeleteEventSubscriptionOutput)

	DeleteReplicationInstance(*databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error)
	DeleteReplicationInstanceWithContext(aws.Context, *databasemigrationservice.DeleteReplicationInstanceInput, ...request.Option) (*databasemigrationservice.DeleteReplicationInstanceOutput, error)
	DeleteReplicationInstanceRequest(*databasemigrationservice.DeleteReplicationInstanceInput) (*request.Request, *databasemigrationservice.DeleteReplicationInstanceOutput)

	DeleteReplicationSubnetGroup(*databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error)
	DeleteReplicationSubnetGroupWithContext(aws.Context, *databasemigrationservice.DeleteReplicationSubnetGroupInput, ...request.Option) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error)
	DeleteReplicationSubnetGroupRequest(*databasemigrationservice.DeleteReplicationSubnetGroupInput) (*request.Request, *databasemigrationservice.DeleteReplicationSubnetGroupOutput)

	DeleteReplicationTask(*databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error)
	DeleteReplicationTaskWithContext(aws.Context, *databasemigrationservice.DeleteReplicationTaskInput, ...request.Option) (*databasemigrationservice.DeleteReplicationTaskOutput, error)
	DeleteReplicationTaskRequest(*databasemigrationservice.DeleteReplicationTaskInput) (*request.Request, *databasemigrationservice.DeleteReplicationTaskOutput)

	DescribeAccountAttributes(*databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesWithContext(aws.Context, *databasemigrationservice.DescribeAccountAttributesInput, ...request.Option) (*databasemigrationservice.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesRequest(*databasemigrationservice.DescribeAccountAttributesInput) (*request.Request, *databasemigrationservice.DescribeAccountAttributesOutput)

	DescribeCertificates(*databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error)
	DescribeCertificatesWithContext(aws.Context, *databasemigrationservice.DescribeCertificatesInput, ...request.Option) (*databasemigrationservice.DescribeCertificatesOutput, error)
	DescribeCertificatesRequest(*databasemigrationservice.DescribeCertificatesInput) (*request.Request, *databasemigrationservice.DescribeCertificatesOutput)

	DescribeCertificatesPages(*databasemigrationservice.DescribeCertificatesInput, func(*databasemigrationservice.DescribeCertificatesOutput, bool) bool) error
	DescribeCertificatesPagesWithContext(aws.Context, *databasemigrationservice.DescribeCertificatesInput, func(*databasemigrationservice.DescribeCertificatesOutput, bool) bool, ...request.Option) error

	DescribeConnections(*databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error)
	DescribeConnectionsWithContext(aws.Context, *databasemigrationservice.DescribeConnectionsInput, ...request.Option) (*databasemigrationservice.DescribeConnectionsOutput, error)
	DescribeConnectionsRequest(*databasemigrationservice.DescribeConnectionsInput) (*request.Request, *databasemigrationservice.DescribeConnectionsOutput)

	DescribeConnectionsPages(*databasemigrationservice.DescribeConnectionsInput, func(*databasemigrationservice.DescribeConnectionsOutput, bool) bool) error
	DescribeConnectionsPagesWithContext(aws.Context, *databasemigrationservice.DescribeConnectionsInput, func(*databasemigrationservice.DescribeConnectionsOutput, bool) bool, ...request.Option) error

	DescribeEndpointTypes(*databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error)
	DescribeEndpointTypesWithContext(aws.Context, *databasemigrationservice.DescribeEndpointTypesInput, ...request.Option) (*databasemigrationservice.DescribeEndpointTypesOutput, error)
	DescribeEndpointTypesRequest(*databasemigrationservice.DescribeEndpointTypesInput) (*request.Request, *databasemigrationservice.DescribeEndpointTypesOutput)

	DescribeEndpointTypesPages(*databasemigrationservice.DescribeEndpointTypesInput, func(*databasemigrationservice.DescribeEndpointTypesOutput, bool) bool) error
	DescribeEndpointTypesPagesWithContext(aws.Context, *databasemigrationservice.DescribeEndpointTypesInput, func(*databasemigrationservice.DescribeEndpointTypesOutput, bool) bool, ...request.Option) error

	DescribeEndpoints(*databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error)
	DescribeEndpointsWithContext(aws.Context, *databasemigrationservice.DescribeEndpointsInput, ...request.Option) (*databasemigrationservice.DescribeEndpointsOutput, error)
	DescribeEndpointsRequest(*databasemigrationservice.DescribeEndpointsInput) (*request.Request, *databasemigrationservice.DescribeEndpointsOutput)

	DescribeEndpointsPages(*databasemigrationservice.DescribeEndpointsInput, func(*databasemigrationservice.DescribeEndpointsOutput, bool) bool) error
	DescribeEndpointsPagesWithContext(aws.Context, *databasemigrationservice.DescribeEndpointsInput, func(*databasemigrationservice.DescribeEndpointsOutput, bool) bool, ...request.Option) error

	DescribeEventCategories(*databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesWithContext(aws.Context, *databasemigrationservice.DescribeEventCategoriesInput, ...request.Option) (*databasemigrationservice.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesRequest(*databasemigrationservice.DescribeEventCategoriesInput) (*request.Request, *databasemigrationservice.DescribeEventCategoriesOutput)

	DescribeEventSubscriptions(*databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsWithContext(aws.Context, *databasemigrationservice.DescribeEventSubscriptionsInput, ...request.Option) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsRequest(*databasemigrationservice.DescribeEventSubscriptionsInput) (*request.Request, *databasemigrationservice.DescribeEventSubscriptionsOutput)

	DescribeEventSubscriptionsPages(*databasemigrationservice.DescribeEventSubscriptionsInput, func(*databasemigrationservice.DescribeEventSubscriptionsOutput, bool) bool) error
	DescribeEventSubscriptionsPagesWithContext(aws.Context, *databasemigrationservice.DescribeEventSubscriptionsInput, func(*databasemigrationservice.DescribeEventSubscriptionsOutput, bool) bool, ...request.Option) error

	DescribeEvents(*databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error)
	DescribeEventsWithContext(aws.Context, *databasemigrationservice.DescribeEventsInput, ...request.Option) (*databasemigrationservice.DescribeEventsOutput, error)
	DescribeEventsRequest(*databasemigrationservice.DescribeEventsInput) (*request.Request, *databasemigrationservice.DescribeEventsOutput)

	DescribeEventsPages(*databasemigrationservice.DescribeEventsInput, func(*databasemigrationservice.DescribeEventsOutput, bool) bool) error
	DescribeEventsPagesWithContext(aws.Context, *databasemigrationservice.DescribeEventsInput, func(*databasemigrationservice.DescribeEventsOutput, bool) bool, ...request.Option) error

	DescribeOrderableReplicationInstances(*databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error)
	DescribeOrderableReplicationInstancesWithContext(aws.Context, *databasemigrationservice.DescribeOrderableReplicationInstancesInput, ...request.Option) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error)
	DescribeOrderableReplicationInstancesRequest(*databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*request.Request, *databasemigrationservice.DescribeOrderableReplicationInstancesOutput)

	DescribeOrderableReplicationInstancesPages(*databasemigrationservice.DescribeOrderableReplicationInstancesInput, func(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, bool) bool) error
	DescribeOrderableReplicationInstancesPagesWithContext(aws.Context, *databasemigrationservice.DescribeOrderableReplicationInstancesInput, func(*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, bool) bool, ...request.Option) error

	DescribeRefreshSchemasStatus(*databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error)
	DescribeRefreshSchemasStatusWithContext(aws.Context, *databasemigrationservice.DescribeRefreshSchemasStatusInput, ...request.Option) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error)
	DescribeRefreshSchemasStatusRequest(*databasemigrationservice.DescribeRefreshSchemasStatusInput) (*request.Request, *databasemigrationservice.DescribeRefreshSchemasStatusOutput)

	DescribeReplicationInstances(*databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
	DescribeReplicationInstancesWithContext(aws.Context, *databasemigrationservice.DescribeReplicationInstancesInput, ...request.Option) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
	DescribeReplicationInstancesRequest(*databasemigrationservice.DescribeReplicationInstancesInput) (*request.Request, *databasemigrationservice.DescribeReplicationInstancesOutput)

	DescribeReplicationInstancesPages(*databasemigrationservice.DescribeReplicationInstancesInput, func(*databasemigrationservice.DescribeReplicationInstancesOutput, bool) bool) error
	DescribeReplicationInstancesPagesWithContext(aws.Context, *databasemigrationservice.DescribeReplicationInstancesInput, func(*databasemigrationservice.DescribeReplicationInstancesOutput, bool) bool, ...request.Option) error

	DescribeReplicationSubnetGroups(*databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error)
	DescribeReplicationSubnetGroupsWithContext(aws.Context, *databasemigrationservice.DescribeReplicationSubnetGroupsInput, ...request.Option) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error)
	DescribeReplicationSubnetGroupsRequest(*databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*request.Request, *databasemigrationservice.DescribeReplicationSubnetGroupsOutput)

	DescribeReplicationSubnetGroupsPages(*databasemigrationservice.DescribeReplicationSubnetGroupsInput, func(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, bool) bool) error
	DescribeReplicationSubnetGroupsPagesWithContext(aws.Context, *databasemigrationservice.DescribeReplicationSubnetGroupsInput, func(*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, bool) bool, ...request.Option) error

	DescribeReplicationTasks(*databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error)
	DescribeReplicationTasksWithContext(aws.Context, *databasemigrationservice.DescribeReplicationTasksInput, ...request.Option) (*databasemigrationservice.DescribeReplicationTasksOutput, error)
	DescribeReplicationTasksRequest(*databasemigrationservice.DescribeReplicationTasksInput) (*request.Request, *databasemigrationservice.DescribeReplicationTasksOutput)

	DescribeReplicationTasksPages(*databasemigrationservice.DescribeReplicationTasksInput, func(*databasemigrationservice.DescribeReplicationTasksOutput, bool) bool) error
	DescribeReplicationTasksPagesWithContext(aws.Context, *databasemigrationservice.DescribeReplicationTasksInput, func(*databasemigrationservice.DescribeReplicationTasksOutput, bool) bool, ...request.Option) error

	DescribeSchemas(*databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error)
	DescribeSchemasWithContext(aws.Context, *databasemigrationservice.DescribeSchemasInput, ...request.Option) (*databasemigrationservice.DescribeSchemasOutput, error)
	DescribeSchemasRequest(*databasemigrationservice.DescribeSchemasInput) (*request.Request, *databasemigrationservice.DescribeSchemasOutput)

	DescribeSchemasPages(*databasemigrationservice.DescribeSchemasInput, func(*databasemigrationservice.DescribeSchemasOutput, bool) bool) error
	DescribeSchemasPagesWithContext(aws.Context, *databasemigrationservice.DescribeSchemasInput, func(*databasemigrationservice.DescribeSchemasOutput, bool) bool, ...request.Option) error

	DescribeTableStatistics(*databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error)
	DescribeTableStatisticsWithContext(aws.Context, *databasemigrationservice.DescribeTableStatisticsInput, ...request.Option) (*databasemigrationservice.DescribeTableStatisticsOutput, error)
	DescribeTableStatisticsRequest(*databasemigrationservice.DescribeTableStatisticsInput) (*request.Request, *databasemigrationservice.DescribeTableStatisticsOutput)

	DescribeTableStatisticsPages(*databasemigrationservice.DescribeTableStatisticsInput, func(*databasemigrationservice.DescribeTableStatisticsOutput, bool) bool) error
	DescribeTableStatisticsPagesWithContext(aws.Context, *databasemigrationservice.DescribeTableStatisticsInput, func(*databasemigrationservice.DescribeTableStatisticsOutput, bool) bool, ...request.Option) error

	ImportCertificate(*databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error)
	ImportCertificateWithContext(aws.Context, *databasemigrationservice.ImportCertificateInput, ...request.Option) (*databasemigrationservice.ImportCertificateOutput, error)
	ImportCertificateRequest(*databasemigrationservice.ImportCertificateInput) (*request.Request, *databasemigrationservice.ImportCertificateOutput)

	ListTagsForResource(*databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *databasemigrationservice.ListTagsForResourceInput, ...request.Option) (*databasemigrationservice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*databasemigrationservice.ListTagsForResourceInput) (*request.Request, *databasemigrationservice.ListTagsForResourceOutput)

	ModifyEndpoint(*databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error)
	ModifyEndpointWithContext(aws.Context, *databasemigrationservice.ModifyEndpointInput, ...request.Option) (*databasemigrationservice.ModifyEndpointOutput, error)
	ModifyEndpointRequest(*databasemigrationservice.ModifyEndpointInput) (*request.Request, *databasemigrationservice.ModifyEndpointOutput)

	ModifyEventSubscription(*databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionWithContext(aws.Context, *databasemigrationservice.ModifyEventSubscriptionInput, ...request.Option) (*databasemigrationservice.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionRequest(*databasemigrationservice.ModifyEventSubscriptionInput) (*request.Request, *databasemigrationservice.ModifyEventSubscriptionOutput)

	ModifyReplicationInstance(*databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error)
	ModifyReplicationInstanceWithContext(aws.Context, *databasemigrationservice.ModifyReplicationInstanceInput, ...request.Option) (*databasemigrationservice.ModifyReplicationInstanceOutput, error)
	ModifyReplicationInstanceRequest(*databasemigrationservice.ModifyReplicationInstanceInput) (*request.Request, *databasemigrationservice.ModifyReplicationInstanceOutput)

	ModifyReplicationSubnetGroup(*databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error)
	ModifyReplicationSubnetGroupWithContext(aws.Context, *databasemigrationservice.ModifyReplicationSubnetGroupInput, ...request.Option) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error)
	ModifyReplicationSubnetGroupRequest(*databasemigrationservice.ModifyReplicationSubnetGroupInput) (*request.Request, *databasemigrationservice.ModifyReplicationSubnetGroupOutput)

	ModifyReplicationTask(*databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error)
	ModifyReplicationTaskWithContext(aws.Context, *databasemigrationservice.ModifyReplicationTaskInput, ...request.Option) (*databasemigrationservice.ModifyReplicationTaskOutput, error)
	ModifyReplicationTaskRequest(*databasemigrationservice.ModifyReplicationTaskInput) (*request.Request, *databasemigrationservice.ModifyReplicationTaskOutput)

	RefreshSchemas(*databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error)
	RefreshSchemasWithContext(aws.Context, *databasemigrationservice.RefreshSchemasInput, ...request.Option) (*databasemigrationservice.RefreshSchemasOutput, error)
	RefreshSchemasRequest(*databasemigrationservice.RefreshSchemasInput) (*request.Request, *databasemigrationservice.RefreshSchemasOutput)

	ReloadTables(*databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error)
	ReloadTablesWithContext(aws.Context, *databasemigrationservice.ReloadTablesInput, ...request.Option) (*databasemigrationservice.ReloadTablesOutput, error)
	ReloadTablesRequest(*databasemigrationservice.ReloadTablesInput) (*request.Request, *databasemigrationservice.ReloadTablesOutput)

	RemoveTagsFromResource(*databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceWithContext(aws.Context, *databasemigrationservice.RemoveTagsFromResourceInput, ...request.Option) (*databasemigrationservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceRequest(*databasemigrationservice.RemoveTagsFromResourceInput) (*request.Request, *databasemigrationservice.RemoveTagsFromResourceOutput)

	StartReplicationTask(*databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error)
	StartReplicationTaskWithContext(aws.Context, *databasemigrationservice.StartReplicationTaskInput, ...request.Option) (*databasemigrationservice.StartReplicationTaskOutput, error)
	StartReplicationTaskRequest(*databasemigrationservice.StartReplicationTaskInput) (*request.Request, *databasemigrationservice.StartReplicationTaskOutput)

	StopReplicationTask(*databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error)
	StopReplicationTaskWithContext(aws.Context, *databasemigrationservice.StopReplicationTaskInput, ...request.Option) (*databasemigrationservice.StopReplicationTaskOutput, error)
	StopReplicationTaskRequest(*databasemigrationservice.StopReplicationTaskInput) (*request.Request, *databasemigrationservice.StopReplicationTaskOutput)

	TestConnection(*databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error)
	TestConnectionWithContext(aws.Context, *databasemigrationservice.TestConnectionInput, ...request.Option) (*databasemigrationservice.TestConnectionOutput, error)
	TestConnectionRequest(*databasemigrationservice.TestConnectionInput) (*request.Request, *databasemigrationservice.TestConnectionOutput)
}

var _ DatabaseMigrationServiceAPI = (*databasemigrationservice.DatabaseMigrationService)(nil)
