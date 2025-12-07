package models

import (
	"testing"
)

func TestTicket_GetterMethods(t *testing.T) {
	// Test with valid ticket
	ticket := &Ticket{
		OrderNum:     "TICKET-001",
		Status:       Running,
		Uid:          "user123",
		Step:         "approval",
		Memo:         "test memo",
		Operator:     []string{"admin1", "admin2"},
		OperatedUser: []string{"user1", "user2"},
	}

	// Test getters with valid ticket
	if got := ticket.GetOrderNum(); got != "TICKET-001" {
		t.Errorf("Ticket.GetOrderNum() = %v, want TICKET-001", got)
	}
	if got := ticket.GetStatus(); got != Running {
		t.Errorf("Ticket.GetStatus() = %v, want %v", got, Running)
	}
	if got := ticket.GetUid(); got != "user123" {
		t.Errorf("Ticket.GetUid() = %v, want user123", got)
	}
	if got := ticket.GetStep(); got != "approval" {
		t.Errorf("Ticket.GetStep() = %v, want approval", got)
	}
	if got := ticket.GetMemo(); got != "test memo" {
		t.Errorf("Ticket.GetMemo() = %v, want test memo", got)
	}

	operators := ticket.GetOperator()
	if len(operators) != 2 {
		t.Errorf("Ticket.GetOperator() length = %v, want 2", len(operators))
	}
	if operators[0] != "admin1" || operators[1] != "admin2" {
		t.Errorf("Ticket.GetOperator() = %v, want [admin1 admin2]", operators)
	}

	operatedUsers := ticket.GetOperatedUser()
	if len(operatedUsers) != 2 {
		t.Errorf("Ticket.GetOperatedUser() length = %v, want 2", len(operatedUsers))
	}
	if operatedUsers[0] != "user1" || operatedUsers[1] != "user2" {
		t.Errorf("Ticket.GetOperatedUser() = %v, want [user1 user2]", operatedUsers)
	}
}

func TestTicket_SetterMethods(t *testing.T) {
	ticket := &Ticket{}

	// Test SetOrderNum
	ticket.SetOrderNum("TICKET-001")
	if ticket.OrderNum != "TICKET-001" {
		t.Errorf("SetOrderNum failed, got %v, want TICKET-001", ticket.OrderNum)
	}

	// Test SetStatus
	ticket.SetStatus(Passed)
	if ticket.Status != Passed {
		t.Errorf("SetStatus failed, got %v, want %v", ticket.Status, Passed)
	}

	// Test SetUid
	ticket.SetUid("user123")
	if ticket.Uid != "user123" {
		t.Errorf("SetUid failed, got %v, want user123", ticket.Uid)
	}

	// Test SetStep
	ticket.SetStep("approval")
	if ticket.Step != "approval" {
		t.Errorf("SetStep failed, got %v, want approval", ticket.Step)
	}

	// Test SetOperator
	operators := []string{"admin1", "admin2"}
	ticket.SetOperator(operators)
	if len(ticket.Operator) != len(operators) {
		t.Errorf("SetOperator failed, got length %v, want %v", len(ticket.Operator), len(operators))
	}

	// Test SetOperatedUser
	operatedUsers := []string{"user1", "user2"}
	ticket.SetOperatedUser(operatedUsers)
	if len(ticket.OperatedUser) != len(operatedUsers) {
		t.Errorf("SetOperatedUser failed, got length %v, want %v", len(ticket.OperatedUser), len(operatedUsers))
	}

	// Test SetMemo
	ticket.SetMemo("test memo")
	if ticket.Memo != "test memo" {
		t.Errorf("SetMemo failed, got %v, want test memo", ticket.Memo)
	}
}

func TestTicket_GetName_SetName(t *testing.T) {
	// 测试非空指针
	ticket := &Ticket{
		Name: "test ticket",
	}

	// 测试 GetName 方法
	if ticket.GetName() != "test ticket" {
		t.Errorf("expected %s, got %s", "test ticket", ticket.GetName())
	}

	// 测试 SetName 方法
	ticket.SetName("updated ticket")
	if ticket.GetName() != "updated ticket" {
		t.Errorf("expected %s, got %s", "updated ticket", ticket.GetName())
	}
}

func TestTicketTemplate_GetName_SetName(t *testing.T) {
	// 测试非空指针
	template := &TicketTemplate{
		Name: "test template",
	}

	// 测试 GetName 方法
	if template.GetName() != "test template" {
		t.Errorf("expected %s, got %s", "test template", template.GetName())
	}

	// 测试 SetName 方法
	template.SetName("updated template")
	if template.GetName() != "updated template" {
		t.Errorf("expected %s, got %s", "updated template", template.GetName())
	}
}



