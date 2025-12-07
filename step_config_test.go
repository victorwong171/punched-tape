package punched_tape

import (
	"testing"

	"github.com/victorwong171/punched-tape/models"
)

// NextStepBuilder Tests
func TestNewNextStepBuilder(t *testing.T) {
	tests := []struct {
		name      string
		step      string
		operation string
		expectNil bool
	}{
		{
			name:      "valid parameters",
			step:      "next1",
			operation: "approve",
			expectNil: false,
		},
		{
			name:      "empty parameters",
			step:      "",
			operation: "",
			expectNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewNextStepBuilder(tt.step, tt.operation)
			if tt.expectNil && builder != nil {
				t.Errorf("NewNextStepBuilder() expected nil, got %v", builder)
			}
			if !tt.expectNil && builder == nil {
				t.Errorf("NewNextStepBuilder() expected non-nil, got nil")
			}
		})
	}
}

func TestNextStepBuilder_Build(t *testing.T) {
	tests := []struct {
		name        string
		step        string
		operation   string
		expectError bool
	}{
		{
			name:        "valid next step",
			step:        "next1",
			operation:   "approve",
			expectError: false,
		},
		{
			name:        "empty step",
			step:        "",
			operation:   "approve",
			expectError: true,
		},
		{
			name:        "empty operation",
			step:        "next1",
			operation:   "",
			expectError: true,
		},
		{
			name:        "both empty",
			step:        "",
			operation:   "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewNextStepBuilder(tt.step, tt.operation)

			nextStep, err := builder.Build()

			if tt.expectError {
				if err == nil {
					t.Errorf("Build() expected error, got nil")
				}
				if nextStep != nil {
					t.Errorf("Build() expected nil nextStep, got %v", nextStep)
				}
			} else {
				if err != nil {
					t.Errorf("Build() unexpected error: %v", err)
				}
				if nextStep == nil {
					t.Errorf("Build() expected nextStep, got nil")
				}
				if nextStep.Step != tt.step {
					t.Errorf("Build() Step = %v, want %v", nextStep.Step, tt.step)
				}
				if nextStep.Operation != tt.operation {
					t.Errorf("Build() Operation = %v, want %v", nextStep.Operation, tt.operation)
				}
			}
		})
	}
}

func TestNextStepBuilder_BuildOrPanic(t *testing.T) {
	tests := []struct {
		name        string
		step        string
		operation   string
		expectPanic bool
	}{
		{
			name:        "valid next step",
			step:        "next1",
			operation:   "approve",
			expectPanic: false,
		},
		{
			name:        "empty step should panic",
			step:        "",
			operation:   "approve",
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewNextStepBuilder(tt.step, tt.operation)

			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("BuildOrPanic() expected panic, got none")
					}
				}()
			}

			nextStep := builder.BuildOrPanic()

			if !tt.expectPanic {
				if nextStep == nil {
					t.Errorf("BuildOrPanic() expected nextStep, got nil")
				}
				if nextStep.Step != tt.step {
					t.Errorf("BuildOrPanic() Step = %v, want %v", nextStep.Step, tt.step)
				}
				if nextStep.Operation != tt.operation {
					t.Errorf("BuildOrPanic() Operation = %v, want %v", nextStep.Operation, tt.operation)
				}
			}
		})
	}
}

// DisposalBuilder Tests
func TestNewDisposalBuilder(t *testing.T) {
	builder := NewDisposalBuilder()
	if builder == nil {
		t.Errorf("NewDisposalBuilder() returned nil")
	}
	if builder.option.JointSignRate != 0.0 {
		t.Errorf("Default JointSignRate = %v, want 0.0", builder.option.JointSignRate)
	}
}

func TestDisposalBuilder_SetSignType(t *testing.T) {
	builder := NewDisposalBuilder()

	tests := []struct {
		name     string
		signType string
		expected string
	}{
		{
			name:     "set jointly sign",
			signType: models.JointlySign,
			expected: models.JointlySign,
		},
		{
			name:     "set serial sign",
			signType: models.SerialSign,
			expected: models.SerialSign,
		},
		{
			name:     "set anyone sign",
			signType: models.AnyoneSign,
			expected: models.AnyoneSign,
		},
		{
			name:     "set custom sign type",
			signType: "custom",
			expected: "custom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetSignType(tt.signType)
			if result != builder {
				t.Errorf("SetSignType() should return builder instance")
			}
			if builder.option.SignType != tt.expected {
				t.Errorf("SetSignType() = %v, want %v", builder.option.SignType, tt.expected)
			}
		})
	}
}

