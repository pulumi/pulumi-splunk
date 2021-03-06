// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: AuthenticationUsers
//
// Create and update user information or delete the user.
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
// 		_, err := splunk.NewAuthenticationUsers(ctx, "user01", &splunk.AuthenticationUsersArgs{
// 			Email:           pulumi.String("user01@example.com"),
// 			ForceChangePass: pulumi.Bool(false),
// 			Password:        pulumi.String("password01"),
// 			Roles: pulumi.StringArray{
// 				pulumi.String("terraform-user01-role"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AuthenticationUsers struct {
	pulumi.CustomResourceState

	// User default app. Overrides the default app inherited from the user roles.
	DefaultApp pulumi.StringOutput `pulumi:"defaultApp"`
	// User email address.
	Email pulumi.StringOutput `pulumi:"email"`
	// Force user to change password indication
	ForceChangePass pulumi.BoolPtrOutput `pulumi:"forceChangePass"`
	// Unique user login name.
	Name pulumi.StringOutput `pulumi:"name"`
	// User login password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Full user name.
	Realname pulumi.StringOutput `pulumi:"realname"`
	// Restart background search job that has not completed when Splunk restarts indication.
	RestartBackgroundJobs pulumi.BoolOutput `pulumi:"restartBackgroundJobs"`
	// Role to assign to this user. At least one existing role is required.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// User timezone.
	Tz pulumi.StringOutput `pulumi:"tz"`
}

// NewAuthenticationUsers registers a new resource with the given unique name, arguments, and options.
func NewAuthenticationUsers(ctx *pulumi.Context,
	name string, args *AuthenticationUsersArgs, opts ...pulumi.ResourceOption) (*AuthenticationUsers, error) {
	if args == nil {
		args = &AuthenticationUsersArgs{}
	}

	var resource AuthenticationUsers
	err := ctx.RegisterResource("splunk:index/authenticationUsers:AuthenticationUsers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthenticationUsers gets an existing AuthenticationUsers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthenticationUsers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthenticationUsersState, opts ...pulumi.ResourceOption) (*AuthenticationUsers, error) {
	var resource AuthenticationUsers
	err := ctx.ReadResource("splunk:index/authenticationUsers:AuthenticationUsers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthenticationUsers resources.
type authenticationUsersState struct {
	// User default app. Overrides the default app inherited from the user roles.
	DefaultApp *string `pulumi:"defaultApp"`
	// User email address.
	Email *string `pulumi:"email"`
	// Force user to change password indication
	ForceChangePass *bool `pulumi:"forceChangePass"`
	// Unique user login name.
	Name *string `pulumi:"name"`
	// User login password.
	Password *string `pulumi:"password"`
	// Full user name.
	Realname *string `pulumi:"realname"`
	// Restart background search job that has not completed when Splunk restarts indication.
	RestartBackgroundJobs *bool `pulumi:"restartBackgroundJobs"`
	// Role to assign to this user. At least one existing role is required.
	Roles []string `pulumi:"roles"`
	// User timezone.
	Tz *string `pulumi:"tz"`
}

type AuthenticationUsersState struct {
	// User default app. Overrides the default app inherited from the user roles.
	DefaultApp pulumi.StringPtrInput
	// User email address.
	Email pulumi.StringPtrInput
	// Force user to change password indication
	ForceChangePass pulumi.BoolPtrInput
	// Unique user login name.
	Name pulumi.StringPtrInput
	// User login password.
	Password pulumi.StringPtrInput
	// Full user name.
	Realname pulumi.StringPtrInput
	// Restart background search job that has not completed when Splunk restarts indication.
	RestartBackgroundJobs pulumi.BoolPtrInput
	// Role to assign to this user. At least one existing role is required.
	Roles pulumi.StringArrayInput
	// User timezone.
	Tz pulumi.StringPtrInput
}

func (AuthenticationUsersState) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationUsersState)(nil)).Elem()
}

