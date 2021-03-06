// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: AuthorizationRoles
//
// Create and update role information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-splunk/sdk/go/splunk"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := splunk.NewAuthorizationRoles(ctx, "role01", &splunk.AuthorizationRolesArgs{
// 			Capabilities: pulumi.StringArray{
// 				pulumi.String("accelerate_datamodel"),
// 				pulumi.String("change_authentication"),
// 				pulumi.String("restart_splunkd"),
// 			},
// 			DefaultApp: pulumi.String("search"),
// 			ImportedRoles: pulumi.StringArray{
// 				pulumi.String("power"),
// 				pulumi.String("user"),
// 			},
// 			SearchIndexesAlloweds: pulumi.StringArray{
// 				pulumi.String("_audit"),
// 				pulumi.String("_internal"),
// 				pulumi.String("main"),
// 			},
// 			SearchIndexesDefaults: pulumi.StringArray{
// 				pulumi.String("_audit"),
// 				pulumi.String("_internal"),
// 				pulumi.String("main"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AuthorizationRoles struct {
	pulumi.CustomResourceState

	// List of capabilities assigned to role.
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`
	// Maximum number of concurrently running real-time searches that all members of this role can have.
	CumulativeRealtimeSearchJobsQuota pulumi.IntOutput `pulumi:"cumulativeRealtimeSearchJobsQuota"`
	// Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
	CumulativeSearchJobsQuota pulumi.IntOutput `pulumi:"cumulativeSearchJobsQuota"`
	// Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
	DefaultApp pulumi.StringOutput `pulumi:"defaultApp"`
	// List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
	ImportedRoles pulumi.StringArrayOutput `pulumi:"importedRoles"`
	// The name of the user role to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
	RealtimeSearchJobsQuota pulumi.IntOutput `pulumi:"realtimeSearchJobsQuota"`
	// Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
	SearchDiskQuota pulumi.IntOutput `pulumi:"searchDiskQuota"`
	// Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
	SearchFilter pulumi.StringOutput `pulumi:"searchFilter"`
	// List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
	SearchIndexesAlloweds pulumi.StringArrayOutput `pulumi:"searchIndexesAlloweds"`
	// List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
	SearchIndexesDefaults pulumi.StringArrayOutput `pulumi:"searchIndexesDefaults"`
	// The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
	SearchJobsQuota pulumi.IntOutput `pulumi:"searchJobsQuota"`
	// Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
	SearchTimeWin pulumi.IntOutput `pulumi:"searchTimeWin"`
}

// NewAuthorizationRoles registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationRoles(ctx *pulumi.Context,
	name string, args *AuthorizationRolesArgs, opts ...pulumi.ResourceOption) (*AuthorizationRoles, error) {
	if args == nil {
		args = &AuthorizationRolesArgs{}
	}

	var resource AuthorizationRoles
	err := ctx.RegisterResource("splunk:index/authorizationRoles:AuthorizationRoles", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationRoles gets an existing AuthorizationRoles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationRoles(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationRolesState, opts ...pulumi.ResourceOption) (*AuthorizationRoles, error) {
	var resource AuthorizationRoles
	err := ctx.ReadResource("splunk:index/authorizationRoles:AuthorizationRoles", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationRoles resources.
type authorizationRolesState struct {
	// List of capabilities assigned to role.
	Capabilities []string `pulumi:"capabilities"`
	// Maximum number of concurrently running real-time searches that all members of this role can have.
	CumulativeRealtimeSearchJobsQuota *int `pulumi:"cumulativeRealtimeSearchJobsQuota"`
	// Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
	CumulativeSearchJobsQuota *int `pulumi:"cumulativeSearchJobsQuota"`
	// Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
	DefaultApp *string `pulumi:"defaultApp"`
	// List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
	ImportedRoles []string `pulumi:"importedRoles"`
	// The name of the user role to create.
	Name *string `pulumi:"name"`
	// Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
	RealtimeSearchJobsQuota *int `pulumi:"realtimeSearchJobsQuota"`
	// Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
	SearchDiskQuota *int `pulumi:"searchDiskQuota"`
	// Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
	SearchFilter *string `pulumi:"searchFilter"`
	// List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
	SearchIndexesAlloweds []string `pulumi:"searchIndexesAlloweds"`
	// List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
	SearchIndexesDefaults []string `pulumi:"searchIndexesDefaults"`
	// The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
	SearchJobsQuota *int `pulumi:"searchJobsQuota"`
	// Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
	SearchTimeWin *int `pulumi:"searchTimeWin"`
}

type AuthorizationRolesState struct {
	// List of capabilities assigned to role.
	Capabilities pulumi.StringArrayInput
	// Maximum number of concurrently running real-time searches that all members of this role can have.
	CumulativeRealtimeSearchJobsQuota pulumi.IntPtrInput
	// Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
	CumulativeSearchJobsQuota pulumi.IntPtrInput
	// Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
	DefaultApp pulumi.StringPtrInput
	// List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
	ImportedRoles pulumi.StringArrayInput
	// The name of the user role to create.
	Name pulumi.StringPtrInput
	// Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
	RealtimeSearchJobsQuota pulumi.IntPtrInput
	// Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
	SearchDiskQuota pulumi.IntPtrInput
	// Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
	SearchFilter pulumi.StringPtrInput
	// List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
	SearchIndexesAlloweds pulumi.StringArrayInput
	// List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
	SearchIndexesDefaults pulumi.StringArrayInput
	// The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
	SearchJobsQuota pulumi.IntPtrInput
	// Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
	SearchTimeWin pulumi.IntPtrInput
}

func (AuthorizationRolesState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationRolesState)(nil)).Elem()
}

type authorizationRolesArgs struct {
	// List of capabilities assigned to role.
	Capabilities []string `pulumi:"capabilities"`
	// Maximum number of concurrently running real-time searches that all members of this role can have.
	CumulativeRealtimeSearchJobsQuota *int `pulumi:"cumulativeRealtimeSearchJobsQuota"`
	// Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
	CumulativeSearchJobsQuota *int `pulumi:"cumulativeSearchJobsQuota"`
	// Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
	DefaultApp *string `pulumi:"defaultApp"`
	// List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
	ImportedRoles []string `pulumi:"importedRoles"`
	// The name of the user role to create.
	Name *string `pulumi:"name"`
	// Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
	RealtimeSearchJobsQuota *int `pulumi:"realtimeSearchJobsQuota"`
	// Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
	SearchDiskQuota *int `pulumi:"searchDiskQuota"`
	// Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
	SearchFilter *string `pulumi:"searchFilter"`
	// List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
	SearchIndexesAlloweds []string `pulumi:"searchIndexesAlloweds"`
	// List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
	SearchIndexesDefaults []string `pulumi:"searchIndexesDefaults"`
	// The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
	SearchJobsQuota *int `pulumi:"searchJobsQuota"`
	// Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
	SearchTimeWin *int `pulumi:"searchTimeWin"`
}

// The set of arguments for constructing a AuthorizationRoles resource.
type AuthorizationRolesArgs struct {
	// List of capabilities assigned to role.
	Capabilities pulumi.StringArrayInput
	// Maximum number of concurrently running real-time searches that all members of this role can have.
	CumulativeRealtimeSearchJobsQuota pulumi.IntPtrInput
	// Maximum number of concurrently running searches for all role members. Warning message logged when limit is reached.
	CumulativeSearchJobsQuota pulumi.IntPtrInput
	// Specify the folder name of the default app to use for this role. A user-specific default app overrides this.
	DefaultApp pulumi.StringPtrInput
	// List of imported roles for this role. <br>Importing other roles imports all aspects of that role, such as capabilities and allowed indexes to search. In combining multiple roles, the effective value for each attribute is value with the broadest permissions.
	ImportedRoles pulumi.StringArrayInput
	// The name of the user role to create.
	Name pulumi.StringPtrInput
	// Specify the maximum number of concurrent real-time search jobs for this role. This count is independent from the normal search jobs limit.
	RealtimeSearchJobsQuota pulumi.IntPtrInput
	// Specifies the maximum disk space in MB that can be used by a user's search jobs. For example, a value of 100 limits this role to 100 MB total.
	SearchDiskQuota pulumi.IntPtrInput
	// Specify a search string that restricts the scope of searches run by this role. Search results for this role only show events that also match the search string you specify. In the case that a user has multiple roles with different search filters, they are combined with an OR.
	SearchFilter pulumi.StringPtrInput
	// List of indexes that this role has permissions to search. These may be wildcarded, but the index name must begin with an underscore to match internal indexes.
	SearchIndexesAlloweds pulumi.StringArrayInput
	// List of indexes to search when no index is specified. These indexes can be wildcarded, with the exception that '*' does not match internal indexes. To match internal indexes, start with '_'. All internal indexes are represented by '_*'. A user with this role can search other indexes using "index= "
	SearchIndexesDefaults pulumi.StringArrayInput
	// The maximum number of concurrent searches a user with this role is allowed to run. For users with multiple roles, the maximum quota value among all of the roles applies.
	SearchJobsQuota pulumi.IntPtrInput
	// Maximum time span of a search, in seconds. By default, searches are not limited to any specific time window. To override any search time windows from imported roles, set srchTimeWin to '0', as the 'admin' role does.
	SearchTimeWin pulumi.IntPtrInput
}

func (AuthorizationRolesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationRolesArgs)(nil)).Elem()
}

type AuthorizationRolesInput interface {
	pulumi.Input

	ToAuthorizationRolesOutput() AuthorizationRolesOutput
	ToAuthorizationRolesOutputWithContext(ctx context.Context) AuthorizationRolesOutput
}

func (*AuthorizationRoles) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationRoles)(nil))
}

func (i *AuthorizationRoles) ToAuthorizationRolesOutput() AuthorizationRolesOutput {
	return i.ToAuthorizationRolesOutputWithContext(context.Background())
}

func (i *AuthorizationRoles) ToAuthorizationRolesOutputWithContext(ctx context.Context) AuthorizationRolesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRolesOutput)
}

func (i *AuthorizationRoles) ToAuthorizationRolesPtrOutput() AuthorizationRolesPtrOutput {
	return i.ToAuthorizationRolesPtrOutputWithContext(context.Background())
}

func (i *AuthorizationRoles) ToAuthorizationRolesPtrOutputWithContext(ctx context.Context) AuthorizationRolesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRolesPtrOutput)
}

type AuthorizationRolesPtrInput interface {
	pulumi.Input

	ToAuthorizationRolesPtrOutput() AuthorizationRolesPtrOutput
	ToAuthorizationRolesPtrOutputWithContext(ctx context.Context) AuthorizationRolesPtrOutput
}

type authorizationRolesPtrType AuthorizationRolesArgs

func (*authorizationRolesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationRoles)(nil))
}

func (i *authorizationRolesPtrType) ToAuthorizationRolesPtrOutput() AuthorizationRolesPtrOutput {
	return i.ToAuthorizationRolesPtrOutputWithContext(context.Background())
}

func (i *authorizationRolesPtrType) ToAuthorizationRolesPtrOutputWithContext(ctx context.Context) AuthorizationRolesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRolesPtrOutput)
}

// AuthorizationRolesArrayInput is an input type that accepts AuthorizationRolesArray and AuthorizationRolesArrayOutput values.
// You can construct a concrete instance of `AuthorizationRolesArrayInput` via:
//
//          AuthorizationRolesArray{ AuthorizationRolesArgs{...} }
type AuthorizationRolesArrayInput interface {
	pulumi.Input

	ToAuthorizationRolesArrayOutput() AuthorizationRolesArrayOutput
	ToAuthorizationRolesArrayOutputWithContext(context.Context) AuthorizationRolesArrayOutput
}

type AuthorizationRolesArray []AuthorizationRolesInput

func (AuthorizationRolesArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AuthorizationRoles)(nil))
}

func (i AuthorizationRolesArray) ToAuthorizationRolesArrayOutput() AuthorizationRolesArrayOutput {
	return i.ToAuthorizationRolesArrayOutputWithContext(context.Background())
}

func (i AuthorizationRolesArray) ToAuthorizationRolesArrayOutputWithContext(ctx context.Context) AuthorizationRolesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRolesArrayOutput)
}

// AuthorizationRolesMapInput is an input type that accepts AuthorizationRolesMap and AuthorizationRolesMapOutput values.
// You can construct a concrete instance of `AuthorizationRolesMapInput` via:
//
//          AuthorizationRolesMap{ "key": AuthorizationRolesArgs{...} }
type AuthorizationRolesMapInput interface {
	pulumi.Input

	ToAuthorizationRolesMapOutput() AuthorizationRolesMapOutput
	ToAuthorizationRolesMapOutputWithContext(context.Context) AuthorizationRolesMapOutput
}

type AuthorizationRolesMap map[string]AuthorizationRolesInput

func (AuthorizationRolesMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AuthorizationRoles)(nil))
}

func (i AuthorizationRolesMap) ToAuthorizationRolesMapOutput() AuthorizationRolesMapOutput {
	return i.ToAuthorizationRolesMapOutputWithContext(context.Background())
}

func (i AuthorizationRolesMap) ToAuthorizationRolesMapOutputWithContext(ctx context.Context) AuthorizationRolesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRolesMapOutput)
}

type AuthorizationRolesOutput struct {
	*pulumi.OutputState
}

func (AuthorizationRolesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationRoles)(nil))
}

func (o AuthorizationRolesOutput) ToAuthorizationRolesOutput() AuthorizationRolesOutput {
	return o
}

func (o AuthorizationRolesOutput) ToAuthorizationRolesOutputWithContext(ctx context.Context) AuthorizationRolesOutput {
	return o
}

func (o AuthorizationRolesOutput) ToAuthorizationRolesPtrOutput() AuthorizationRolesPtrOutput {
	return o.ToAuthorizationRolesPtrOutputWithContext(context.Background())
}

func (o AuthorizationRolesOutput) ToAuthorizationRolesPtrOutputWithContext(ctx context.Context) AuthorizationRolesPtrOutput {
	return o.ApplyT(func(v AuthorizationRoles) *AuthorizationRoles {
		return &v
	}).(AuthorizationRolesPtrOutput)
}

type AuthorizationRolesPtrOutput struct {
	*pulumi.OutputState
}

func (AuthorizationRolesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationRoles)(nil))
}

func (o AuthorizationRolesPtrOutput) ToAuthorizationRolesPtrOutput() AuthorizationRolesPtrOutput {
	return o
}

func (o AuthorizationRolesPtrOutput) ToAuthorizationRolesPtrOutputWithContext(ctx context.Context) AuthorizationRolesPtrOutput {
	return o
}

type AuthorizationRolesArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationRolesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizationRoles)(nil))
}

func (o AuthorizationRolesArrayOutput) ToAuthorizationRolesArrayOutput() AuthorizationRolesArrayOutput {
	return o
}

func (o AuthorizationRolesArrayOutput) ToAuthorizationRolesArrayOutputWithContext(ctx context.Context) AuthorizationRolesArrayOutput {
	return o
}

func (o AuthorizationRolesArrayOutput) Index(i pulumi.IntInput) AuthorizationRolesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthorizationRoles {
		return vs[0].([]AuthorizationRoles)[vs[1].(int)]
	}).(AuthorizationRolesOutput)
}

type AuthorizationRolesMapOutput struct{ *pulumi.OutputState }

func (AuthorizationRolesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthorizationRoles)(nil))
}

func (o AuthorizationRolesMapOutput) ToAuthorizationRolesMapOutput() AuthorizationRolesMapOutput {
	return o
}

func (o AuthorizationRolesMapOutput) ToAuthorizationRolesMapOutputWithContext(ctx context.Context) AuthorizationRolesMapOutput {
	return o
}

func (o AuthorizationRolesMapOutput) MapIndex(k pulumi.StringInput) AuthorizationRolesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthorizationRoles {
		return vs[0].(map[string]AuthorizationRoles)[vs[1].(string)]
	}).(AuthorizationRolesOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationRolesOutput{})
	pulumi.RegisterOutputType(AuthorizationRolesPtrOutput{})
	pulumi.RegisterOutputType(AuthorizationRolesArrayOutput{})
	pulumi.RegisterOutputType(AuthorizationRolesMapOutput{})
}
