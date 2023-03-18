// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCognitoIdentityProviderInvalidProviderTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_identity_provider" "foo" {
	provider_type = "Apple"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCognitoIdentityProviderInvalidProviderTypeRule(),
					Message: fmt.Sprintf(`"%s" is an invalid value as %s`, truncateLongMessage("Apple"), "provider_type"),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_identity_provider" "foo" {
	provider_type = "LoginWithAmazon"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCognitoIdentityProviderInvalidProviderTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