func TestTicket_AddMethods(t *testing.T) {
	ticket := &Ticket{
		Operator:     []string{"admin1"},
		OperatedUser: []string{"user1"},
	}

	// Test AddOperator
	ticket.AddOperator("admin2", "admin3")
	expectedOperators := []string{"admin1", "admin2", "admin3"}
	if len(ticket.Operator) != len(expectedOperators) {
		t.Errorf("AddOperator failed, got length %v, want %v", len(ticket.Operator), len(expectedOperators))
	}

	// Test AddOperatedUser
	ticket.AddOperatedUser("user2", "user3")
	expectedOperatedUsers := []string{"user1", "user2", "user3"}
	if len(ticket.OperatedUser) != len(expectedOperatedUsers) {
		t.Errorf("AddOperatedUser failed, got length %v, want %v", len(ticket.OperatedUser), len(expectedOperatedUsers))
	}
}

func TestDisposal_GetterMethods(t *testing.T) {
	disposal := &Disposal{
		SignType:      JointlySign,
		JointSignRate: 0.5,
	}

	// Test getters with valid disposal
	if got := disposal.GetSignType(); got != JointlySign {
		t.Errorf("Disposal.GetSignType() = %v, want %v", got, JointlySign)
	}
	if got := disposal.GetJointSignRate(); got != 0.5 {
		t.Errorf("Disposal.GetJointSignRate() = %v, want 0.5", got)
	}
}

func TestDisposal_SetterMethods(t *testing.T) {
	disposal := &Disposal{}

	// Test SetSignType
	disposal.SetSignType(JointlySign)
	if disposal.SignType != JointlySign {
		t.Errorf("SetSignType failed, got %v, want %v", disposal.SignType, JointlySign)
	}

	// Test SetJointSignRate
	disposal.SetJointSignRate(0.75)
	if disposal.JointSignRate != 0.75 {
		t.Errorf("SetJointSignRate failed, got %v, want 0.75", disposal.JointSignRate)
	}
}

func TestTicketTemplate_GetterMethods(t *testing.T) {
	template := &TicketTemplate{
		Uid:       "template-001",
		EndStep:   []string{"approved", "rejected"},
		StartStep: "submit",
		Config:    []*StepConfig{{Step: "step1", State: "state1"}},
		Builtin:   true,
	}

	// Test getters with valid template
	if got := template.GetUid(); got != "template-001" {
		t.Errorf("TicketTemplate.GetUid() = %v, want template-001", got)
	}
	if got := template.GetStartStep(); got != "submit" {
		t.Errorf("TicketTemplate.GetStartStep() = %v, want submit", got)
	}
	if got := template.GetBuiltin(); got != true {
		t.Errorf("TicketTemplate.GetBuiltin() = %v, want true", got)
	}

	endSteps := template.GetEndStep()
	if len(endSteps) != 2 {
		t.Errorf("TicketTemplate.GetEndStep() length = %v, want 2", len(endSteps))
	}
	if endSteps[0] != "approved" || endSteps[1] != "rejected" {
		t.Errorf("TicketTemplate.GetEndStep() = %v, want [approved rejected]", endSteps)
	}

	configs := template.GetConfig()
	if len(configs) != 1 {
		t.Errorf("TicketTemplate.GetConfig() length = %v, want 1", len(configs))
	}
}

func TestTicketTemplate_SetterMethods(t *testing.T) {
	template := &TicketTemplate{}

	// Test SetUid
	template.SetUid("template-001")
	if template.Uid != "template-001" {
		t.Errorf("SetUid failed, got %v, want template-001", template.Uid)
	}

	// Test SetEndStep
	endSteps := []string{"approved", "rejected"}
	template.SetEndStep(endSteps)
	if len(template.EndStep) != len(endSteps) {
		t.Errorf("SetEndStep failed, got length %v, want %v", len(template.EndStep), len(endSteps))
	}

	// Test SetStartStep
	template.SetStartStep("submit")
	if template.StartStep != "submit" {
		t.Errorf("SetStartStep failed, got %v, want submit", template.StartStep)
	}

	// Test SetConfig
	configs := []*StepConfig{
		{Step: "step1", State: "state1"},
		{Step: "step2", State: "state2"},
	}
	template.SetConfig(configs)
	if len(template.Config) != len(configs) {
		t.Errorf("SetConfig failed, got length %v, want %v", len(template.Config), len(configs))
	}

	// Test SetBuiltin
	template.SetBuiltin(true)
	if template.Builtin != true {
		t.Errorf("SetBuiltin failed, got %v, want true", template.Builtin)
	}
}

func TestTicketTemplate_AddMethods(t *testing.T) {
	template := &TicketTemplate{
		EndStep: []string{"approved"},
		Config:  []*StepConfig{{Step: "step1", State: "state1"}},
	}

	// Test AddEndStep
	template.AddEndStep("rejected", "cancelled")
	expectedEndSteps := []string{"approved", "rejected", "cancelled"}
	if len(template.EndStep) != len(expectedEndSteps) {
		t.Errorf("AddEndStep failed, got length %v, want %v", len(template.EndStep), len(expectedEndSteps))
	}

	// Test AddConfig
	template.AddConfig(&StepConfig{Step: "step2", State: "state2"})
	if len(template.Config) != 2 {
		t.Errorf("AddConfig failed, got length %v, want 2", len(template.Config))
	}
}

