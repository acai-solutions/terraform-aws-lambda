package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestLambdaUC3(t *testing.T) {
	t.Log("Starting Sample Module test")

	terraformDir := "../../examples/use-case-3"
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

	lambdaResultOutput1 := terraform.OutputMap(t, terraformModule, "use_case_3_lambda1_result")
	t.Logf("Lambda Output1: %s", lambdaResultOutput1)

	// Extract the statusCode and assert it
	statusCode1 := lambdaResultOutput1["statusCode"]
	// Print the status code
	t.Logf("Derived StatusCode1: %s", statusCode1)
	assert.Equal(t, "200", statusCode1, "Expected statusCode to be 200")

	lambdaResultOutput2 := terraform.OutputMap(t, terraformModule, "use_case_3_lambda2_result")
	t.Logf("Lambda Output2: %s", lambdaResultOutput2)

	// Extract the statusCode and assert it
	statusCode2 := lambdaResultOutput2["statusCode"]
	// Print the status code
	t.Logf("Derived StatusCode: %s", statusCode2)
	assert.Equal(t, "200", statusCode2, "Expected statusCode to be 200")
}
