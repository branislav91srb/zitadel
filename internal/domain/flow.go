package domain

import "strconv"

type FlowState int32

const (
	FlowStateActive FlowState = iota
	FlowStateInactive
	flowStateCount
)

func (s FlowState) Valid() bool {
	return s >= 0 && s < flowStateCount
}

type FlowType int32

const (
	FlowTypeUnspecified FlowType = iota
	FlowTypeExternalAuthentication
	flowTypeCount
)

func (s FlowType) Valid() bool {
	return s > 0 && s < flowTypeCount
}

func (s FlowType) HasTrigger(triggerType TriggerType) bool {
	switch triggerType {
	case TriggerTypePostAuthentication:
		return s == FlowTypeExternalAuthentication
	case TriggerTypePreCreation:
		return s == FlowTypeExternalAuthentication
	case TriggerTypePostCreation:
		return s == FlowTypeExternalAuthentication
	default:
		return false
	}
}

func (s FlowType) String() string {
	if !s.Valid() {
		return FlowTypeUnspecified.String()
	}
	return strconv.Itoa(int(s))
}

type TriggerType int32

const (
	TriggerTypeUnspecified TriggerType = iota
	TriggerTypePostAuthentication
	TriggerTypePreCreation
	TriggerTypePostCreation
	triggerTypeCount
)

func (s TriggerType) Valid() bool {
	return s >= 0 && s < triggerTypeCount
}

func (s TriggerType) String() string {
	if !s.Valid() {
		return TriggerTypeUnspecified.String()
	}
	return strconv.Itoa(int(s))
}
