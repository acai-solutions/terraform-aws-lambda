package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestLambdaUC1(t *testing.T) {
	t.Log("Starting Sample Module test")

	terraformDir := "../../_example/use-case-1"
	backendConfig := loadBackendConfig(t)

	// Create IAM Role
	terraformPreparation := &terraform.Options{
		TerraformDir:  terraformDir,
		NoColor:       false,
		Lock:          true,
		BackendConfig: backendConfig,
		Reconfigure:   true,
		Targets: []string{
			"module.create_provisioner",
		},
	}
	defer terraform.Destroy(t, terraformPreparation)
	terraform.InitAndApply(t, terraformPreparation)

	terraformModule := &terraform.Options{
		TerraformDir:  terraformDir,
		NoColor:       false,
		Lock:          true,
		BackendConfig: backendConfig,
		Reconfigure:   true,
	}
	defer terraform.Destroy(t, terraformModule)
	terraform.InitAndApply(t, terraformModule)

	lambdaResultOutput := terraform.OutputMap(t, terraformModule, "use_case_1_lambda_result")
	t.Logf("Lambda Output: %s", lambdaResultOutput)

	// Extract the statusCode and assert it
	statusCode := lambdaResultOutput["statusCode"]
	// Print the status code
	t.Logf("Derived StatusCode: %s", statusCode)
	assert.Equal(t, "200", statusCode, "Expected statusCode to be 200")
}
