package test

import (
	"testing"
	"time"

	ct "github.com/daviddengcn/go-colortext"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func Test_E2eInfraExample(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	endPoint := terraform.OutputRequired(t, terraformOptions, "app_link_infra")
	ValidateModule(t, endPoint)

}

func ValidateModule(t *testing.T, endPoint string) {
	http_helper.HttpGetWithRetryWithCustomValidation(
		t, endPoint, nil, 5, 5*time.Second,
		func(statusCode int, body string) bool {
			return statusCode == 404
		},
	)
	ct.Foreground(ct.Green, true)
	logger.Logf(t, "Validation complete! "+endPoint)
	ct.Foreground(ct.White, false)
}
