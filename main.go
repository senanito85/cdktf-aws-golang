package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/instance"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	//"github.com/cdktf/cdktf-provider-aws-go/aws/v10/instance"   #not needed
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region: jsii.String("ap-southeast-2"),
	})

	instance := instance.NewInstance(stack, jsii.String("compute"), &instance.InstanceConfig{
		Ami:          jsii.String("ami-0361bbf2b99f46c1d"),
		InstanceType: jsii.String("t2.micro"),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	stack := NewMyStack(app, "aws_instance")
	// To Switch to Remote Backend use below block
	//cdktf.NewRemoteBackend(stack, &cdktf.RemoteBackendProps{
	// 	Hostname:     jsii.String("app.terraform.io"),
	// 	Organization: jsii.String("WooliesX"),
	// 	Workspaces:   cdktf.NewNamedRemoteWorkspace(jsii.String("learn-cdktf")),
	// })
	cdktf.NewLocalBackend(stack, &cdktf.LocalBackendConfig{
		Path: jsii.String("./"),
	})

	app.Synth()
}
