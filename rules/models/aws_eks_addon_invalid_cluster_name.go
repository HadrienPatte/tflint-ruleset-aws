// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEksAddonInvalidClusterNameRule checks the pattern is valid
type AwsEksAddonInvalidClusterNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEksAddonInvalidClusterNameRule returns new rule with default attributes
func NewAwsEksAddonInvalidClusterNameRule() *AwsEksAddonInvalidClusterNameRule {
	return &AwsEksAddonInvalidClusterNameRule{
		resourceType:  "aws_eks_addon",
		attributeName: "cluster_name",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[0-9A-Za-z][A-Za-z0-9\-_]*`),
	}
}

// Name returns the rule name
func (r *AwsEksAddonInvalidClusterNameRule) Name() string {
	return "aws_eks_addon_invalid_cluster_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEksAddonInvalidClusterNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEksAddonInvalidClusterNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEksAddonInvalidClusterNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEksAddonInvalidClusterNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"cluster_name must be 100 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"cluster_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9A-Za-z][A-Za-z0-9\-_]*`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}