// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCognitoIdentityPoolInvalidIdentityPoolNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_identity_pool" "foo" {
	identity_pool_name = "identity:pool"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCognitoIdentityPoolInvalidIdentityPoolNameRule(),
					Message: fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage("identity:pool"), `^[\w\s+=,.@-]+$`),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_identity_pool" "foo" {
	identity_pool_name = "identity pool"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCognitoIdentityPoolInvalidIdentityPoolNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
