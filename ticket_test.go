package punched_tape

import (
	"testing"

	"github.com/FatWang1/punched-tape/models"
)

func TestNewTicketBuilder(t *testing.T) {
	tests := []struct {
		name      string
		uid       string
		orderNum  string
		step      string
		expectNil bool
	}{
		{
			name:      "valid parameters",
			uid:       "user123",
			orderNum:  "TICKET-001",
			step:      "approval",
			expectNil: false,
		},
		{
			name:      "empty parameters",
			uid:       "",
			orderNum:  "",
			step:      "",
			expectNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewTicketBuilder(tt.uid, tt.orderNum, tt.step, tt.name)
			if tt.expectNil && builder != nil {
				t.Errorf("NewTicketBuilder() expected nil, got %v", builder)
			}
			if !tt.expectNil && builder == nil {
				t.Errorf("NewTicketBuilder() expected non-nil, got nil")
			}
		})
	}
}

func TestTicketBuilder_SetStatus(t *testing.T) {
	builder := NewTicketBuilder("user123", "TICKET-001", "approval", "test")

	tests := []struct {
		name     string
		status   string
		expected string
	}{
		{
			name:     "set running status",
			status:   models.Running,
			expected: models.Running,
		},
		{
			name:     "set passed status",
			status:   models.Passed,
			expected: models.Passed,
		},
		{
			name:     "set rejected status",
			status:   models.Rejected,
			expected: models.Rejected,
		},
		{
			name:     "set custom status",
			status:   "custom",
			expected: "custom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetStatus(tt.status)
			if result != builder {
				t.Errorf("SetStatus() should return builder instance")
			}
			if builder.option.Status != tt.expected {
				t.Errorf("SetStatus() = %v, want %v", builder.option.Status, tt.expected)
			}
		})
	}
}