func TestDisposalBuilder_SetJointSignRate(t *testing.T) {
	builder := NewDisposalBuilder()

	tests := []struct {
		name     string
		rate     float32
		expected float32
	}{
		{
			name:     "set valid rate",
			rate:     0.5,
			expected: 0.5,
		},
		{
			name:     "set zero rate",
			rate:     0.0,
			expected: 0.0,
		},
		{
			name:     "set one rate",
			rate:     1.0,
			expected: 1.0,
		},
		{
			name:     "set negative rate",
			rate:     -0.1,
			expected: -0.1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetJointSignRate(tt.rate)
			if result != builder {
				t.Errorf("SetJointSignRate() should return builder instance")
			}
			if builder.option.JointSignRate != tt.expected {
				t.Errorf("SetJointSignRate() = %v, want %v", builder.option.JointSignRate, tt.expected)
			}
		})
	}
}

func TestDisposalBuilder_Build(t *testing.T) {
	tests := []struct {
		name        string
		signType    string
		rate        float32
		expectError bool
	}{
		{
			name:        "valid disposal",
			signType:    models.JointlySign,
			rate:        0.5,
			expectError: false,
		},
		{
			name:        "valid disposal with serial sign",
			signType:    models.SerialSign,
			rate:        0.0,
			expectError: false,
		},
		{
			name:        "valid disposal with anyone sign",
			signType:    models.AnyoneSign,
			rate:        0.0,
			expectError: false,
		},
		{
			name:        "invalid sign type",
			signType:    "invalid",
			rate:        0.0,
			expectError: true,
		},
		{
			name:        "jointly sign with invalid rate",
			signType:    models.JointlySign,
			rate:        -0.1,
			expectError: true,
		},
		{
			name:        "jointly sign with rate > 1",
			signType:    models.JointlySign,
			rate:        1.1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewDisposalBuilder()
			builder.SetSignType(tt.signType)
			builder.SetJointSignRate(tt.rate)

			disposal, err := builder.Build()

			if tt.expectError {
				if err == nil {
					t.Errorf("Build() expected error, got nil")
				}
				if disposal != nil {
					t.Errorf("Build() expected nil disposal, got %v", disposal)
				}
			} else {
				if err != nil {
					t.Errorf("Build() unexpected error: %v", err)
				}
				if disposal == nil {
					t.Errorf("Build() expected disposal, got nil")
				}
				if disposal.SignType != tt.signType {
					t.Errorf("Build() SignType = %v, want %v", disposal.SignType, tt.signType)
				}
				if disposal.JointSignRate != tt.rate {
					t.Errorf("Build() JointSignRate = %v, want %v", disposal.JointSignRate, tt.rate)
				}
			}
		})
	}
}

func TestDisposalBuilder_BuildOrPanic(t *testing.T) {
	tests := []struct {
		name        string
		signType    string
		rate        float32
		expectPanic bool
	}{
		{
			name:        "valid disposal",
			signType:    models.JointlySign,
			rate:        0.5,
			expectPanic: false,
		},
		{
			name:        "invalid sign type should panic",
			signType:    "invalid",
			rate:        0.0,
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewDisposalBuilder()
			builder.SetSignType(tt.signType)
			builder.SetJointSignRate(tt.rate)

			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("BuildOrPanic() expected panic, got none")
					}
				}()
			}

			disposal := builder.BuildOrPanic()

			if !tt.expectPanic {
				if disposal == nil {
					t.Errorf("BuildOrPanic() expected disposal, got nil")
				}
				if disposal.SignType != tt.signType {
					t.Errorf("BuildOrPanic() SignType = %v, want %v", disposal.SignType, tt.signType)
				}
				if disposal.JointSignRate != tt.rate {
					t.Errorf("BuildOrPanic() JointSignRate = %v, want %v", disposal.JointSignRate, tt.rate)
				}
			}
		})
	}
}

// StepConfigBuilder Tests
func TestNewStepConfigBuilder(t *testing.T) {
	tests := []struct {
		name      string
		step      string
		state     string
		expectNil bool
	}{
		{
			name:      "valid parameters",
			step:      "approval",
			state:     "pending",
			expectNil: false,
		},
		{
			name:      "empty parameters",
			step:      "",
			state:     "",
			expectNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewStepConfigBuilder(tt.step, tt.state)
			if tt.expectNil && builder != nil {
				t.Errorf("NewStepConfigBuilder() expected nil, got %v", builder)
			}
			if !tt.expectNil && builder == nil {
				t.Errorf("NewStepConfigBuilder() expected non-nil, got nil")
			}
		})
	}
}

