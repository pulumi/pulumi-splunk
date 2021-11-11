// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package splunk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: InputsScript
//
// Create or update scripted inputs.
type InputsScript struct {
	pulumi.CustomResourceState

	// The app/user context that is the namespace for the resource
	Acl InputsScriptAclOutput `pulumi:"acl"`
	// Specifies whether the input script is disabled.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Sets the host for events from this input. Defaults to whatever host sent the event.
	Host pulumi.StringOutput `pulumi:"host"`
	// Sets the index for events from this input. Defaults to the main index.
	Index pulumi.StringOutput `pulumi:"index"`
	// Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
	Interval pulumi.IntOutput `pulumi:"interval"`
	// Specify the name of the scripted input.
	Name pulumi.StringOutput `pulumi:"name"`
	// User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
	Passauth pulumi.StringOutput `pulumi:"passauth"`
	// Specify a new name for the source field for the script.
	RenameSource pulumi.StringOutput `pulumi:"renameSource"`
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source pulumi.StringOutput `pulumi:"source"`
	// Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
	// Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
	Sourcetype pulumi.StringOutput `pulumi:"sourcetype"`
}

// NewInputsScript registers a new resource with the given unique name, arguments, and options.
func NewInputsScript(ctx *pulumi.Context,
	name string, args *InputsScriptArgs, opts ...pulumi.ResourceOption) (*InputsScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interval == nil {
		return nil, errors.New("invalid value for required argument 'Interval'")
	}
	var resource InputsScript
	err := ctx.RegisterResource("splunk:index/inputsScript:InputsScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInputsScript gets an existing InputsScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInputsScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputsScriptState, opts ...pulumi.ResourceOption) (*InputsScript, error) {
	var resource InputsScript
	err := ctx.ReadResource("splunk:index/inputsScript:InputsScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InputsScript resources.
type inputsScriptState struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsScriptAcl `pulumi:"acl"`
	// Specifies whether the input script is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Sets the host for events from this input. Defaults to whatever host sent the event.
	Host *string `pulumi:"host"`
	// Sets the index for events from this input. Defaults to the main index.
	Index *string `pulumi:"index"`
	// Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
	Interval *int `pulumi:"interval"`
	// Specify the name of the scripted input.
	Name *string `pulumi:"name"`
	// User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
	Passauth *string `pulumi:"passauth"`
	// Specify a new name for the source field for the script.
	RenameSource *string `pulumi:"renameSource"`
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source *string `pulumi:"source"`
	// Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
	// Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
	Sourcetype *string `pulumi:"sourcetype"`
}

type InputsScriptState struct {
	// The app/user context that is the namespace for the resource
	Acl InputsScriptAclPtrInput
	// Specifies whether the input script is disabled.
	Disabled pulumi.BoolPtrInput
	// Sets the host for events from this input. Defaults to whatever host sent the event.
	Host pulumi.StringPtrInput
	// Sets the index for events from this input. Defaults to the main index.
	Index pulumi.StringPtrInput
	// Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
	Interval pulumi.IntPtrInput
	// Specify the name of the scripted input.
	Name pulumi.StringPtrInput
	// User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
	Passauth pulumi.StringPtrInput
	// Specify a new name for the source field for the script.
	RenameSource pulumi.StringPtrInput
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source pulumi.StringPtrInput
	// Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
	// Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
	Sourcetype pulumi.StringPtrInput
}

func (InputsScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsScriptState)(nil)).Elem()
}

type inputsScriptArgs struct {
	// The app/user context that is the namespace for the resource
	Acl *InputsScriptAcl `pulumi:"acl"`
	// Specifies whether the input script is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Sets the host for events from this input. Defaults to whatever host sent the event.
	Host *string `pulumi:"host"`
	// Sets the index for events from this input. Defaults to the main index.
	Index *string `pulumi:"index"`
	// Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
	Interval int `pulumi:"interval"`
	// Specify the name of the scripted input.
	Name *string `pulumi:"name"`
	// User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
	Passauth *string `pulumi:"passauth"`
	// Specify a new name for the source field for the script.
	RenameSource *string `pulumi:"renameSource"`
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source *string `pulumi:"source"`
	// Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
	// Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
	Sourcetype *string `pulumi:"sourcetype"`
}

