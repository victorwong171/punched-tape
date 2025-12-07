package punched_tape

import (
	"fmt"

	"github.com/FatWang1/punched-tape/models"
	"gopkg.in/errgo.v2/errors"
)

type TicketBuilder struct {
	option models.Ticket
}

// NewTicketBuilder 创建工单构建器，必填字段在构造函数中指定
func NewTicketBuilder(uid, orderNum, step, name string) *TicketBuilder {
	return &TicketBuilder{
		option: models.Ticket{
			Uid:          uid,
			OrderNum:     orderNum,
			Name:         name,
			Step:         step,
			Status:       models.Running, // 设置默认状态
			Operator:     make([]string, 0),
			OperatedUser: make([]string, 0),
		},
	}
}

// SetName 设置工单名称
func (b *TicketBuilder) SetName(name string) *TicketBuilder {
	b.option.Name = name
	return b
}

// SetStatus 设置工单状态
func (b *TicketBuilder) SetStatus(status string) *TicketBuilder {
	b.option.Status = status
	return b
}

// SetOperator 设置操作人列表
func (b *TicketBuilder) SetOperator(operator []string) *TicketBuilder {
	b.option.Operator = operator
	return b
}

// AddOperator 添加操作人
func (b *TicketBuilder) AddOperator(operator ...string) *TicketBuilder {
	b.option.Operator = append(b.option.Operator, operator...)
	return b
}

// SetOperatedUser 设置已操作用户列表
func (b *TicketBuilder) SetOperatedUser(operatedUser []string) *TicketBuilder {
	b.option.OperatedUser = operatedUser
	return b
}

// AddOperatedUser 添加已操作用户
func (b *TicketBuilder) AddOperatedUser(operatedUser ...string) *TicketBuilder {
	b.option.OperatedUser = append(b.option.OperatedUser, operatedUser...)
	return b
}

// SetMemo 设置备注
func (b *TicketBuilder) SetMemo(memo string) *TicketBuilder {
	b.option.Memo = memo
	return b
}

// Build 构建Ticket对象，包含验证
func (b *TicketBuilder) Build() (*models.Ticket, error) {
	// 验证状态值
	if !models.TicketStatus.HasKey(b.option.Status) {
		return nil, errors.New(fmt.Sprintf("invalid status: %s", b.option.Status))
	}

	return &b.option, nil
}

// BuildOrPanic 构建Ticket对象，验证失败时panic
func (b *TicketBuilder) BuildOrPanic() *models.Ticket {
	ticket, err := b.Build()
	if err != nil {
		panic(err)
	}
	return ticket
}
