// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCloudformationStackInvalidOnFailureRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudformation_stack" "foo" {
	on_failure = "DO_ANYTHING"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCloudformationStackInvalidOnFailureRule(),
					Message: fmt.Sprintf(`"%s" is an invalid value as %s`, truncateLongMessage("DO_ANYTHING"), "on_failure"),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudformation_stack" "foo" {
	on_failure = "DO_NOTHING"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCloudformationStackInvalidOnFailureRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
