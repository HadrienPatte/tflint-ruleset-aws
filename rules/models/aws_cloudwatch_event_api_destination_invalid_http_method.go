// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule checks the pattern is valid
type AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule returns new rule with default attributes
func NewAwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule() *AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule {
	return &AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule{
		resourceType:  "aws_cloudwatch_event_api_destination",
		attributeName: "http_method",
		enum: []string{
			"POST",
			"GET",
			"HEAD",
			"OPTIONS",
			"PUT",
			"PATCH",
			"DELETE",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule) Name() string {
	return "aws_cloudwatch_event_api_destination_invalid_http_method"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventAPIDestinationInvalidHTTPMethodRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as http_method`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}