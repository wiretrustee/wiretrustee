package iptables

// Rule to handle management of rules
type Rule struct {
	ruleID    string
	ipsetName string

	specs []string
	ip    string
	dst   bool
}

// GetRuleID returns the rule id
func (r *Rule) GetRuleID() string {
	return r.ruleID
}
