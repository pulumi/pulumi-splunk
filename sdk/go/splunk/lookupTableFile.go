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

// ## # Resource: LookupTableFile
//
// Create and manage lookup table files.
//
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
//			_, err := splunk.NewLookupTableFile(ctx, "lookup_table_file", &splunk.LookupTableFileArgs{
//				App:      pulumi.String("search"),
//				Owner:    pulumi.String("nobody"),
//				FileName: pulumi.String("lookup.csv"),
//				FileContents: pulumi.StringArrayArray{
//					pulumi.StringArray{
//						pulumi.String("status"),
//						pulumi.String("status_description"),
//						pulumi.String("status_type"),
//					},
//					pulumi.StringArray{
//						pulumi.String("100"),
//						pulumi.String("Continue"),
//						pulumi.String("Informational"),
//					},
//					pulumi.StringArray{
//						pulumi.String("101"),
//						pulumi.String("Switching Protocols"),
//						pulumi.String("Informational"),
//					},
//					pulumi.StringArray{
//						pulumi.String("200"),
//						pulumi.String("OK"),
//						pulumi.String("Successful"),
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
type LookupTableFile struct {
	pulumi.CustomResourceState

	// The app context for the resource.
	App pulumi.StringOutput `pulumi:"app"`
	// The column header and row value contents for the lookup table file.
	FileContents pulumi.StringArrayArrayOutput `pulumi:"fileContents"`
	// A name for the lookup table file. Generally ends with ".csv"
	FileName pulumi.StringOutput `pulumi:"fileName"`
	// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
	Owner pulumi.StringOutput `pulumi:"owner"`
}

// NewLookupTableFile registers a new resource with the given unique name, arguments, and options.
func NewLookupTableFile(ctx *pulumi.Context,
	name string, args *LookupTableFileArgs, opts ...pulumi.ResourceOption) (*LookupTableFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.App == nil {
		return nil, errors.New("invalid value for required argument 'App'")
	}
	if args.FileContents == nil {
		return nil, errors.New("invalid value for required argument 'FileContents'")
	}
	if args.FileName == nil {
		return nil, errors.New("invalid value for required argument 'FileName'")
	}
	if args.Owner == nil {
		return nil, errors.New("invalid value for required argument 'Owner'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LookupTableFile
	err := ctx.RegisterResource("splunk:index/lookupTableFile:LookupTableFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLookupTableFile gets an existing LookupTableFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLookupTableFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LookupTableFileState, opts ...pulumi.ResourceOption) (*LookupTableFile, error) {
	var resource LookupTableFile
	err := ctx.ReadResource("splunk:index/lookupTableFile:LookupTableFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LookupTableFile resources.
type lookupTableFileState struct {
	// The app context for the resource.
	App *string `pulumi:"app"`
	// The column header and row value contents for the lookup table file.
	FileContents [][]string `pulumi:"fileContents"`
	// A name for the lookup table file. Generally ends with ".csv"
	FileName *string `pulumi:"fileName"`
	// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
	Owner *string `pulumi:"owner"`
}

type LookupTableFileState struct {
	// The app context for the resource.
	App pulumi.StringPtrInput
	// The column header and row value contents for the lookup table file.
	FileContents pulumi.StringArrayArrayInput
	// A name for the lookup table file. Generally ends with ".csv"
	FileName pulumi.StringPtrInput
	// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
	Owner pulumi.StringPtrInput
}

func (LookupTableFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*lookupTableFileState)(nil)).Elem()
}

type lookupTableFileArgs struct {
	// The app context for the resource.
	App string `pulumi:"app"`
	// The column header and row value contents for the lookup table file.
	FileContents [][]string `pulumi:"fileContents"`
	// A name for the lookup table file. Generally ends with ".csv"
	FileName string `pulumi:"fileName"`
	// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
	Owner string `pulumi:"owner"`
}

// The set of arguments for constructing a LookupTableFile resource.
type LookupTableFileArgs struct {
	// The app context for the resource.
	App pulumi.StringInput
	// The column header and row value contents for the lookup table file.
	FileContents pulumi.StringArrayArrayInput
	// A name for the lookup table file. Generally ends with ".csv"
	FileName pulumi.StringInput
	// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
	Owner pulumi.StringInput
}

func (LookupTableFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lookupTableFileArgs)(nil)).Elem()
}

type LookupTableFileInput interface {
	pulumi.Input

	ToLookupTableFileOutput() LookupTableFileOutput
	ToLookupTableFileOutputWithContext(ctx context.Context) LookupTableFileOutput
}

func (*LookupTableFile) ElementType() reflect.Type {
	return reflect.TypeOf((**LookupTableFile)(nil)).Elem()
}

func (i *LookupTableFile) ToLookupTableFileOutput() LookupTableFileOutput {
	return i.ToLookupTableFileOutputWithContext(context.Background())
}

func (i *LookupTableFile) ToLookupTableFileOutputWithContext(ctx context.Context) LookupTableFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupTableFileOutput)
}