func TestStepConfigBuilder_SetOperator(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	tests := []struct {
		name     string
		operator []string
		expected []string
	}{
		{
			name:     "set single operator",
			operator: []string{"admin1"},
			expected: []string{"admin1"},
		},
		{
			name:     "set multiple operators",
			operator: []string{"admin1", "admin2", "admin3"},
			expected: []string{"admin1", "admin2", "admin3"},
		},
		{
			name:     "set empty operators",
			operator: []string{},
			expected: []string{},
		},
		{
			name:     "set nil operators",
			operator: nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetOperator(tt.operator)
			if result != builder {
				t.Errorf("SetOperator() should return builder instance")
			}
			if len(builder.option.Operator) != len(tt.expected) {
				t.Errorf("SetOperator() length = %v, want %v", len(builder.option.Operator), len(tt.expected))
				return
			}
			for i, v := range builder.option.Operator {
				if v != tt.expected[i] {
					t.Errorf("SetOperator()[%d] = %v, want %v", i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestStepConfigBuilder_AddOperator(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	// Test initial state
	if len(builder.option.Operator) != 0 {
		t.Errorf("Initial operator length should be 0, got %v", len(builder.option.Operator))
	}

	// Test adding single operator
	result := builder.AddOperator("admin1")
	if result != builder {
		t.Errorf("AddOperator() should return builder instance")
	}
	if len(builder.option.Operator) != 1 {
		t.Errorf("AddOperator() length = %v, want 1", len(builder.option.Operator))
	}
	if builder.option.Operator[0] != "admin1" {
		t.Errorf("AddOperator()[0] = %v, want admin1", builder.option.Operator[0])
	}

	// Test adding multiple operators
	result = builder.AddOperator("admin2", "admin3")
	if len(builder.option.Operator) != 3 {
		t.Errorf("AddOperator() length = %v, want 3", len(builder.option.Operator))
	}
	expected := []string{"admin1", "admin2", "admin3"}
	for i, v := range builder.option.Operator {
		if v != expected[i] {
			t.Errorf("AddOperator()[%d] = %v, want %v", i, v, expected[i])
		}
	}
}

func TestStepConfigBuilder_SetNext(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	nextSteps := []*models.NextStep{
		{Step: "next1", Operation: "op1"},
		{Step: "next2", Operation: "op2"},
	}

	result := builder.SetNext(nextSteps)
	if result != builder {
		t.Errorf("SetNext() should return builder instance")
	}
	if len(builder.option.Next) != len(nextSteps) {
		t.Errorf("SetNext() length = %v, want %v", len(builder.option.Next), len(nextSteps))
	}
}

func TestStepConfigBuilder_AddNext(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	// Test initial state
	if len(builder.option.Next) != 0 {
		t.Errorf("Initial next length should be 0, got %v", len(builder.option.Next))
	}

	// Test adding single next step
	nextStep1 := &models.NextStep{Step: "next1", Operation: "op1"}
	result := builder.AddNext(nextStep1)
	if result != builder {
		t.Errorf("AddNext() should return builder instance")
	}
	if len(builder.option.Next) != 1 {
		t.Errorf("AddNext() length = %v, want 1", len(builder.option.Next))
	}

	// Test adding multiple next steps
	nextStep2 := &models.NextStep{Step: "next2", Operation: "op2"}
	nextStep3 := &models.NextStep{Step: "next3", Operation: "op3"}
	result = builder.AddNext(nextStep2, nextStep3)
	if len(builder.option.Next) != 3 {
		t.Errorf("AddNext() length = %v, want 3", len(builder.option.Next))
	}
}

func TestStepConfigBuilder_AddNextStep(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	// Test initial state
	if len(builder.option.Next) != 0 {
		t.Errorf("Initial next length should be 0, got %v", len(builder.option.Next))
	}

	// Test adding next step
	result := builder.AddNextStep("next1", "op1")
	if result != builder {
		t.Errorf("AddNextStep() should return builder instance")
	}
	if len(builder.option.Next) != 1 {
		t.Errorf("AddNextStep() length = %v, want 1", len(builder.option.Next))
	}
	if builder.option.Next[0].Step != "next1" {
		t.Errorf("AddNextStep() Step = %v, want next1", builder.option.Next[0].Step)
	}
	if builder.option.Next[0].Operation != "op1" {
		t.Errorf("AddNextStep() Operation = %v, want op1", builder.option.Next[0].Operation)
	}
}

func TestStepConfigBuilder_SetDisposal(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	disposal := models.Disposal{
		SignType:      models.JointlySign,
		JointSignRate: 0.5,
	}

	result := builder.SetDisposal(disposal)
	if result != builder {
		t.Errorf("SetDisposal() should return builder instance")
	}
	if builder.option.Disposal.SignType != disposal.SignType {
		t.Errorf("SetDisposal() SignType = %v, want %v", builder.option.Disposal.SignType, disposal.SignType)
	}
	if builder.option.Disposal.JointSignRate != disposal.JointSignRate {
		t.Errorf("SetDisposal() JointSignRate = %v, want %v", builder.option.Disposal.JointSignRate, disposal.JointSignRate)
	}
}

func TestStepConfigBuilder_SetDisposalSignType(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	tests := []struct {
		name     string
		signType string
		expected string
	}{
		{
			name:     "set jointly sign",
			signType: models.JointlySign,
			expected: models.JointlySign,
		},
		{
			name:     "set serial sign",
			signType: models.SerialSign,
			expected: models.SerialSign,
		},
		{
			name:     "set anyone sign",
			signType: models.AnyoneSign,
			expected: models.AnyoneSign,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetDisposalSignType(tt.signType)
			if result != builder {
				t.Errorf("SetDisposalSignType() should return builder instance")
			}
			if builder.option.Disposal.SignType != tt.expected {
				t.Errorf("SetDisposalSignType() = %v, want %v", builder.option.Disposal.SignType, tt.expected)
			}
		})
	}
}

func TestStepConfigBuilder_SetDisposalJointSignRate(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	tests := []struct {
		name     string
		rate     float32
		expected float32
	}{
		{
			name:     "set valid rate",
			rate:     0.5,
			expected: 0.5,
		},
		{
			name:     "set zero rate",
			rate:     0.0,
			expected: 0.0,
		},
		{
			name:     "set one rate",
			rate:     1.0,
			expected: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetDisposalJointSignRate(tt.rate)
			if result != builder {
				t.Errorf("SetDisposalJointSignRate() should return builder instance")
			}
			if builder.option.Disposal.JointSignRate != tt.expected {
				t.Errorf("SetDisposalJointSignRate() = %v, want %v", builder.option.Disposal.JointSignRate, tt.expected)
			}
		})
	}
}

func TestStepConfigBuilder_Build(t *testing.T) {
	tests := []struct {
		name        string
		step        string
		state       string
		signType    string
		rate        float32
		expectError bool
	}{
		{
			name:        "valid step config",
			step:        "approval",
			state:       "pending",
			signType:    models.JointlySign,
			rate:        0.5,
			expectError: false,
		},
		{
			name:        "valid step config with serial sign",
			step:        "approval",
			state:       "pending",
			signType:    models.SerialSign,
			rate:        0.0,
			expectError: false,
		},
		{
			name:        "valid step config with anyone sign",
			step:        "approval",
			state:       "pending",
			signType:    models.AnyoneSign,
			rate:        0.0,
			expectError: false,
		},
		{
			name:        "invalid disposal sign type",
			step:        "approval",
			state:       "pending",
			signType:    "invalid",
			rate:        0.0,
			expectError: true,
		},
		{
			name:        "jointly sign with invalid rate",
			step:        "approval",
			state:       "pending",
			signType:    models.JointlySign,
			rate:        -0.1,
			expectError: true,
		},
		{
			name:        "jointly sign with rate > 1",
			step:        "approval",
			state:       "pending",
			signType:    models.JointlySign,
			rate:        1.1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewStepConfigBuilder(tt.step, tt.state)
			builder.SetDisposalSignType(tt.signType)
			builder.SetDisposalJointSignRate(tt.rate)

			config, err := builder.Build()

			if tt.expectError {
				if err == nil {
					t.Errorf("Build() expected error, got nil")
				}
				if config != nil {
					t.Errorf("Build() expected nil config, got %v", config)
				}
			} else {
				if err != nil {
					t.Errorf("Build() unexpected error: %v", err)
				}
				if config == nil {
					t.Errorf("Build() expected config, got nil")
				}
				if config.Step != tt.step {
					t.Errorf("Build() Step = %v, want %v", config.Step, tt.step)
				}
				if config.State != tt.state {
					t.Errorf("Build() State = %v, want %v", config.State, tt.state)
				}
				if config.Disposal.SignType != tt.signType {
					t.Errorf("Build() Disposal.SignType = %v, want %v", config.Disposal.SignType, tt.signType)
				}
				if config.Disposal.JointSignRate != tt.rate {
					t.Errorf("Build() Disposal.JointSignRate = %v, want %v", config.Disposal.JointSignRate, tt.rate)
				}
			}
		})
	}
}

func TestStepConfigBuilder_BuildOrPanic(t *testing.T) {
	tests := []struct {
		name        string
		step        string
		state       string
		signType    string
		rate        float32
		expectPanic bool
	}{
		{
			name:        "valid step config",
			step:        "approval",
			state:       "pending",
			signType:    models.JointlySign,
			rate:        0.5,
			expectPanic: false,
		},
		{
			name:        "invalid disposal sign type should panic",
			step:        "approval",
			state:       "pending",
			signType:    "invalid",
			rate:        0.0,
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewStepConfigBuilder(tt.step, tt.state)
			builder.SetDisposalSignType(tt.signType)
			builder.SetDisposalJointSignRate(tt.rate)

			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("BuildOrPanic() expected panic, got none")
					}
				}()
			}

			config := builder.BuildOrPanic()

			if !tt.expectPanic {
				if config == nil {
					t.Errorf("BuildOrPanic() expected config, got nil")
				}
				if config.Step != tt.step {
					t.Errorf("BuildOrPanic() Step = %v, want %v", config.Step, tt.step)
				}
				if config.State != tt.state {
					t.Errorf("BuildOrPanic() State = %v, want %v", config.State, tt.state)
				}
				if config.Disposal.SignType != tt.signType {
					t.Errorf("BuildOrPanic() Disposal.SignType = %v, want %v", config.Disposal.SignType, tt.signType)
				}
				if config.Disposal.JointSignRate != tt.rate {
					t.Errorf("BuildOrPanic() Disposal.JointSignRate = %v, want %v", config.Disposal.JointSignRate, tt.rate)
				}
			}
		})
	}
}

