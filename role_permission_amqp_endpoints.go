package setdata_acl

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type RolePermissionAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewRolePermissionAmqpEndpoints(ch setdata_common.CommandHandler) RolePermissionAmqpEndpoints {
	return RolePermissionAmqpEndpoints{ch: ch}
}

func(r *RolePermissionAmqpEndpoints) MakeCreateRolePermissionAmqpEndpoint() amqp.Handler{
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateRolePermissionCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := r.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func(r *RolePermissionAmqpEndpoints) MakeGetRolePermissionAmqpEndpoint() amqp.Handler{
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetRolePermissionCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := r.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func(r *RolePermissionAmqpEndpoints) MakeListRolePermissionAmqpEndpoint() amqp.Handler{
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListRolePermissionCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := r.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func(r *RolePermissionAmqpEndpoints) MakeDeleteRolePermissionAmqpEndpoint() amqp.Handler{
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeleteRolePermissionCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := r.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}