// LookupTableFileArrayInput is an input type that accepts LookupTableFileArray and LookupTableFileArrayOutput values.
// You can construct a concrete instance of `LookupTableFileArrayInput` via:
//
//	LookupTableFileArray{ LookupTableFileArgs{...} }
type LookupTableFileArrayInput interface {
	pulumi.Input

	ToLookupTableFileArrayOutput() LookupTableFileArrayOutput
	ToLookupTableFileArrayOutputWithContext(context.Context) LookupTableFileArrayOutput
}

type LookupTableFileArray []LookupTableFileInput

func (LookupTableFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LookupTableFile)(nil)).Elem()
}

func (i LookupTableFileArray) ToLookupTableFileArrayOutput() LookupTableFileArrayOutput {
	return i.ToLookupTableFileArrayOutputWithContext(context.Background())
}

func (i LookupTableFileArray) ToLookupTableFileArrayOutputWithContext(ctx context.Context) LookupTableFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupTableFileArrayOutput)
}

// LookupTableFileMapInput is an input type that accepts LookupTableFileMap and LookupTableFileMapOutput values.
// You can construct a concrete instance of `LookupTableFileMapInput` via:
//
//	LookupTableFileMap{ "key": LookupTableFileArgs{...} }
type LookupTableFileMapInput interface {
	pulumi.Input

	ToLookupTableFileMapOutput() LookupTableFileMapOutput
	ToLookupTableFileMapOutputWithContext(context.Context) LookupTableFileMapOutput
}

type LookupTableFileMap map[string]LookupTableFileInput

func (LookupTableFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LookupTableFile)(nil)).Elem()
}

func (i LookupTableFileMap) ToLookupTableFileMapOutput() LookupTableFileMapOutput {
	return i.ToLookupTableFileMapOutputWithContext(context.Background())
}

func (i LookupTableFileMap) ToLookupTableFileMapOutputWithContext(ctx context.Context) LookupTableFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupTableFileMapOutput)
}

type LookupTableFileOutput struct{ *pulumi.OutputState }

func (LookupTableFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LookupTableFile)(nil)).Elem()
}

func (o LookupTableFileOutput) ToLookupTableFileOutput() LookupTableFileOutput {
	return o
}

func (o LookupTableFileOutput) ToLookupTableFileOutputWithContext(ctx context.Context) LookupTableFileOutput {
	return o
}

// The app context for the resource.
func (o LookupTableFileOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v *LookupTableFile) pulumi.StringOutput { return v.App }).(pulumi.StringOutput)
}

// The column header and row value contents for the lookup table file.
func (o LookupTableFileOutput) FileContents() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v *LookupTableFile) pulumi.StringArrayArrayOutput { return v.FileContents }).(pulumi.StringArrayArrayOutput)
}

// A name for the lookup table file. Generally ends with ".csv"
func (o LookupTableFileOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v *LookupTableFile) pulumi.StringOutput { return v.FileName }).(pulumi.StringOutput)
}

// User name of resource owner. Defaults to the resource creator. Required for updating any knowledge object ACL properties. nobody = All users may access the resource, but write access to the resource might be restricted.
func (o LookupTableFileOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *LookupTableFile) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

type LookupTableFileArrayOutput struct{ *pulumi.OutputState }

func (LookupTableFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LookupTableFile)(nil)).Elem()
}

func (o LookupTableFileArrayOutput) ToLookupTableFileArrayOutput() LookupTableFileArrayOutput {
	return o
}

func (o LookupTableFileArrayOutput) ToLookupTableFileArrayOutputWithContext(ctx context.Context) LookupTableFileArrayOutput {
	return o
}

func (o LookupTableFileArrayOutput) Index(i pulumi.IntInput) LookupTableFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LookupTableFile {
		return vs[0].([]*LookupTableFile)[vs[1].(int)]
	}).(LookupTableFileOutput)
}

type LookupTableFileMapOutput struct{ *pulumi.OutputState }

func (LookupTableFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LookupTableFile)(nil)).Elem()
}

func (o LookupTableFileMapOutput) ToLookupTableFileMapOutput() LookupTableFileMapOutput {
	return o
}

func (o LookupTableFileMapOutput) ToLookupTableFileMapOutputWithContext(ctx context.Context) LookupTableFileMapOutput {
	return o
}

func (o LookupTableFileMapOutput) MapIndex(k pulumi.StringInput) LookupTableFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LookupTableFile {
		return vs[0].(map[string]*LookupTableFile)[vs[1].(string)]
	}).(LookupTableFileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LookupTableFileInput)(nil)).Elem(), &LookupTableFile{})
	pulumi.RegisterInputType(reflect.TypeOf((*LookupTableFileArrayInput)(nil)).Elem(), LookupTableFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LookupTableFileMapInput)(nil)).Elem(), LookupTableFileMap{})
	pulumi.RegisterOutputType(LookupTableFileOutput{})
	pulumi.RegisterOutputType(LookupTableFileArrayOutput{})
	pulumi.RegisterOutputType(LookupTableFileMapOutput{})
}
