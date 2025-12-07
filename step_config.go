package punched_tape

import (
	"fmt"

	"github.com/victorwong171/punched-tape/models"
	"gopkg.in/errgo.v2/errors"
)

type NextStepBuilder struct {
	option models.NextStep
}

// NewNextStepBuilder 创建下一步骤构建器，必填字段在构造函数中指定
func NewNextStepBuilder(step, operation string) *NextStepBuilder {
	return &NextStepBuilder{
		option: models.NextStep{
			Step:      step,
			Operation: operation,
		},
	}
}

// Build 构建NextStep对象，包含验证
func (b *NextStepBuilder) Build() (*models.NextStep, error) {
	// 验证必填字段
	if b.option.Step == "" {
		return nil, fmt.Errorf("step is required")
	}
	if b.option.Operation == "" {
		return nil, fmt.Errorf("operation is required")
	}

	return &b.option, nil
}

// BuildOrPanic 构建NextStep对象，验证失败时panic
func (b *NextStepBuilder) BuildOrPanic() *models.NextStep {
	nextStep, err := b.Build()
	if err != nil {
		panic(err)
	}
	return nextStep
}

type DisposalBuilder struct {
	option models.Disposal
}

func NewDisposalBuilder() *DisposalBuilder {
	return &DisposalBuilder{
		option: models.Disposal{
			JointSignRate: 0.0,
		},
	}
}

// SetSignType 设置签名类型
func (b *DisposalBuilder) SetSignType(signType string) *DisposalBuilder {
	b.option.SignType = signType
	return b
}

// SetJointSignRate 设置联合签名比例
func (b *DisposalBuilder) SetJointSignRate(rate float32) *DisposalBuilder {
	b.option.JointSignRate = rate
	return b
}

// Build 构建Disposal对象，包含验证
func (b *DisposalBuilder) Build() (*models.Disposal, error) {
	// 验证签名类型
	if b.option.SignType != "" {
		validSignTypes := map[string]bool{
			models.JointlySign: true,
			models.SerialSign:  true,
			models.AnyoneSign:  true,
		}
		if !validSignTypes[b.option.SignType] {
			return nil, fmt.Errorf("invalid sign_type: %s", b.option.SignType)
		}

		// 如果是联合签名，验证比例
		if b.option.SignType == models.JointlySign {
			if b.option.JointSignRate <= 0 || b.option.JointSignRate > 1 {
				return nil, fmt.Errorf("joint_sign_rate must be between 0 and 1")
			}
		}
	}

	return &b.option, nil
}

// BuildOrPanic 构建Disposal对象，验证失败时panic
func (b *DisposalBuilder) BuildOrPanic() *models.Disposal {
	disposal, err := b.Build()
	if err != nil {
		panic(err)
	}
	return disposal
}

type StepConfigBuilder struct {
	option *models.StepConfig
}

// NewStepConfigBuilder 创建步骤配置构建器，必填字段在构造函数中指定
func NewStepConfigBuilder(step, state string) *StepConfigBuilder {
	return &StepConfigBuilder{
		option: &models.StepConfig{
			Step:     step,
			State:    state,
			Operator: make([]string, 0),
			Next:     make([]*models.NextStep, 0),
		},
	}
}

// SetOperator 设置操作人列表
func (b *StepConfigBuilder) SetOperator(operator []string) *StepConfigBuilder {
	b.option.Operator = operator
	return b
}

// AddOperator 添加操作人
func (b *StepConfigBuilder) AddOperator(operator ...string) *StepConfigBuilder {
	b.option.Operator = append(b.option.Operator, operator...)
	return b
}

// SetNext 设置下一步骤列表
func (b *StepConfigBuilder) SetNext(next []*models.NextStep) *StepConfigBuilder {
	b.option.Next = next
	return b
}

// AddNext 添加下一步骤
func (b *StepConfigBuilder) AddNext(next ...*models.NextStep) *StepConfigBuilder {
	b.option.Next = append(b.option.Next, next...)
	return b
}

// AddNextStep 便捷方法：添加下一步骤
func (b *StepConfigBuilder) AddNextStep(step, operation string) *StepConfigBuilder {
	nextStep := &models.NextStep{
		Step:      step,
		Operation: operation,
	}
	b.option.Next = append(b.option.Next, nextStep)
	return b
}

// SetDisposal 设置处置方式
func (b *StepConfigBuilder) SetDisposal(disposal models.Disposal) *StepConfigBuilder {
	b.option.Disposal = disposal
	return b
}

// SetDisposalSignType 设置处置方式的签名类型
func (b *StepConfigBuilder) SetDisposalSignType(signType string) *StepConfigBuilder {
	b.option.Disposal.SignType = signType
	return b
}

// SetDisposalJointSignRate 设置联合签名比例
func (b *StepConfigBuilder) SetDisposalJointSignRate(rate float32) *StepConfigBuilder {
	b.option.Disposal.JointSignRate = rate
	return b
}

// Build 构建StepConfig对象，包含验证
func (b *StepConfigBuilder) Build() (*models.StepConfig, error) {
	// 验证处置方式
	if !models.DisposalSignType.HasKey(b.option.Disposal.SignType) {
		return nil, errors.New(fmt.Sprintf("invalid disposal sign type: %s", b.option.Disposal.SignType))
	}
	if b.option.Disposal.SignType == models.JointlySign {
		if b.option.Disposal.JointSignRate < 0 || b.option.Disposal.JointSignRate > 1 {
			return nil, errors.New(fmt.Sprintf("invalid disposal joint sign rate: %f", b.option.Disposal.JointSignRate))
		}
	}

	return b.option, nil
}

// BuildOrPanic 构建StepConfig对象，验证失败时panic
func (b *StepConfigBuilder) BuildOrPanic() *models.StepConfig {
	config, err := b.Build()
	if err != nil {
		panic(err)
	}
	return config
}
