package models

import "github.com/FatWang1/fatwang-go-utils/desc/set"

const (
	JointlySign = "jointly_sign"
	SerialSign  = "serial_sign"
	AnyoneSign  = "anyone_sign"

	Running  = "running"
	Passed   = "passed"
	Rejected = "rejected"

	Reject = "reject"
)

var (
	TicketStatus     = set.Setify(Running, Passed, Rejected)
	DisposalSignType = set.Setify(JointlySign, SerialSign, AnyoneSign)
)