// The set of arguments for constructing a InputsScript resource.
type InputsScriptArgs struct {
	// The app/user context that is the namespace for the resource
	Acl InputsScriptAclPtrInput
	// Specifies whether the input script is disabled.
	Disabled pulumi.BoolPtrInput
	// Sets the host for events from this input. Defaults to whatever host sent the event.
	Host pulumi.StringPtrInput
	// Sets the index for events from this input. Defaults to the main index.
	Index pulumi.StringPtrInput
	// Specify an integer or cron schedule. This parameter specifies how often to execute the specified script, in seconds or a valid cron schedule. If you specify a cron schedule, the script is not executed on start-up.
	Interval pulumi.IntInput
	// Specify the name of the scripted input.
	Name pulumi.StringPtrInput
	// User to run the script as. If you provide a username, Splunk software generates an auth token for that user and passes it to the script.
	Passauth pulumi.StringPtrInput
	// Specify a new name for the source field for the script.
	RenameSource pulumi.StringPtrInput
	// Sets the source key/field for events from this input. Defaults to the input file path.
	// Sets the source key initial value. The key is used during parsing/indexing, in particular to set the source field during indexing. It is also the source field used at search time. As a convenience, the chosen string is prepended with 'source::'.
	Source pulumi.StringPtrInput
	// Sets the sourcetype key/field for events from this input. If unset, Splunk software picks a source type based on various aspects of the data. As a convenience, the chosen string is prepended with 'sourcetype::'. There is no hard-coded default.
	// Sets the sourcetype key initial value. The key is used during parsing/indexing, in particular to set the source type field during indexing. It is also the source type field used at search time.
	Sourcetype pulumi.StringPtrInput
}

func (InputsScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputsScriptArgs)(nil)).Elem()
}

type InputsScriptInput interface {
	pulumi.Input

	ToInputsScriptOutput() InputsScriptOutput
	ToInputsScriptOutputWithContext(ctx context.Context) InputsScriptOutput
}

func (*InputsScript) ElementType() reflect.Type {
	return reflect.TypeOf((*InputsScript)(nil))
}

func (i *InputsScript) ToInputsScriptOutput() InputsScriptOutput {
	return i.ToInputsScriptOutputWithContext(context.Background())
}

func (i *InputsScript) ToInputsScriptOutputWithContext(ctx context.Context) InputsScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsScriptOutput)
}

func (i *InputsScript) ToInputsScriptPtrOutput() InputsScriptPtrOutput {
	return i.ToInputsScriptPtrOutputWithContext(context.Background())
}

func (i *InputsScript) ToInputsScriptPtrOutputWithContext(ctx context.Context) InputsScriptPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsScriptPtrOutput)
}

type InputsScriptPtrInput interface {
	pulumi.Input

	ToInputsScriptPtrOutput() InputsScriptPtrOutput
	ToInputsScriptPtrOutputWithContext(ctx context.Context) InputsScriptPtrOutput
}

type inputsScriptPtrType InputsScriptArgs

func (*inputsScriptPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsScript)(nil))
}

func (i *inputsScriptPtrType) ToInputsScriptPtrOutput() InputsScriptPtrOutput {
	return i.ToInputsScriptPtrOutputWithContext(context.Background())
}

func (i *inputsScriptPtrType) ToInputsScriptPtrOutputWithContext(ctx context.Context) InputsScriptPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsScriptPtrOutput)
}

// InputsScriptArrayInput is an input type that accepts InputsScriptArray and InputsScriptArrayOutput values.
// You can construct a concrete instance of `InputsScriptArrayInput` via:
//
//          InputsScriptArray{ InputsScriptArgs{...} }
type InputsScriptArrayInput interface {
	pulumi.Input

	ToInputsScriptArrayOutput() InputsScriptArrayOutput
	ToInputsScriptArrayOutputWithContext(context.Context) InputsScriptArrayOutput
}

type InputsScriptArray []InputsScriptInput

func (InputsScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InputsScript)(nil)).Elem()
}

func (i InputsScriptArray) ToInputsScriptArrayOutput() InputsScriptArrayOutput {
	return i.ToInputsScriptArrayOutputWithContext(context.Background())
}

