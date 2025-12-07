package punched_tape

import (
	"testing"

	"github.com/victorwong171/punched-tape/models"
)

func TestNewTemplateBuilder(t *testing.T) {
	tests := []struct {
		name      string
		uid       string
		startStep string
		expectNil bool
	}{
		{
			name:      "valid parameters",
			uid:       "template-001",
			startStep: "submit",
			expectNil: false,
		},
		{
			name:      "empty parameters",
			uid:       "",
			startStep: "",
			expectNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewTemplateBuilder(tt.uid, tt.startStep)
			if tt.expectNil && builder != nil {
				t.Errorf("NewTemplateBuilder() expected nil, got %v", builder)
			}
			if !tt.expectNil && builder == nil {
				t.Errorf("NewTemplateBuilder() expected non-nil, got nil")
			}
		})
	}
}

func TestTemplateBuilder_SetEndStep(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	tests := []struct {
		name     string
		endStep  []string
		expected []string
	}{
		{
			name:     "set single end step",
			endStep:  []string{"approved"},
			expected: []string{"approved"},
		},
		{
			name:     "set multiple end steps",
			endStep:  []string{"approved", "rejected", "cancelled"},
			expected: []string{"approved", "rejected", "cancelled"},
		},
		{
			name:     "set empty end steps",
			endStep:  []string{},
			expected: []string{},
		},
		{
			name:     "set nil end steps",
			endStep:  nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetEndStep(tt.endStep)
			if result != builder {
				t.Errorf("SetEndStep() should return builder instance")
			}
			if len(builder.option.EndStep) != len(tt.expected) {
				t.Errorf("SetEndStep() length = %v, want %v", len(builder.option.EndStep), len(tt.expected))
				return
			}
			for i, v := range builder.option.EndStep {
				if v != tt.expected[i] {
					t.Errorf("SetEndStep()[%d] = %v, want %v", i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestTemplateBuilder_AddEndStep(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	// Test initial state
	if len(builder.option.EndStep) != 0 {
		t.Errorf("Initial end step length should be 0, got %v", len(builder.option.EndStep))
	}

	// Test adding single end step
	result := builder.AddEndStep("approved")
	if result != builder {
		t.Errorf("AddEndStep() should return builder instance")
	}
	if len(builder.option.EndStep) != 1 {
		t.Errorf("AddEndStep() length = %v, want 1", len(builder.option.EndStep))
	}
	if builder.option.EndStep[0] != "approved" {
		t.Errorf("AddEndStep()[0] = %v, want approved", builder.option.EndStep[0])
	}

	// Test adding multiple end steps
	result = builder.AddEndStep("rejected", "cancelled")
	if len(builder.option.EndStep) != 3 {
		t.Errorf("AddEndStep() length = %v, want 3", len(builder.option.EndStep))
	}
	expected := []string{"approved", "rejected", "cancelled"}
	for i, v := range builder.option.EndStep {
		if v != expected[i] {
			t.Errorf("AddEndStep()[%d] = %v, want %v", i, v, expected[i])
		}
	}
}

func TestTemplateBuilder_SetBuiltin(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	tests := []struct {
		name     string
		builtin  bool
		expected bool
	}{
		{
			name:     "set builtin true",
			builtin:  true,
			expected: true,
		},
		{
			name:     "set builtin false",
			builtin:  false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetBuiltin(tt.builtin)
			if result != builder {
				t.Errorf("SetBuiltin() should return builder instance")
			}
			if builder.option.Builtin != tt.expected {
				t.Errorf("SetBuiltin() = %v, want %v", builder.option.Builtin, tt.expected)
			}
		})
	}
}

func TestTemplateBuilder_SetConfig(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	configs := []*models.StepConfig{
		{Step: "step1", State: "state1"},
		{Step: "step2", State: "state2"},
	}

	result := builder.SetConfig(configs)
	if result != builder {
		t.Errorf("SetConfig() should return builder instance")
	}
	if len(builder.option.Config) != len(configs) {
		t.Errorf("SetConfig() length = %v, want %v", len(builder.option.Config), len(configs))
	}
}

func TestTemplateBuilder_AddConfig(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	// Test initial state
	if len(builder.option.Config) != 0 {
		t.Errorf("Initial config length should be 0, got %v", len(builder.option.Config))
	}

	// Test adding single config
	config1 := &models.StepConfig{Step: "step1", State: "state1"}
	result := builder.AddConfig(config1)
	if result != builder {
		t.Errorf("AddConfig() should return builder instance")
	}
	if len(builder.option.Config) != 1 {
		t.Errorf("AddConfig() length = %v, want 1", len(builder.option.Config))
	}

	// Test adding multiple configs
	config2 := &models.StepConfig{Step: "step2", State: "state2"}
	config3 := &models.StepConfig{Step: "step3", State: "state3"}
	result = builder.AddConfig(config2, config3)
	if len(builder.option.Config) != 3 {
		t.Errorf("AddConfig() length = %v, want 3", len(builder.option.Config))
	}
}

func TestTemplateBuilder_AddStepConfig(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	// Test initial state
	if len(builder.option.Config) != 0 {
		t.Errorf("Initial config length should be 0, got %v", len(builder.option.Config))
	}

	// Test adding step config
	result := builder.AddStepConfig("step1", "state1", []string{"operator1"})
	if result != builder {
		t.Errorf("AddStepConfig() should return builder instance")
	}
	if len(builder.option.Config) != 1 {
		t.Errorf("AddStepConfig() length = %v, want 1", len(builder.option.Config))
	}
	if builder.option.Config[0].Step != "step1" {
		t.Errorf("AddStepConfig() Step = %v, want step1", builder.option.Config[0].Step)
	}
	if builder.option.Config[0].State != "state1" {
		t.Errorf("AddStepConfig() State = %v, want state1", builder.option.Config[0].State)
	}
	if len(builder.option.Config[0].Operator) != 1 {
		t.Errorf("AddStepConfig() Operator length = %v, want 1", len(builder.option.Config[0].Operator))
	}
	if builder.option.Config[0].Operator[0] != "operator1" {
		t.Errorf("AddStepConfig() Operator[0] = %v, want operator1", builder.option.Config[0].Operator[0])
	}
}

func TestTemplateBuilder_Build(t *testing.T) {
	tests := []struct {
		name        string
		uid         string
		startStep   string
		configs     []*models.StepConfig
		expectError bool
	}{
		{
			name:      "valid template",
			uid:       "template-001",
			startStep: "submit",
			configs: []*models.StepConfig{
				{Step: "submit", State: "pending", Next: []*models.NextStep{{Step: "review", Operation: "submit"}}, Disposal: models.Disposal{SignType: models.AnyoneSign}},
				{Step: "review", State: "reviewing", Next: []*models.NextStep{{Step: "approved", Operation: "approve"}, {Step: "rejected", Operation: "reject"}}, Disposal: models.Disposal{SignType: models.JointlySign, JointSignRate: 0.5}},
				{Step: "approved", State: "approved", Disposal: models.Disposal{SignType: models.AnyoneSign}},
				{Step: "rejected", State: "rejected", Disposal: models.Disposal{SignType: models.AnyoneSign}},
			},
			expectError: false,
		},
		{
			name:        "empty config should error",
			uid:         "template-001",
			startStep:   "submit",
			configs:     []*models.StepConfig{},
			expectError: true,
		},
		{
			name:        "nil config should error",
			uid:         "template-001",
			startStep:   "submit",
			configs:     nil,
			expectError: true,
		},
		{
			name:      "config with nil step config should error",
			uid:       "template-001",
			startStep: "submit",
			configs: []*models.StepConfig{
				nil,
			},
			expectError: true,
		},
		{
			name:      "config with empty step should error",
			uid:       "template-001",
			startStep: "submit",
			configs: []*models.StepConfig{
				{Step: "", State: "pending"},
			},
			expectError: true,
		},
		{
			name:      "config with empty state should error",
			uid:       "template-001",
			startStep: "submit",
			configs: []*models.StepConfig{
				{Step: "submit", State: ""},
			},
			expectError: true,
		},
		{
			name:      "start step not found in config should error",
			uid:       "template-001",
			startStep: "submit",
			configs: []*models.StepConfig{
				{Step: "review", State: "reviewing"},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewTemplateBuilder(tt.uid, tt.startStep)
			if tt.configs != nil {
				builder.SetConfig(tt.configs)
				// Set end steps for valid templates
				if !tt.expectError {
					builder.SetEndStep([]string{"approved", "rejected"})
				}
			}

			template, err := builder.Build()

			if tt.expectError {
				if err == nil {
					t.Errorf("Build() expected error, got nil")
				}
				if template != nil {
					t.Errorf("Build() expected nil template, got %v", template)
				}
			} else {
				if err != nil {
					t.Errorf("Build() unexpected error: %v", err)
				}
				if template == nil {
					t.Errorf("Build() expected template, got nil")
				}
				if template.Uid != tt.uid {
					t.Errorf("Build() Uid = %v, want %v", template.Uid, tt.uid)
				}
				if template.StartStep != tt.startStep {
					t.Errorf("Build() StartStep = %v, want %v", template.StartStep, tt.startStep)
				}
			}
		})
	}
}

func TestTemplateBuilder_BuildOrPanic(t *testing.T) {
	tests := []struct {
		name        string
		uid         string
		startStep   string
		configs     []*models.StepConfig
		expectPanic bool
	}{
		{
			name:      "valid template",
			uid:       "template-001",
			startStep: "submit",
			configs: []*models.StepConfig{
				{Step: "submit", State: "pending", Next: []*models.NextStep{{Step: "review", Operation: "submit"}}, Disposal: models.Disposal{SignType: models.AnyoneSign}},
				{Step: "review", State: "reviewing", Next: []*models.NextStep{{Step: "approved", Operation: "approve"}, {Step: "rejected", Operation: "reject"}}, Disposal: models.Disposal{SignType: models.JointlySign, JointSignRate: 0.5}},
				{Step: "approved", State: "approved", Disposal: models.Disposal{SignType: models.AnyoneSign}},
				{Step: "rejected", State: "rejected", Disposal: models.Disposal{SignType: models.AnyoneSign}},
			},
			expectPanic: false,
		},
		{
			name:        "empty config should panic",
			uid:         "template-001",
			startStep:   "submit",
			configs:     []*models.StepConfig{},
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewTemplateBuilder(tt.uid, tt.startStep)
			if tt.configs != nil {
				builder.SetConfig(tt.configs)
				// Set end steps for valid templates
				if !tt.expectPanic {
					builder.SetEndStep([]string{"approved", "rejected"})
				}
			}

			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("BuildOrPanic() expected panic, got none")
					}
				}()
			}

			template := builder.BuildOrPanic()

			if !tt.expectPanic {
				if template == nil {
					t.Errorf("BuildOrPanic() expected template, got nil")
				}
				if template.Uid != tt.uid {
					t.Errorf("BuildOrPanic() Uid = %v, want %v", template.Uid, tt.uid)
				}
				if template.StartStep != tt.startStep {
					t.Errorf("BuildOrPanic() StartStep = %v, want %v", template.StartStep, tt.startStep)
				}
			}
		})
	}
}

func TestTemplateBuilder_Chaining(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	// Test method chaining
	result := builder.
		SetEndStep([]string{"approved", "rejected"}).
		AddEndStep("cancelled").
		SetBuiltin(true).
		AddStepConfig("submit", "pending", []string{"user"}).
		AddStepConfig("review", "reviewing", []string{"manager"})

	if result != builder {
		t.Errorf("Method chaining should return builder instance")
	}

	// Verify all values were set correctly
	if len(builder.option.EndStep) != 3 {
		t.Errorf("EndStep length = %v, want 3", len(builder.option.EndStep))
	}
	if builder.option.Builtin != true {
		t.Errorf("Builtin = %v, want true", builder.option.Builtin)
	}
	if len(builder.option.Config) != 2 {
		t.Errorf("Config length = %v, want 2", len(builder.option.Config))
	}
}

func TestTemplateBuilder_DefaultValues(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	// Test default values
	if builder.option.Uid != "template-001" {
		t.Errorf("Uid = %v, want template-001", builder.option.Uid)
	}
	if builder.option.StartStep != "submit" {
		t.Errorf("StartStep = %v, want submit", builder.option.StartStep)
	}
	if len(builder.option.EndStep) != 0 {
		t.Errorf("Default EndStep length = %v, want 0", len(builder.option.EndStep))
	}
	if len(builder.option.Config) != 0 {
		t.Errorf("Default Config length = %v, want 0", len(builder.option.Config))
	}
	if builder.option.Builtin != false {
		t.Errorf("Default Builtin = %v, want false", builder.option.Builtin)
	}
}

func TestTemplateBuilder_EndStepValidation(t *testing.T) {
	builder := NewTemplateBuilder("template-001", "submit")

	// Add config with submit step and proper disposal
	builder.AddStepConfig("submit", "pending", []string{"user"})
	builder.AddStepConfig("review", "reviewing", []string{"manager"})

	// Set proper disposal for each step
	builder.option.Config[0].Disposal = models.Disposal{SignType: models.AnyoneSign}
	builder.option.Config[1].Disposal = models.Disposal{SignType: models.AnyoneSign}

	// Add next steps to make the flow reachable
	builder.option.Config[0].Next = []*models.NextStep{{Step: "review", Operation: "submit"}}

	// Test valid end steps
	builder.SetEndStep([]string{"review"})
	template, err := builder.Build()
	if err != nil {
		t.Errorf("Build() with valid end steps should not error: %v", err)
	}
	if template == nil {
		t.Errorf("Build() should return template")
	}

	// Test invalid end step
	builder.SetEndStep([]string{"invalid_step"})
	template, err = builder.Build()
	if err == nil {
		t.Errorf("Build() with invalid end step should error")
	}
	if template != nil {
		t.Errorf("Build() should return nil template")
	}
}
