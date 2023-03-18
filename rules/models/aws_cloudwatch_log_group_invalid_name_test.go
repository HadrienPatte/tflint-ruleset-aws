// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCloudwatchLogGroupInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_log_group" "foo" {
	name = "Yoda:prod"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCloudwatchLogGroupInvalidNameRule(),
					Message: fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage("Yoda:prod"), `^[\.\-_/#A-Za-z0-9]+$`),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_log_group" "foo" {
	name = "Yada"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCloudwatchLogGroupInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
