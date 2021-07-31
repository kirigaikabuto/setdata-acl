package setdata_acl

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type PermissionAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewPermissionAmqpEndpoints(ch setdata_common.CommandHandler) PermissionAmqpEndpoints {
	return PermissionAmqpEndpoints{ch: ch}
}

func (p *PermissionAmqpEndpoints) MakeCreatePermissionAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreatePermissionCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := p.ch.ExecCommand(cmd)
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

func (p *PermissionAmqpEndpoints) MakeGetPermissionAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetPermissionCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := p.ch.ExecCommand(cmd)
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

func (p *PermissionAmqpEndpoints) MakeListPermissionAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListPermissionCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := p.ch.ExecCommand(cmd)
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

func (p *PermissionAmqpEndpoints) MakeDeletePermissionAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeletePermissionCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := p.ch.ExecCommand(cmd)
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