type authenticationUsersArgs struct {
	// User default app. Overrides the default app inherited from the user roles.
	DefaultApp *string `pulumi:"defaultApp"`
	// User email address.
	Email *string `pulumi:"email"`
	// Force user to change password indication
	ForceChangePass *bool `pulumi:"forceChangePass"`
	// Unique user login name.
	Name *string `pulumi:"name"`
	// User login password.
	Password *string `pulumi:"password"`
	// Full user name.
	Realname *string `pulumi:"realname"`
	// Restart background search job that has not completed when Splunk restarts indication.
	RestartBackgroundJobs *bool `pulumi:"restartBackgroundJobs"`
	// Role to assign to this user. At least one existing role is required.
	Roles []string `pulumi:"roles"`
	// User timezone.
	Tz *string `pulumi:"tz"`
}

// The set of arguments for constructing a AuthenticationUsers resource.
type AuthenticationUsersArgs struct {
	// User default app. Overrides the default app inherited from the user roles.
	DefaultApp pulumi.StringPtrInput
	// User email address.
	Email pulumi.StringPtrInput
	// Force user to change password indication
	ForceChangePass pulumi.BoolPtrInput
	// Unique user login name.
	Name pulumi.StringPtrInput
	// User login password.
	Password pulumi.StringPtrInput
	// Full user name.
	Realname pulumi.StringPtrInput
	// Restart background search job that has not completed when Splunk restarts indication.
	RestartBackgroundJobs pulumi.BoolPtrInput
	// Role to assign to this user. At least one existing role is required.
	Roles pulumi.StringArrayInput
	// User timezone.
	Tz pulumi.StringPtrInput
}

func (AuthenticationUsersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationUsersArgs)(nil)).Elem()
}

type AuthenticationUsersInput interface {
	pulumi.Input

	ToAuthenticationUsersOutput() AuthenticationUsersOutput
	ToAuthenticationUsersOutputWithContext(ctx context.Context) AuthenticationUsersOutput
}

func (*AuthenticationUsers) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationUsers)(nil))
}

func (i *AuthenticationUsers) ToAuthenticationUsersOutput() AuthenticationUsersOutput {
	return i.ToAuthenticationUsersOutputWithContext(context.Background())
}

func (i *AuthenticationUsers) ToAuthenticationUsersOutputWithContext(ctx context.Context) AuthenticationUsersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationUsersOutput)
}

func (i *AuthenticationUsers) ToAuthenticationUsersPtrOutput() AuthenticationUsersPtrOutput {
	return i.ToAuthenticationUsersPtrOutputWithContext(context.Background())
}

func (i *AuthenticationUsers) ToAuthenticationUsersPtrOutputWithContext(ctx context.Context) AuthenticationUsersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationUsersPtrOutput)
}

type AuthenticationUsersPtrInput interface {
	pulumi.Input

	ToAuthenticationUsersPtrOutput() AuthenticationUsersPtrOutput
	ToAuthenticationUsersPtrOutputWithContext(ctx context.Context) AuthenticationUsersPtrOutput
}

type authenticationUsersPtrType AuthenticationUsersArgs

func (*authenticationUsersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationUsers)(nil))
}

func (i *authenticationUsersPtrType) ToAuthenticationUsersPtrOutput() AuthenticationUsersPtrOutput {
	return i.ToAuthenticationUsersPtrOutputWithContext(context.Background())
}

func (i *authenticationUsersPtrType) ToAuthenticationUsersPtrOutputWithContext(ctx context.Context) AuthenticationUsersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationUsersPtrOutput)
}

// AuthenticationUsersArrayInput is an input type that accepts AuthenticationUsersArray and AuthenticationUsersArrayOutput values.
// You can construct a concrete instance of `AuthenticationUsersArrayInput` via:
//
//          AuthenticationUsersArray{ AuthenticationUsersArgs{...} }
type AuthenticationUsersArrayInput interface {
	pulumi.Input

	ToAuthenticationUsersArrayOutput() AuthenticationUsersArrayOutput
	ToAuthenticationUsersArrayOutputWithContext(context.Context) AuthenticationUsersArrayOutput
}

type AuthenticationUsersArray []AuthenticationUsersInput

func (AuthenticationUsersArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AuthenticationUsers)(nil))
}

func (i AuthenticationUsersArray) ToAuthenticationUsersArrayOutput() AuthenticationUsersArrayOutput {
	return i.ToAuthenticationUsersArrayOutputWithContext(context.Background())
}

func (i AuthenticationUsersArray) ToAuthenticationUsersArrayOutputWithContext(ctx context.Context) AuthenticationUsersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationUsersArrayOutput)
}

