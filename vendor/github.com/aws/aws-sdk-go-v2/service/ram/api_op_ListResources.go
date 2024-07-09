// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resources that you added to a resource share or the resources that
// are shared with you.
func (c *Client) ListResources(ctx context.Context, params *ListResourcesInput, optFns ...func(*Options)) (*ListResourcesOutput, error) {
	if params == nil {
		params = &ListResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResources", params, optFns, c.addOperationListResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesInput struct {

	// Specifies that you want to list only the resource shares that match the
	// following:
	//
	//   - SELF – resources that your account shares with other accounts
	//
	//   - OTHER-ACCOUNTS – resources that other accounts share with your account
	//
	// This member is required.
	ResourceOwner types.ResourceOwner

	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	// Specifies that you want to list only the resource shares that are associated
	// with the specified principal.
	Principal *string

	// Specifies that you want to list only the resource shares that include resources
	// with the specified [Amazon Resource Names (ARNs)].
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	ResourceArns []string

	// Specifies that you want the results to include only resources that have the
	// specified scope.
	//
	//   - ALL – the results include both global and regional resources or resource
	//   types.
	//
	//   - GLOBAL – the results include only global resources or resource types.
	//
	//   - REGIONAL – the results include only regional resources or resource types.
	//
	// The default value is ALL .
	ResourceRegionScope types.ResourceRegionScopeFilter

	// Specifies that you want to list only resources in the resource shares
	// identified by the specified [Amazon Resource Names (ARNs)].
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	ResourceShareArns []string

	// Specifies that you want to list only the resource shares that include resources
	// of the specified resource type.
	//
	// For valid values, query the ListResourceTypes operation.
	ResourceType *string

	noSmithyDocumentSerde
}

type ListResourcesOutput struct {

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// An array of objects that contain information about the resources.
	Resources []types.Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResources"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpListResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResources(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListResourcesPaginatorOptions is the paginator options for ListResources
type ListResourcesPaginatorOptions struct {
	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourcesPaginator is a paginator for ListResources
type ListResourcesPaginator struct {
	options   ListResourcesPaginatorOptions
	client    ListResourcesAPIClient
	params    *ListResourcesInput
	nextToken *string
	firstPage bool
}

// NewListResourcesPaginator returns a new ListResourcesPaginator
func NewListResourcesPaginator(client ListResourcesAPIClient, params *ListResourcesInput, optFns ...func(*ListResourcesPaginatorOptions)) *ListResourcesPaginator {
	if params == nil {
		params = &ListResourcesInput{}
	}

	options := ListResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResources page.
func (p *ListResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListResources(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListResourcesAPIClient is a client that implements the ListResources operation.
type ListResourcesAPIClient interface {
	ListResources(context.Context, *ListResourcesInput, ...func(*Options)) (*ListResourcesOutput, error)
}

var _ ListResourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResources",
	}
}
