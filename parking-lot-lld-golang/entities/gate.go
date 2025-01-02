package entities

import "main/enums"

type Gate struct {
	Id       int
	GateType enums.GateType
}

func NewGate(id int, gateType enums.GateType) *Gate {
	return &Gate{
		Id:       id,
		GateType: gateType,
	}
}
