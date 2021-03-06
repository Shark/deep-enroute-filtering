package rules

import (
  "fmt"

  "gitlab.hpi.de/felix.seidel/iotsec-enroute-filtering/filter/types"

  "github.com/Shark/canopus"
)

type MethodRule struct {
  AllowedMethods []string
}

func (r MethodRule) Process(message *types.COAPMessage) types.RuleProcessingResult {
  methodString := canopus.MethodString(message.Message.GetCode())

  for _, method := range r.AllowedMethods {
    if(method == methodString) {
      return types.RuleProcessingResult{Rule: r, Allowed: true}
    }
  }

  ruleMessage := fmt.Sprintf("%s is not allowed, only %v", methodString, r.AllowedMethods)
  return types.RuleProcessingResult{
    false,
    r,
    &ruleMessage,
  }
}

func (r MethodRule) Name() string {
  return "MethodRule"
}
