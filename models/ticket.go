package models

import "github.com/FatWang1/fatwang-go-utils/utils"

const (
	JointlySign = "jointly_sign"
	SerialSign  = "serial_sign"
	AnyoneSign  = "anyone_sign"

	Running  = "running"
	Passed   = "passed"
	Rejected = "rejected"

	Reject = "reject"
)

type Ticket struct {
	OrderNum     string   `json:"order_num"`     // 工单号
	Name         string   `json:"name"`          // 工单名称
	Status       string   `json:"status"`        // running/passed/rejected
	Uid          string   `json:"uid"`           // 工单唯一标识
	Step         string   `json:"step"`          // 当前步骤
	Operator     []string `json:"operator"`      // 操作人列表
	OperatedUser []string `json:"operated_user"` // 在Disposal.SignType为jointly_sign/serial_sign时使用
	Memo         string   `json:"memo"`          // 备注
}

// Getter methods for Ticket
func (t *Ticket) GetOrderNum() string {
	return utils.TernaryOperator(t == nil, "", t.OrderNum)
}

func (t *Ticket) GetName() string {
	return utils.TernaryOperator(t == nil, "", t.Name)
}

func (t *Ticket) GetStatus() string {
	return utils.TernaryOperator(t == nil, "", t.Status)
}

func (t *Ticket) GetUid() string {
	return utils.TernaryOperator(t == nil, "", t.Uid)
}

func (t *Ticket) GetStep() string {
	return utils.TernaryOperator(t == nil, "", t.Step)
}

func (t *Ticket) GetOperator() []string {
	return utils.TernaryOperator(t == nil, nil, t.Operator)
}

func (t *Ticket) GetOperatedUser() []string {
	return utils.TernaryOperator(t == nil, nil, t.OperatedUser)
}

func (t *Ticket) GetMemo() string {
	return utils.TernaryOperator(t == nil, "", t.Memo)
}

// Setter methods for Ticket
func (t *Ticket) SetName(name string) {
	if t != nil {
		t.Name = name
	}
}

func (t *Ticket) SetOrderNum(orderNum string) {
	if t != nil {
		t.OrderNum = orderNum
	}
}

func (t *Ticket) SetStatus(status string) {
	if t != nil {
		t.Status = status
	}
}

func (t *Ticket) SetUid(uid string) {
	if t != nil {
		t.Uid = uid
	}
}

func (t *Ticket) SetStep(step string) {
	if t != nil {
		t.Step = step
	}
}

func (t *Ticket) SetOperator(operator []string) {
	if t != nil {
		t.Operator = operator
	}
}

func (t *Ticket) SetOperatedUser(operatedUser []string) {
	if t != nil {
		t.OperatedUser = operatedUser
	}
}

func (t *Ticket) SetMemo(memo string) {
	if t != nil {
		t.Memo = memo
	}
}

// Add methods for slice fields
func (t *Ticket) AddOperator(operator ...string) {
	if t != nil {
		t.Operator = append(t.Operator, operator...)
	}
}

func (t *Ticket) AddOperatedUser(operatedUser ...string) {
	if t != nil {
		t.OperatedUser = append(t.OperatedUser, operatedUser...)
	}
}

type Disposal struct {
	SignType      string  `json:"sign_type"`       // jointly_sign/serial_sign/anyone_sign
	JointSignRate float32 `json:"joint_sign_rate"` // 仅jointly_sign时使用
}

// Getter methods for Disposal
func (d *Disposal) GetSignType() string {
	return utils.TernaryOperator(d == nil, "", d.SignType)
}

func (d *Disposal) GetJointSignRate() float32 {
	return utils.TernaryOperator(d == nil, 0.0, d.JointSignRate)
}

// Setter methods for Disposal
func (d *Disposal) SetSignType(signType string) {
	if d != nil {
		d.SignType = signType
	}
}

func (d *Disposal) SetJointSignRate(rate float32) {
	if d != nil {
		d.JointSignRate = rate
	}
}

// 发起工单时 可以直接使用模版 或者自定义模版 自定义模版需要
type TicketTemplate struct {
	Uid       string        `json:"uid"`        // 模板唯一标识
	Name      string        `json:"name"`       // 模板名称
	EndStep   []string      `json:"end_step"`   // 结束节点
	StartStep string        `json:"start_step"` // 开始节点
	Config    []*StepConfig `json:"config"`     // 配置
	Builtin   bool          `json:"builtin"`    // 是否内置
}

// Getter methods for TicketTemplate
func (tt *TicketTemplate) GetName() string {
	return utils.TernaryOperator(tt == nil, "", tt.Name)
}

