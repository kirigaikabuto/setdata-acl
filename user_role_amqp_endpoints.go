package setdata_acl

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type UserRoleAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewUserRoleAmqpEndpoints(ch setdata_common.CommandHandler) UserRoleAmqpEndpoints {
	return UserRoleAmqpEndpoints{ch: ch}
}

func (u *UserRoleAmqpEndpoints) MakeCreateUserRoleAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateUserRoleCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
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

func (u *UserRoleAmqpEndpoints) MakeGetUserRoleAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetUserRoleCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
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

func (u *UserRoleAmqpEndpoints) MakeListUserRoleAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListUserRoleCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
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

func (u *UserRoleAmqpEndpoints) MakeDeleteUserRoleAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeleteUserRoleCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
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

func (u *UserRoleAmqpEndpoints) MakeGetUserRolePermissionsAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetUserRolePermissionsCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
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
