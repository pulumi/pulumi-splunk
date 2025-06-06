// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-splunk/sdk/go/splunk/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-splunk/sdk/go/splunk"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := splunk.NewGenericAcl(ctx, "my_app", &splunk.GenericAclArgs{
//				Path: pulumi.String("apps/local/my_app"),
//				Acl: &splunk.GenericAclAclArgs{
//					App:   pulumi.String("system"),
//					Owner: pulumi.String("nobody"),
//					Reads: pulumi.StringArray{
//						pulumi.String("*"),
//					},
//					Writes: pulumi.StringArray{
//						pulumi.String("admin"),
//						pulumi.String("power"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = splunk.NewGenericAcl(ctx, "my_dashboard", &splunk.GenericAclArgs{
//				Path: pulumi.String("data/ui/views/my_dashboard"),
//				Acl: &splunk.GenericAclAclArgs{
//					App:   pulumi.String("my_app"),
//					Owner: pulumi.String("joe_user"),
//					Reads: pulumi.StringArray{
//						pulumi.String("team_joe"),
//					},
//					Writes: pulumi.StringArray{
//						pulumi.String("team_joe"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Generic ACL resources can be imported by specifying their owner, app, and path with a colon-delimited string as the ID:
//
// ```sh
// $ pulumi import splunk:index/genericAcl:GenericAcl splunk_generic_acl <owner>:<app>:<path>
// ```
type GenericAcl struct {
	pulumi.CustomResourceState

	// The ACL to apply to the object, including app/owner to properly identify the object.
	// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
	// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
	// app and owner for objects that don't fit in the normal namespace.
	Acl GenericAclAclOutput `pulumi:"acl"`
	// REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
	Path pulumi.StringOutput `pulumi:"path"`
}

// NewGenericAcl registers a new resource with the given unique name, arguments, and options.
func NewGenericAcl(ctx *pulumi.Context,
	name string, args *GenericAclArgs, opts ...pulumi.ResourceOption) (*GenericAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GenericAcl
	err := ctx.RegisterResource("splunk:index/genericAcl:GenericAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGenericAcl gets an existing GenericAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGenericAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GenericAclState, opts ...pulumi.ResourceOption) (*GenericAcl, error) {
	var resource GenericAcl
	err := ctx.ReadResource("splunk:index/genericAcl:GenericAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GenericAcl resources.
type genericAclState struct {
	// The ACL to apply to the object, including app/owner to properly identify the object.
	// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
	// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
	// app and owner for objects that don't fit in the normal namespace.
	Acl *GenericAclAcl `pulumi:"acl"`
	// REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
	Path *string `pulumi:"path"`
}

type GenericAclState struct {
	// The ACL to apply to the object, including app/owner to properly identify the object.
	// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
	// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
	// app and owner for objects that don't fit in the normal namespace.
	Acl GenericAclAclPtrInput
	// REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
	Path pulumi.StringPtrInput
}

func (GenericAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*genericAclState)(nil)).Elem()
}

type genericAclArgs struct {
	// The ACL to apply to the object, including app/owner to properly identify the object.
	// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
	// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
	// app and owner for objects that don't fit in the normal namespace.
	Acl *GenericAclAcl `pulumi:"acl"`
	// REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
	Path string `pulumi:"path"`
}

// The set of arguments for constructing a GenericAcl resource.
type GenericAclArgs struct {
	// The ACL to apply to the object, including app/owner to properly identify the object.
	// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
	// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
	// app and owner for objects that don't fit in the normal namespace.
	Acl GenericAclAclPtrInput
	// REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
	Path pulumi.StringInput
}

func (GenericAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*genericAclArgs)(nil)).Elem()
}

type GenericAclInput interface {
	pulumi.Input

	ToGenericAclOutput() GenericAclOutput
	ToGenericAclOutputWithContext(ctx context.Context) GenericAclOutput
}

func (*GenericAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericAcl)(nil)).Elem()
}

func (i *GenericAcl) ToGenericAclOutput() GenericAclOutput {
	return i.ToGenericAclOutputWithContext(context.Background())
}

func (i *GenericAcl) ToGenericAclOutputWithContext(ctx context.Context) GenericAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericAclOutput)
}

