package setdata_acl

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type RoleAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewRoleAmqpEndpoints(ch setdata_common.CommandHandler) RoleAmqpEndpoints {
	return RoleAmqpEndpoints{ch: ch}
}

func (r *RoleAmqpEndpoints) MakeCreateRoleAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateRoleCommand{}
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
