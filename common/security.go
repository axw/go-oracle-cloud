package common

import "errors"

type SecRuleAction string

func (s SecRuleAction) Validate() (err error) {
	if s == "" {
		return errors.New("go-oracle-cloud: Empty secure rule permission")
	}
	return nil
}

const (
	SecRulePermit SecRuleAction = "PERMIT"
	SecRuleDeny   SecRuleAction = "DENY"
)

type FlowDirection string

const (
	Egress  FlowDirection = "egress"
	Ingress FlowDirection = "ingress"
)