// GenericAclArrayInput is an input type that accepts GenericAclArray and GenericAclArrayOutput values.
// You can construct a concrete instance of `GenericAclArrayInput` via:
//
//	GenericAclArray{ GenericAclArgs{...} }
type GenericAclArrayInput interface {
	pulumi.Input

	ToGenericAclArrayOutput() GenericAclArrayOutput
	ToGenericAclArrayOutputWithContext(context.Context) GenericAclArrayOutput
}

type GenericAclArray []GenericAclInput

func (GenericAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GenericAcl)(nil)).Elem()
}

func (i GenericAclArray) ToGenericAclArrayOutput() GenericAclArrayOutput {
	return i.ToGenericAclArrayOutputWithContext(context.Background())
}

func (i GenericAclArray) ToGenericAclArrayOutputWithContext(ctx context.Context) GenericAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericAclArrayOutput)
}

// GenericAclMapInput is an input type that accepts GenericAclMap and GenericAclMapOutput values.
// You can construct a concrete instance of `GenericAclMapInput` via:
//
//	GenericAclMap{ "key": GenericAclArgs{...} }
type GenericAclMapInput interface {
	pulumi.Input

	ToGenericAclMapOutput() GenericAclMapOutput
	ToGenericAclMapOutputWithContext(context.Context) GenericAclMapOutput
}

type GenericAclMap map[string]GenericAclInput

func (GenericAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GenericAcl)(nil)).Elem()
}

func (i GenericAclMap) ToGenericAclMapOutput() GenericAclMapOutput {
	return i.ToGenericAclMapOutputWithContext(context.Background())
}

func (i GenericAclMap) ToGenericAclMapOutputWithContext(ctx context.Context) GenericAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericAclMapOutput)
}

type GenericAclOutput struct{ *pulumi.OutputState }

func (GenericAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericAcl)(nil)).Elem()
}

func (o GenericAclOutput) ToGenericAclOutput() GenericAclOutput {
	return o
}

func (o GenericAclOutput) ToGenericAclOutputWithContext(ctx context.Context) GenericAclOutput {
	return o
}

// The ACL to apply to the object, including app/owner to properly identify the object.
// Though technically optional, it should be explicitly set for this resource to really be valid. Some objects, such as
// apps, require specific values for app and owner. Consult the REST API documentation regarding which values to use for
// app and owner for objects that don't fit in the normal namespace.
func (o GenericAclOutput) Acl() GenericAclAclOutput {
	return o.ApplyT(func(v *GenericAcl) GenericAclAclOutput { return v.Acl }).(GenericAclAclOutput)
}

// REST API Endpoint path to the object, relative to servicesNS/<owner>/<app>
func (o GenericAclOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericAcl) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

type GenericAclArrayOutput struct{ *pulumi.OutputState }

func (GenericAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GenericAcl)(nil)).Elem()
}

func (o GenericAclArrayOutput) ToGenericAclArrayOutput() GenericAclArrayOutput {
	return o
}

func (o GenericAclArrayOutput) ToGenericAclArrayOutputWithContext(ctx context.Context) GenericAclArrayOutput {
	return o
}

func (o GenericAclArrayOutput) Index(i pulumi.IntInput) GenericAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GenericAcl {
		return vs[0].([]*GenericAcl)[vs[1].(int)]
	}).(GenericAclOutput)
}

type GenericAclMapOutput struct{ *pulumi.OutputState }

func (GenericAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GenericAcl)(nil)).Elem()
}

func (o GenericAclMapOutput) ToGenericAclMapOutput() GenericAclMapOutput {
	return o
}

func (o GenericAclMapOutput) ToGenericAclMapOutputWithContext(ctx context.Context) GenericAclMapOutput {
	return o
}

func (o GenericAclMapOutput) MapIndex(k pulumi.StringInput) GenericAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GenericAcl {
		return vs[0].(map[string]*GenericAcl)[vs[1].(string)]
	}).(GenericAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GenericAclInput)(nil)).Elem(), &GenericAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenericAclArrayInput)(nil)).Elem(), GenericAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenericAclMapInput)(nil)).Elem(), GenericAclMap{})
	pulumi.RegisterOutputType(GenericAclOutput{})
	pulumi.RegisterOutputType(GenericAclArrayOutput{})
	pulumi.RegisterOutputType(GenericAclMapOutput{})
}
