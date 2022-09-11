package action

import (
	"google.golang.org/protobuf/types/known/durationpb"

	object_grpc "github.com/zitadel/zitadel/internal/api/grpc/object"
	"github.com/zitadel/zitadel/internal/domain"
	"github.com/zitadel/zitadel/internal/query"
	action_pb "github.com/zitadel/zitadel/pkg/grpc/action"
)

func FlowTypeToDomain(flowType string) domain.FlowType {
	switch flowType {
	case "FLOW_TYPE_EXTERNAL_AUTHENTICATION", "1":
		return domain.FlowTypeExternalAuthentication
	default:
		return domain.FlowTypeUnspecified
	}
}

// TriggerTypeToDomain maps the pb type to domain
// for backward compatability the old enum identifiers are mapped as well
func TriggerTypeToDomain(triggerType string) domain.TriggerType {
	switch triggerType {
	case "TRIGGER_TYPE_POST_AUTHENTICATION", "1":
		return domain.TriggerTypePostAuthentication
	case "TRIGGER_TYPE_PRE_CREATION", "2":
		return domain.TriggerTypePreCreation
	case "TRIGGER_TYPE_POST_CREATION", "3":
		return domain.TriggerTypePostCreation
	default:
		return domain.TriggerTypeUnspecified
	}
}

func FlowToPb(flow *query.Flow) *action_pb.Flow {
	return &action_pb.Flow{
		Type:           flow.Type.String(),
		Details:        object_grpc.ChangeToDetailsPb(flow.Sequence, flow.ChangeDate, flow.ResourceOwner),
		State:          action_pb.FlowState_FLOW_STATE_ACTIVE, //TODO: state in next release
		TriggerActions: TriggerActionsToPb(flow.TriggerActions),
	}
}

func TriggerActionToPb(trigger domain.TriggerType, actions []*query.Action) *action_pb.TriggerAction {
	return &action_pb.TriggerAction{
		TriggerType: trigger.String(),
		Actions:     ActionsToPb(actions),
	}
}

func TriggerActionsToPb(triggers map[domain.TriggerType][]*query.Action) []*action_pb.TriggerAction {
	list := make([]*action_pb.TriggerAction, 0)
	for trigger, actions := range triggers {
		list = append(list, TriggerActionToPb(trigger, actions))
	}
	return list
}

func ActionsToPb(actions []*query.Action) []*action_pb.Action {
	list := make([]*action_pb.Action, len(actions))
	for i, action := range actions {
		list[i] = ActionToPb(action)
	}
	return list
}

func ActionToPb(action *query.Action) *action_pb.Action {
	return &action_pb.Action{
		Id:            action.ID,
		Details:       object_grpc.ChangeToDetailsPb(action.Sequence, action.ChangeDate, action.ResourceOwner),
		State:         ActionStateToPb(action.State),
		Name:          action.Name,
		Script:        action.Script,
		Timeout:       durationpb.New(action.Timeout),
		AllowedToFail: action.AllowedToFail,
	}
}

func ActionStateToPb(state domain.ActionState) action_pb.ActionState {
	switch state {
	case domain.ActionStateActive:
		return action_pb.ActionState_ACTION_STATE_ACTIVE
	case domain.ActionStateInactive:
		return action_pb.ActionState_ACTION_STATE_INACTIVE
	default:
		return action_pb.ActionState_ACTION_STATE_UNSPECIFIED
	}
}

func ActionNameQuery(q *action_pb.ActionNameQuery) (query.SearchQuery, error) {
	return query.NewActionNameSearchQuery(object_grpc.TextMethodToQuery(q.Method), q.Name)
}

func ActionStateQuery(q *action_pb.ActionStateQuery) (query.SearchQuery, error) {
	return query.NewActionStateSearchQuery(ActionStateToDomain(q.State))
}
func ActionIDQuery(q *action_pb.ActionIDQuery) (query.SearchQuery, error) {
	return query.NewActionIDSearchQuery(q.Id)
}

func ActionStateToDomain(state action_pb.ActionState) domain.ActionState {
	switch state {
	case action_pb.ActionState_ACTION_STATE_ACTIVE:
		return domain.ActionStateActive
	case action_pb.ActionState_ACTION_STATE_INACTIVE:
		return domain.ActionStateInactive
	default:
		return domain.ActionStateUnspecified
	}
}