func TestTicketBuilder_SetOperator(t *testing.T) {
	builder := NewTicketBuilder("user123", "TICKET-001", "approval", "test")

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

func TestTicketBuilder_AddOperator(t *testing.T) {
	builder := NewTicketBuilder("user123", "TICKET-001", "approval", "test")

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

func TestTicketBuilder_SetOperatedUser(t *testing.T) {
	builder := NewTicketBuilder("user123", "TICKET-001", "approval", "test")

	tests := []struct {
		name         string
		operatedUser []string
		expected     []string
	}{
		{
			name:         "set single operated user",
			operatedUser: []string{"user1"},
			expected:     []string{"user1"},
		},
		{
			name:         "set multiple operated users",
			operatedUser: []string{"user1", "user2", "user3"},
			expected:     []string{"user1", "user2", "user3"},
		},
		{
			name:         "set empty operated users",
			operatedUser: []string{},
			expected:     []string{},
		},
		{
			name:         "set nil operated users",
			operatedUser: nil,
			expected:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetOperatedUser(tt.operatedUser)
			if result != builder {
				t.Errorf("SetOperatedUser() should return builder instance")
			}
			if len(builder.option.OperatedUser) != len(tt.expected) {
				t.Errorf("SetOperatedUser() length = %v, want %v", len(builder.option.OperatedUser), len(tt.expected))
				return
			}
			for i, v := range builder.option.OperatedUser {
				if v != tt.expected[i] {
					t.Errorf("SetOperatedUser()[%d] = %v, want %v", i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestTicketBuilder_AddOperatedUser(t *testing.T) {
	builder := NewTicketBuilder("user123", "TICKET-001", "approval", "test")

	// Test initial state
	if len(builder.option.OperatedUser) != 0 {
		t.Errorf("Initial operated user length should be 0, got %v", len(builder.option.OperatedUser))
	}

	// Test adding single operated user
	result := builder.AddOperatedUser("user1")
	if result != builder {
		t.Errorf("AddOperatedUser() should return builder instance")
	}
	if len(builder.option.OperatedUser) != 1 {
		t.Errorf("AddOperatedUser() length = %v, want 1", len(builder.option.OperatedUser))
	}
	if builder.option.OperatedUser[0] != "user1" {
		t.Errorf("AddOperatedUser()[0] = %v, want user1", builder.option.OperatedUser[0])
	}

	// Test adding multiple operated users
	result = builder.AddOperatedUser("user2", "user3")
	if len(builder.option.OperatedUser) != 3 {
		t.Errorf("AddOperatedUser() length = %v, want 3", len(builder.option.OperatedUser))
	}
	expected := []string{"user1", "user2", "user3"}
	for i, v := range builder.option.OperatedUser {
		if v != expected[i] {
			t.Errorf("AddOperatedUser()[%d] = %v, want %v", i, v, expected[i])
		}
	}
}

func TestTicketBuilder_SetMemo(t *testing.T) {
	builder := NewTicketBuilder("user123", "TICKET-001", "approval", "test")

	tests := []struct {
		name     string
		memo     string
		expected string
	}{
		{
			name:     "set memo",
			memo:     "test memo",
			expected: "test memo",
		},
		{
			name:     "set empty memo",
			memo:     "",
			expected: "",
		},
		{
			name:     "set long memo",
			memo:     "this is a very long memo with lots of text",
			expected: "this is a very long memo with lots of text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := builder.SetMemo(tt.memo)
			if result != builder {
				t.Errorf("SetMemo() should return builder instance")
			}
			if builder.option.Memo != tt.expected {
				t.Errorf("SetMemo() = %v, want %v", builder.option.Memo, tt.expected)
			}
		})
	}
}

func TestTicketBuilder_Build(t *testing.T) {
	tests := []struct {
		name        string
		uid         string
		orderNum    string
		step        string
		status      string
		expectError bool
	}{
		{
			name:        "valid ticket",
			uid:         "user123",
			orderNum:    "TICKET-001",
			step:        "approval",
			status:      models.Running,
			expectError: false,
		},
		{
			name:        "valid ticket with passed status",
			uid:         "user123",
			orderNum:    "TICKET-001",
			step:        "approval",
			status:      models.Passed,
			expectError: false,
		},
		{
			name:        "valid ticket with rejected status",
			uid:         "user123",
			orderNum:    "TICKET-001",
			step:        "approval",
			status:      models.Rejected,
			expectError: false,
		},
		{
			name:        "invalid status",
			uid:         "user123",
			orderNum:    "TICKET-001",
			step:        "approval",
			status:      "invalid_status",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewTicketBuilder(tt.uid, tt.orderNum, tt.step, "test")
			builder.SetStatus(tt.status)

			ticket, err := builder.Build()

			if tt.expectError {
				if err == nil {
					t.Errorf("Build() expected error, got nil")
				}
				if ticket != nil {
					t.Errorf("Build() expected nil ticket, got %v", ticket)
				}
			} else {
				if err != nil {
					t.Errorf("Build() unexpected error: %v", err)
				}
				if ticket == nil {
					t.Errorf("Build() expected ticket, got nil")
				}
				if ticket.Uid != tt.uid {
					t.Errorf("Build() Uid = %v, want %v", ticket.Uid, tt.uid)
				}
				if ticket.OrderNum != tt.orderNum {
					t.Errorf("Build() OrderNum = %v, want %v", ticket.OrderNum, tt.orderNum)
				}
				if ticket.Step != tt.step {
					t.Errorf("Build() Step = %v, want %v", ticket.Step, tt.step)
				}
				if ticket.Status != tt.status {
					t.Errorf("Build() Status = %v, want %v", ticket.Status, tt.status)
				}
			}
		})
	}
}

func TestTicketBuilder_BuildOrPanic(t *testing.T) {
	tests := []struct {
		name        string
		uid         string
		orderNum    string
		step        string
		status      string
		expectPanic bool
	}{
		{
			name:        "valid ticket",
			uid:         "user123",
			orderNum:    "TICKET-001",
			step:        "approval",
			status:      models.Running,
			expectPanic: false,
		},
		{
			name:        "invalid status should panic",
			uid:         "user123",
			orderNum:    "TICKET-001",
			step:        "approval",
			status:      "invalid_status",
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewTicketBuilder(tt.uid, tt.orderNum, tt.step, "test")
			builder.SetStatus(tt.status)

			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("BuildOrPanic() expected panic, got none")
					}
				}()
			}

			ticket := builder.BuildOrPanic()

			if !tt.expectPanic {
				if ticket == nil {
					t.Errorf("BuildOrPanic() expected ticket, got nil")
				}
				if ticket.Uid != tt.uid {
					t.Errorf("BuildOrPanic() Uid = %v, want %v", ticket.Uid, tt.uid)
				}
				if ticket.OrderNum != tt.orderNum {
					t.Errorf("BuildOrPanic() OrderNum = %v, want %v", ticket.OrderNum, tt.orderNum)
				}
				if ticket.Step != tt.step {
					t.Errorf("BuildOrPanic() Step = %v, want %v", ticket.Step, tt.step)
				}
				if ticket.Status != tt.status {
					t.Errorf("BuildOrPanic() Status = %v, want %v", ticket.Status, tt.status)
				}
			}
		})
	}
}

func TestTicketBuilder_Chaining(t *testing.T) {
	builder := NewTicketBuilder("user123", "TICKET-001", "approval", "test")

	// Test method chaining
	result := builder.
		SetStatus(models.Passed).
		SetOperator([]string{"admin1", "admin2"}).
		AddOperator("admin3").
		SetOperatedUser([]string{"user1"}).
		AddOperatedUser("user2").
		SetMemo("test memo")

	if result != builder {
		t.Errorf("Method chaining should return builder instance")
	}

	// Verify all values were set correctly
	if builder.option.Status != models.Passed {
		t.Errorf("Status = %v, want %v", builder.option.Status, models.Passed)
	}
	if len(builder.option.Operator) != 3 {
		t.Errorf("Operator length = %v, want 3", len(builder.option.Operator))
	}
	if len(builder.option.OperatedUser) != 2 {
		t.Errorf("OperatedUser length = %v, want 2", len(builder.option.OperatedUser))
	}
	if builder.option.Memo != "test memo" {
		t.Errorf("Memo = %v, want test memo", builder.option.Memo)
	}
}

func TestTicketBuilder_DefaultValues(t *testing.T) {
	builder := NewTicketBuilder("user123", "TICKET-001", "approval", "test")

	// Test default values
	if builder.option.Status != models.Running {
		t.Errorf("Default Status = %v, want %v", builder.option.Status, models.Running)
	}
	if builder.option.Uid != "user123" {
		t.Errorf("Uid = %v, want user123", builder.option.Uid)
	}
	if builder.option.OrderNum != "TICKET-001" {
		t.Errorf("OrderNum = %v, want TICKET-001", builder.option.OrderNum)
	}
	if builder.option.Step != "approval" {
		t.Errorf("Step = %v, want approval", builder.option.Step)
	}
	if len(builder.option.Operator) != 0 {
		t.Errorf("Default Operator length = %v, want 0", len(builder.option.Operator))
	}
	if len(builder.option.OperatedUser) != 0 {
		t.Errorf("Default OperatedUser length = %v, want 0", len(builder.option.OperatedUser))
	}
	if builder.option.Memo != "" {
		t.Errorf("Default Memo = %v, want empty string", builder.option.Memo)
	}
}
