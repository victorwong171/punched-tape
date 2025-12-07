package punched_tape

import (
	"github.com/victorwong171/punched-tape/models"
	"github.com/victorwong171/punched-tape/ticket/template"
)

type TemplateBuilder struct {
	option models.TicketTemplate
}

// NewTemplateBuilder 创建模板构建器，必填字段在构造函数中指定
func NewTemplateBuilder(uid, startStep string) *TemplateBuilder {
	return &TemplateBuilder{
		option: models.TicketTemplate{
			Uid:       uid,
			StartStep: startStep,
			EndStep:   make([]string, 0),
			Config:    make([]*models.StepConfig, 0),
			Builtin:   false, // 默认非内置
		},
	}
}

// SetEndStep 设置结束步骤列表
func (b *TemplateBuilder) SetEndStep(endStep []string) *TemplateBuilder {
	b.option.EndStep = endStep
	return b
}

// AddEndStep 添加结束步骤
func (b *TemplateBuilder) AddEndStep(endStep ...string) *TemplateBuilder {
	b.option.EndStep = append(b.option.EndStep, endStep...)
	return b
}

// SetBuiltin 设置是否为内置模板
func (b *TemplateBuilder) SetBuiltin(builtin bool) *TemplateBuilder {
	b.option.Builtin = builtin
	return b
}

// SetConfig 设置步骤配置列表
func (b *TemplateBuilder) SetConfig(config []*models.StepConfig) *TemplateBuilder {
	b.option.Config = config
	return b
}

// AddConfig 添加步骤配置
func (b *TemplateBuilder) AddConfig(config ...*models.StepConfig) *TemplateBuilder {
	b.option.Config = append(b.option.Config, config...)
	return b
}

// AddStepConfig 便捷方法：添加步骤配置
func (b *TemplateBuilder) AddStepConfig(step, state string, operator []string) *TemplateBuilder {
	stepConfig := &models.StepConfig{
		Step:     step,
		State:    state,
		Operator: operator,
		Next:     make([]*models.NextStep, 0),
	}
	b.option.Config = append(b.option.Config, stepConfig)
	return b
}

// Build 构建TicketTemplate对象，包含验证
func (b *TemplateBuilder) Build() (*models.TicketTemplate, error) {
	// 验证配置
	validator := template.NewValidator()
	if err := validator.Validate(b.option); err != nil {
		return nil, err
	}

	return &b.option, nil
}

// BuildOrPanic 构建TicketTemplate对象，验证失败时panic
func (b *TemplateBuilder) BuildOrPanic() *models.TicketTemplate {
	template, err := b.Build()
	if err != nil {
		panic(err)
	}
	return template
}