func (i InputsScriptArray) ToInputsScriptArrayOutputWithContext(ctx context.Context) InputsScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsScriptArrayOutput)
}

// InputsScriptMapInput is an input type that accepts InputsScriptMap and InputsScriptMapOutput values.
// You can construct a concrete instance of `InputsScriptMapInput` via:
//
//          InputsScriptMap{ "key": InputsScriptArgs{...} }
type InputsScriptMapInput interface {
	pulumi.Input

	ToInputsScriptMapOutput() InputsScriptMapOutput
	ToInputsScriptMapOutputWithContext(context.Context) InputsScriptMapOutput
}

type InputsScriptMap map[string]InputsScriptInput

func (InputsScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InputsScript)(nil)).Elem()
}

func (i InputsScriptMap) ToInputsScriptMapOutput() InputsScriptMapOutput {
	return i.ToInputsScriptMapOutputWithContext(context.Background())
}

func (i InputsScriptMap) ToInputsScriptMapOutputWithContext(ctx context.Context) InputsScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputsScriptMapOutput)
}

type InputsScriptOutput struct{ *pulumi.OutputState }

func (InputsScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputsScript)(nil))
}

func (o InputsScriptOutput) ToInputsScriptOutput() InputsScriptOutput {
	return o
}

func (o InputsScriptOutput) ToInputsScriptOutputWithContext(ctx context.Context) InputsScriptOutput {
	return o
}

func (o InputsScriptOutput) ToInputsScriptPtrOutput() InputsScriptPtrOutput {
	return o.ToInputsScriptPtrOutputWithContext(context.Background())
}

func (o InputsScriptOutput) ToInputsScriptPtrOutputWithContext(ctx context.Context) InputsScriptPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputsScript) *InputsScript {
		return &v
	}).(InputsScriptPtrOutput)
}

type InputsScriptPtrOutput struct{ *pulumi.OutputState }

func (InputsScriptPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputsScript)(nil))
}

func (o InputsScriptPtrOutput) ToInputsScriptPtrOutput() InputsScriptPtrOutput {
	return o
}

func (o InputsScriptPtrOutput) ToInputsScriptPtrOutputWithContext(ctx context.Context) InputsScriptPtrOutput {
	return o
}

func (o InputsScriptPtrOutput) Elem() InputsScriptOutput {
	return o.ApplyT(func(v *InputsScript) InputsScript {
		if v != nil {
			return *v
		}
		var ret InputsScript
		return ret
	}).(InputsScriptOutput)
}

type InputsScriptArrayOutput struct{ *pulumi.OutputState }

func (InputsScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputsScript)(nil))
}

func (o InputsScriptArrayOutput) ToInputsScriptArrayOutput() InputsScriptArrayOutput {
	return o
}

func (o InputsScriptArrayOutput) ToInputsScriptArrayOutputWithContext(ctx context.Context) InputsScriptArrayOutput {
	return o
}

func (o InputsScriptArrayOutput) Index(i pulumi.IntInput) InputsScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputsScript {
		return vs[0].([]InputsScript)[vs[1].(int)]
	}).(InputsScriptOutput)
}

type InputsScriptMapOutput struct{ *pulumi.OutputState }

func (InputsScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InputsScript)(nil))
}

func (o InputsScriptMapOutput) ToInputsScriptMapOutput() InputsScriptMapOutput {
	return o
}

func (o InputsScriptMapOutput) ToInputsScriptMapOutputWithContext(ctx context.Context) InputsScriptMapOutput {
	return o
}

func (o InputsScriptMapOutput) MapIndex(k pulumi.StringInput) InputsScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InputsScript {
		return vs[0].(map[string]InputsScript)[vs[1].(string)]
	}).(InputsScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputsScriptInput)(nil)).Elem(), &InputsScript{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsScriptPtrInput)(nil)).Elem(), &InputsScript{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsScriptArrayInput)(nil)).Elem(), InputsScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputsScriptMapInput)(nil)).Elem(), InputsScriptMap{})
	pulumi.RegisterOutputType(InputsScriptOutput{})
	pulumi.RegisterOutputType(InputsScriptPtrOutput{})
	pulumi.RegisterOutputType(InputsScriptArrayOutput{})
	pulumi.RegisterOutputType(InputsScriptMapOutput{})
}