func TestStepConfig_GetterMethods(t *testing.T) {
	stepConfig := &StepConfig{
		Step:     "approval",
		State:    "pending",
		Operator: []string{"admin1", "admin2"},
		Next:     []*NextStep{{Step: "next1", Operation: "op1"}},
		Disposal: Disposal{SignType: JointlySign, JointSignRate: 0.5},
	}

	// Test getters with valid step config
	if got := stepConfig.GetStep(); got != "approval" {
		t.Errorf("StepConfig.GetStep() = %v, want approval", got)
	}
	if got := stepConfig.GetState(); got != "pending" {
		t.Errorf("StepConfig.GetState() = %v, want pending", got)
	}

	operators := stepConfig.GetOperator()
	if len(operators) != 2 {
		t.Errorf("StepConfig.GetOperator() length = %v, want 2", len(operators))
	}
	if operators[0] != "admin1" || operators[1] != "admin2" {
		t.Errorf("StepConfig.GetOperator() = %v, want [admin1 admin2]", operators)
	}

	nextSteps := stepConfig.GetNext()
	if len(nextSteps) != 1 {
		t.Errorf("StepConfig.GetNext() length = %v, want 1", len(nextSteps))
	}

	disposal := stepConfig.GetDisposal()
	if disposal.SignType != JointlySign {
		t.Errorf("StepConfig.GetDisposal().SignType = %v, want %v", disposal.SignType, JointlySign)
	}
	if disposal.JointSignRate != 0.5 {
		t.Errorf("StepConfig.GetDisposal().JointSignRate = %v, want 0.5", disposal.JointSignRate)
	}
}

func TestStepConfig_SetterMethods(t *testing.T) {
	stepConfig := &StepConfig{}

	// Test SetStep
	stepConfig.SetStep("approval")
	if stepConfig.Step != "approval" {
		t.Errorf("SetStep failed, got %v, want approval", stepConfig.Step)
	}

	// Test SetState
	stepConfig.SetState("pending")
	if stepConfig.State != "pending" {
		t.Errorf("SetState failed, got %v, want pending", stepConfig.State)
	}

	// Test SetOperator
	operators := []string{"admin1", "admin2"}
	stepConfig.SetOperator(operators)
	if len(stepConfig.Operator) != len(operators) {
		t.Errorf("SetOperator failed, got length %v, want %v", len(stepConfig.Operator), len(operators))
	}

	// Test SetNext
	nextSteps := []*NextStep{
		{Step: "next1", Operation: "op1"},
		{Step: "next2", Operation: "op2"},
	}
	stepConfig.SetNext(nextSteps)
	if len(stepConfig.Next) != len(nextSteps) {
		t.Errorf("SetNext failed, got length %v, want %v", len(stepConfig.Next), len(nextSteps))
	}

	// Test SetDisposal
	disposal := Disposal{SignType: JointlySign, JointSignRate: 0.5}
	stepConfig.SetDisposal(disposal)
	if stepConfig.Disposal.SignType != disposal.SignType {
		t.Errorf("SetDisposal failed, got SignType %v, want %v", stepConfig.Disposal.SignType, disposal.SignType)
	}
	if stepConfig.Disposal.JointSignRate != disposal.JointSignRate {
		t.Errorf("SetDisposal failed, got JointSignRate %v, want %v", stepConfig.Disposal.JointSignRate, disposal.JointSignRate)
	}
}

func TestStepConfig_AddMethods(t *testing.T) {
	stepConfig := &StepConfig{
		Operator: []string{"admin1"},
		Next:     []*NextStep{{Step: "next1", Operation: "op1"}},
	}

	// Test AddOperator
	stepConfig.AddOperator("admin2", "admin3")
	expectedOperators := []string{"admin1", "admin2", "admin3"}
	if len(stepConfig.Operator) != len(expectedOperators) {
		t.Errorf("AddOperator failed, got length %v, want %v", len(stepConfig.Operator), len(expectedOperators))
	}

	// Test AddNext
	stepConfig.AddNext(&NextStep{Step: "next2", Operation: "op2"})
	if len(stepConfig.Next) != 2 {
		t.Errorf("AddNext failed, got length %v, want 2", len(stepConfig.Next))
	}
}

func TestNextStep_GetterMethods(t *testing.T) {
	nextStep := &NextStep{
		Step:      "next1",
		Operation: "approve",
	}

	// Test getters with valid next step
	if got := nextStep.GetStep(); got != "next1" {
		t.Errorf("NextStep.GetStep() = %v, want next1", got)
	}
	if got := nextStep.GetOperation(); got != "approve" {
		t.Errorf("NextStep.GetOperation() = %v, want approve", got)
	}
}

func TestNextStep_SetterMethods(t *testing.T) {
	nextStep := &NextStep{}

	// Test SetStep
	nextStep.SetStep("next1")
	if nextStep.Step != "next1" {
		t.Errorf("SetStep failed, got %v, want next1", nextStep.Step)
	}

	// Test SetOperation
	nextStep.SetOperation("approve")
	if nextStep.Operation != "approve" {
		t.Errorf("SetOperation failed, got %v, want approve", nextStep.Operation)
	}
}