func (tt *TicketTemplate) GetUid() string {
	return utils.TernaryOperator(tt == nil, "", tt.Uid)
}

func (tt *TicketTemplate) GetEndStep() []string {
	return utils.TernaryOperator(tt == nil, nil, tt.EndStep)
}

func (tt *TicketTemplate) GetStartStep() string {
	return utils.TernaryOperator(tt == nil, "", tt.StartStep)
}

func (tt *TicketTemplate) GetConfig() []*StepConfig {
	return utils.TernaryOperator(tt == nil, nil, tt.Config)
}

func (tt *TicketTemplate) GetBuiltin() bool {
	return utils.TernaryOperator(tt == nil, false, tt.Builtin)
}

// Setter methods for TicketTemplate
func (tt *TicketTemplate) SetName(name string) {
	if tt != nil {
		tt.Name = name
	}
}

func (tt *TicketTemplate) SetUid(uid string) {
	if tt != nil {
		tt.Uid = uid
	}
}

func (tt *TicketTemplate) SetEndStep(endStep []string) {
	if tt != nil {
		tt.EndStep = endStep
	}
}

func (tt *TicketTemplate) SetStartStep(startStep string) {
	if tt != nil {
		tt.StartStep = startStep
	}
}

func (tt *TicketTemplate) SetConfig(config []*StepConfig) {
	if tt != nil {
		tt.Config = config
	}
}

func (tt *TicketTemplate) SetBuiltin(builtin bool) {
	if tt != nil {
		tt.Builtin = builtin
	}
}

// Add methods for slice fields
func (tt *TicketTemplate) AddEndStep(endStep ...string) {
	if tt != nil {
		tt.EndStep = append(tt.EndStep, endStep...)
	}
}

func (tt *TicketTemplate) AddConfig(config ...*StepConfig) {
	if tt != nil {
		tt.Config = append(tt.Config, config...)
	}
}

type StepConfig struct {
	Step     string      `json:"step"`     // 步骤名
	State    string      `json:"state"`    // 步骤所属状态
	Operator []string    `json:"operator"` // 预设操作人
	Next     []*NextStep `json:"next"`     // 下一节点
	Disposal Disposal    `json:"disposal"` // 处置方式
}

// Getter methods for StepConfig
func (sc *StepConfig) GetStep() string {
	return utils.TernaryOperator(sc == nil, "", sc.Step)
}

func (sc *StepConfig) GetState() string {
	return utils.TernaryOperator(sc == nil, "", sc.State)
}

func (sc *StepConfig) GetOperator() []string {
	return utils.TernaryOperator(sc == nil, nil, sc.Operator)
}

func (sc *StepConfig) GetNext() []*NextStep {
	return utils.TernaryOperator(sc == nil, nil, sc.Next)
}

func (sc *StepConfig) GetDisposal() Disposal {
	return utils.TernaryOperator(sc == nil, Disposal{}, sc.Disposal)
}

// Setter methods for StepConfig
func (sc *StepConfig) SetStep(step string) {
	if sc != nil {
		sc.Step = step
	}
}

func (sc *StepConfig) SetState(state string) {
	if sc != nil {
		sc.State = state
	}
}

func (sc *StepConfig) SetOperator(operator []string) {
	if sc != nil {
		sc.Operator = operator
	}
}

func (sc *StepConfig) SetNext(next []*NextStep) {
	if sc != nil {
		sc.Next = next
	}
}

func (sc *StepConfig) SetDisposal(disposal Disposal) {
	if sc != nil {
		sc.Disposal = disposal
	}
}

// Add methods for slice fields
func (sc *StepConfig) AddOperator(operator ...string) {
	if sc != nil {
		sc.Operator = append(sc.Operator, operator...)
	}
}

func (sc *StepConfig) AddNext(next ...*NextStep) {
	if sc != nil {
		sc.Next = append(sc.Next, next...)
	}
}

type NextStep struct {
	Step      string `json:"step"`      // 步骤名
	Operation string `json:"operation"` // 操作名
}

// Getter methods for NextStep
func (ns *NextStep) GetStep() string {
	return utils.TernaryOperator(ns == nil, "", ns.Step)
}

func (ns *NextStep) GetOperation() string {
	return utils.TernaryOperator(ns == nil, "", ns.Operation)
}

// Setter methods for NextStep
func (ns *NextStep) SetStep(step string) {
	if ns != nil {
		ns.Step = step
	}
}

func (ns *NextStep) SetOperation(operation string) {
	if ns != nil {
		ns.Operation = operation
	}
}