func TestStepConfigBuilder_Chaining(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	// Test method chaining
	result := builder.
		SetOperator([]string{"admin1", "admin2"}).
		AddOperator("admin3").
		AddNextStep("next1", "op1").
		SetDisposalSignType(models.JointlySign).
		SetDisposalJointSignRate(0.5)

	if result != builder {
		t.Errorf("Method chaining should return builder instance")
	}

	// Verify all values were set correctly
	if len(builder.option.Operator) != 3 {
		t.Errorf("Operator length = %v, want 3", len(builder.option.Operator))
	}
	if len(builder.option.Next) != 1 {
		t.Errorf("Next length = %v, want 1", len(builder.option.Next))
	}
	if builder.option.Disposal.SignType != models.JointlySign {
		t.Errorf("Disposal.SignType = %v, want %v", builder.option.Disposal.SignType, models.JointlySign)
	}
	if builder.option.Disposal.JointSignRate != 0.5 {
		t.Errorf("Disposal.JointSignRate = %v, want 0.5", builder.option.Disposal.JointSignRate)
	}
}

func TestStepConfigBuilder_DefaultValues(t *testing.T) {
	builder := NewStepConfigBuilder("approval", "pending")

	// Test default values
	if builder.option.Step != "approval" {
		t.Errorf("Step = %v, want approval", builder.option.Step)
	}
	if builder.option.State != "pending" {
		t.Errorf("State = %v, want pending", builder.option.State)
	}
	if len(builder.option.Operator) != 0 {
		t.Errorf("Default Operator length = %v, want 0", len(builder.option.Operator))
	}
	if len(builder.option.Next) != 0 {
		t.Errorf("Default Next length = %v, want 0", len(builder.option.Next))
	}
	if builder.option.Disposal.SignType != "" {
		t.Errorf("Default Disposal.SignType = %v, want empty string", builder.option.Disposal.SignType)
	}
	if builder.option.Disposal.JointSignRate != 0.0 {
		t.Errorf("Default Disposal.JointSignRate = %v, want 0.0", builder.option.Disposal.JointSignRate)
	}
}