// AuthenticationUsersMapInput is an input type that accepts AuthenticationUsersMap and AuthenticationUsersMapOutput values.
// You can construct a concrete instance of `AuthenticationUsersMapInput` via:
//
//          AuthenticationUsersMap{ "key": AuthenticationUsersArgs{...} }
type AuthenticationUsersMapInput interface {
	pulumi.Input

	ToAuthenticationUsersMapOutput() AuthenticationUsersMapOutput
	ToAuthenticationUsersMapOutputWithContext(context.Context) AuthenticationUsersMapOutput
}

type AuthenticationUsersMap map[string]AuthenticationUsersInput

func (AuthenticationUsersMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AuthenticationUsers)(nil))
}

func (i AuthenticationUsersMap) ToAuthenticationUsersMapOutput() AuthenticationUsersMapOutput {
	return i.ToAuthenticationUsersMapOutputWithContext(context.Background())
}

func (i AuthenticationUsersMap) ToAuthenticationUsersMapOutputWithContext(ctx context.Context) AuthenticationUsersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationUsersMapOutput)
}

type AuthenticationUsersOutput struct {
	*pulumi.OutputState
}

func (AuthenticationUsersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationUsers)(nil))
}

func (o AuthenticationUsersOutput) ToAuthenticationUsersOutput() AuthenticationUsersOutput {
	return o
}

func (o AuthenticationUsersOutput) ToAuthenticationUsersOutputWithContext(ctx context.Context) AuthenticationUsersOutput {
	return o
}

func (o AuthenticationUsersOutput) ToAuthenticationUsersPtrOutput() AuthenticationUsersPtrOutput {
	return o.ToAuthenticationUsersPtrOutputWithContext(context.Background())
}

func (o AuthenticationUsersOutput) ToAuthenticationUsersPtrOutputWithContext(ctx context.Context) AuthenticationUsersPtrOutput {
	return o.ApplyT(func(v AuthenticationUsers) *AuthenticationUsers {
		return &v
	}).(AuthenticationUsersPtrOutput)
}

type AuthenticationUsersPtrOutput struct {
	*pulumi.OutputState
}

func (AuthenticationUsersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationUsers)(nil))
}

func (o AuthenticationUsersPtrOutput) ToAuthenticationUsersPtrOutput() AuthenticationUsersPtrOutput {
	return o
}

func (o AuthenticationUsersPtrOutput) ToAuthenticationUsersPtrOutputWithContext(ctx context.Context) AuthenticationUsersPtrOutput {
	return o
}

type AuthenticationUsersArrayOutput struct{ *pulumi.OutputState }

func (AuthenticationUsersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthenticationUsers)(nil))
}

func (o AuthenticationUsersArrayOutput) ToAuthenticationUsersArrayOutput() AuthenticationUsersArrayOutput {
	return o
}

func (o AuthenticationUsersArrayOutput) ToAuthenticationUsersArrayOutputWithContext(ctx context.Context) AuthenticationUsersArrayOutput {
	return o
}

func (o AuthenticationUsersArrayOutput) Index(i pulumi.IntInput) AuthenticationUsersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthenticationUsers {
		return vs[0].([]AuthenticationUsers)[vs[1].(int)]
	}).(AuthenticationUsersOutput)
}

type AuthenticationUsersMapOutput struct{ *pulumi.OutputState }

func (AuthenticationUsersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthenticationUsers)(nil))
}

func (o AuthenticationUsersMapOutput) ToAuthenticationUsersMapOutput() AuthenticationUsersMapOutput {
	return o
}

func (o AuthenticationUsersMapOutput) ToAuthenticationUsersMapOutputWithContext(ctx context.Context) AuthenticationUsersMapOutput {
	return o
}

func (o AuthenticationUsersMapOutput) MapIndex(k pulumi.StringInput) AuthenticationUsersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthenticationUsers {
		return vs[0].(map[string]AuthenticationUsers)[vs[1].(string)]
	}).(AuthenticationUsersOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthenticationUsersOutput{})
	pulumi.RegisterOutputType(AuthenticationUsersPtrOutput{})
	pulumi.RegisterOutputType(AuthenticationUsersArrayOutput{})
	pulumi.RegisterOutputType(AuthenticationUsersMapOutput{})
}
