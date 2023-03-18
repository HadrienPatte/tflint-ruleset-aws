// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsAppsyncDatasourceInvalidTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_appsync_datasource" "foo" {
	type = "AMAZON_SIMPLEDB"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsAppsyncDatasourceInvalidTypeRule(),
					Message: fmt.Sprintf(`"%s" is an invalid value as %s`, truncateLongMessage("AMAZON_SIMPLEDB"), "type"),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_appsync_datasource" "foo" {
	type = "AWS_LAMBDA"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsAppsyncDatasourceInvalidTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
