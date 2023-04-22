package model

import (
	"strconv"
	"strings"
)

/*
{
   "cmd_key":"link",
   "case":"create",
   "step":0,
   "payload":"button 2 is pressed"
}
*/

type CallbackDataBot string

// may will be need
type PayloadData struct {
	ObjectType string
	ObjectID   string
}

type CallbackData struct {
	CommandKey CommandKey
	Case       Case
	Step       Step
	Payload    string
}

func (cd *CallbackData) Encode() string {
	stepNumber := strconv.Itoa(int(cd.Step))
	return strings.Join(
		[]string{
			string(cd.CommandKey),
			string(cd.Case),
			stepNumber,
			cd.Payload,
		}, ";",
	)
}

// "commandKey;case;step;payload" —> "link;create;0;online"

func (cdb *CallbackDataBot) Decode() *CallbackData {
	flowItems := strings.Split(string(*cdb), ";")
	step, err := strconv.Atoi(flowItems[2])
	if err != nil {
		return nil
	}
	return &CallbackData{
		CommandKey: CommandKey(flowItems[0]),
		Case:       Case(flowItems[1]),
		Step:       Step(step),
		Payload:    flowItems[3],
	}
}
