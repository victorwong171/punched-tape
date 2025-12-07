package ticket

import (
	"errors"

	"github.com/victorwong171/punched-tape/models"

	"github.com/victorwong171/go-utils/desc/set"
	"github.com/victorwong171/go-utils/utils"
)

var (
	ErrOperatorNotInOperatorList = errors.New("operator not in operator list")
	ErrAlreadyApproved           = errors.New("already approved")
	ErrBadArguments              = errors.New("bad arguments")
	ErrInvalidStep               = errors.New("invalid step")
)

type Approval interface {
	Approval(next, operation, operator string, ticket *models.Ticket, config *models.StepConfig) (*models.Ticket, error)
}

type DisposalHandler interface {
}

var updateStrategy = map[string]func(operator string, ticket *models.Ticket, jointlySignRate float32, nextStep *models.StepConfig, endStep []string) *models.Ticket{
	models.JointlySign: jointlySignUpdater,
	models.SerialSign:  serialSignUpdater,
	models.AnyoneSign:  anyoneSignUpdater,
}

type Helper struct {
}

func (h *Helper) Approval(
	next,
	operation,
	operator string,
	admin bool,
	endStep []string,
	ticket *models.Ticket,
	stepConfig map[string]*models.StepConfig) (*models.Ticket, error) {
	if len(next) == 0 || len(operation) == 0 || len(operator) == 0 {
		return nil, ErrBadArguments
	}
	if ticket.Status != models.Running {
		return nil, ErrAlreadyApproved
	}

	step := stepConfig[ticket.Step]
	if step == nil {
		return nil, ErrInvalidStep
	}

	// todo: 已经操作过的人是否可以 reject？
	if utils.Contain(ticket.OperatedUser, operator) {
		return nil, ErrAlreadyApproved
	}

	if !(admin || utils.Contain(ticket.Operator, operator)) {
		return nil, ErrOperatorNotInOperatorList
	}

	var reachable bool
	for _, nextStep := range step.GetNext() {
		if nextStep.GetStep() == next && nextStep.GetOperation() == operation {
			reachable = true
		}
	}
	if !reachable {
		return nil, ErrInvalidStep
	}
	nextStep := stepConfig[next]
	ticket = updateStrategy[step.Disposal.SignType](operator, ticket, step.Disposal.JointSignRate, nextStep, endStep)
	return ticket, nil
}

func updateTicket(ticket *models.Ticket, nextStep *models.StepConfig, endStep []string) *models.Ticket {

	ticket.Step = nextStep.Step
	endStepSet := set.Setify(endStep...)
	if endStepSet.HasKey(nextStep.Step) {
		ticket.Operator = nil
		ticket.OperatedUser = nil
		ticket.Status = models.Passed
	} else {
		ticket.Operator = nextStep.Operator
		ticket.OperatedUser = nil
	}
	return ticket
}

func jointlySignUpdater(operator string, ticket *models.Ticket, jointlySignRate float32, nextStep *models.StepConfig, endStep []string) *models.Ticket {
	userSet := set.Setify(ticket.OperatedUser...)
	userSet.Set(ticket.Operator...)
	passRate := float32(1+len(ticket.OperatedUser)) / float32(userSet.Len())
	if passRate < jointlySignRate {
		ticket.Operator = utils.RemoveItemByValue(ticket.Operator, operator)
		ticket.OperatedUser = append(ticket.OperatedUser, operator)
	} else {
		ticket = updateTicket(ticket, nextStep, endStep)
	}
	return ticket
}

func serialSignUpdater(operator string, ticket *models.Ticket, _ float32, nextStep *models.StepConfig, endStep []string) *models.Ticket {
	ticket.Operator = utils.RemoveItemByValue(ticket.Operator, operator)
	if len(ticket.Operator) != 0 {
		ticket.OperatedUser = append(ticket.OperatedUser, operator)
	} else {
		ticket = updateTicket(ticket, nextStep, endStep)
	}
	return ticket
}

func anyoneSignUpdater(_ string, ticket *models.Ticket, _ float32, nextStep *models.StepConfig, endStep []string) *models.Ticket {
	return updateTicket(ticket, nextStep, endStep)
}
