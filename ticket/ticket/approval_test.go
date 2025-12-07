package ticket

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/victorwong171/punched-tape/models"
)

func TestHelper_Approval(t *testing.T) {
	type args struct {
		next       string
		operation  string
		operator   string
		admin      bool
		endStep    []string
		ticket     *models.Ticket
		stepConfig map[string]*models.StepConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Ticket
		wantErr bool
	}{
		{
			name: "all is ok",
			args: args{
				next:      "b",
				operation: "submit",
				operator:  "user",
				admin:     false,
				endStep:   []string{"end"},
				ticket: &models.Ticket{
					Status:       models.Running,
					Step:         "a",
					Operator:     []string{"user"},
					OperatedUser: nil,
				},
				stepConfig: map[string]*models.StepConfig{
					"a": {
						Disposal: models.Disposal{
							SignType:      models.AnyoneSign,
							JointSignRate: 0,
						},
						Next: []*models.NextStep{
							{
								Operation: "submit",
								Step:      "b",
							},
						},
					},
					"b": {
						Operator: []string{"user"},
						Next: []*models.NextStep{
							{
								Operation: "pass",
								Step:      "end",
							},
						},
					},
				},
			},
			want: &models.Ticket{
				Status:   models.Running,
				Operator: []string{"user"},
			},
			wantErr: false,
		},
		{
			name: "unreachable",
			args: args{
				next:      "b",
				operation: "submit",
				operator:  "user",
				admin:     false,
				endStep:   []string{"end"},
				ticket: &models.Ticket{
					Status:       models.Running,
					Step:         "a",
					Operator:     []string{"user"},
					OperatedUser: nil,
				},
				stepConfig: map[string]*models.StepConfig{
					"a": {
						Disposal: models.Disposal{
							SignType:      models.AnyoneSign,
							JointSignRate: 0,
						},
						Next: []*models.NextStep{
							{
								Operation: "approval",
								Step:      "b",
							},
						},
					},
					"b": {
						Operator: []string{"user"},
						Next: []*models.NextStep{
							{
								Operation: "pass",
								Step:      "end",
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "cannot operate",
			args: args{
				next:      "b",
				operation: "submit",
				operator:  "user a",
				admin:     false,
				endStep:   []string{"end"},
				ticket: &models.Ticket{
					Status:       models.Running,
					Step:         "a",
					Operator:     []string{"user"},
					OperatedUser: nil,
				},
				stepConfig: map[string]*models.StepConfig{
					"a": {
						Disposal: models.Disposal{
							SignType:      models.AnyoneSign,
							JointSignRate: 0,
						},
						Next: []*models.NextStep{
							{
								Operation: "submit",
								Step:      "b",
							},
						},
					},
					"b": {
						Operator: []string{"user"},
						Next: []*models.NextStep{
							{
								Operation: "pass",
								Step:      "end",
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "already approved",
			args: args{
				next:      "b",
				operation: "submit",
				operator:  "user",
				admin:     false,
				endStep:   []string{"end"},
				ticket: &models.Ticket{
					Status:       models.Running,
					Step:         "a",
					Operator:     nil,
					OperatedUser: []string{"user"},
				},
				stepConfig: map[string]*models.StepConfig{
					"a": {
						Disposal: models.Disposal{
							SignType:      models.AnyoneSign,
							JointSignRate: 0,
						},
						Next: []*models.NextStep{
							{
								Operation: "submit",
								Step:      "b",
							},
						},
					},
					"b": {
						Operator: []string{"user"},
						Next: []*models.NextStep{
							{
								Operation: "pass",
								Step:      "end",
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong step",
			args: args{
				next:      "b",
				operation: "submit",
				operator:  "user a",
				admin:     false,
				endStep:   []string{"end"},
				ticket: &models.Ticket{
					Status:       models.Running,
					Step:         "c",
					Operator:     []string{"user"},
					OperatedUser: nil,
				},
				stepConfig: map[string]*models.StepConfig{
					"a": {
						Disposal: models.Disposal{
							SignType:      models.AnyoneSign,
							JointSignRate: 0,
						},
						Next: []*models.NextStep{
							{
								Operation: "submit",
								Step:      "b",
							},
						},
					},
					"b": {
						Operator: []string{"user"},
						Next: []*models.NextStep{
							{
								Operation: "pass",
								Step:      "end",
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ticket finished",
			args: args{
				next:      "b",
				operation: "submit",
				operator:  "user a",
				admin:     false,
				endStep:   []string{"end"},
				ticket: &models.Ticket{
					Status:       models.Passed,
					Step:         "a",
					Operator:     []string{"user"},
					OperatedUser: nil,
				},
				stepConfig: map[string]*models.StepConfig{
					"a": {
						Disposal: models.Disposal{
							SignType:      models.AnyoneSign,
							JointSignRate: 0,
						},
						Next: []*models.NextStep{
							{
								Operation: "submit",
								Step:      "b",
							},
						},
					},
					"b": {
						Operator: []string{"user"},
						Next: []*models.NextStep{
							{
								Operation: "pass",
								Step:      "end",
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "bad arguments",
			args: args{
				operation: "submit",
				operator:  "user a",
				admin:     false,
				endStep:   []string{"end"},
				ticket: &models.Ticket{
					Status:       models.Running,
					Step:         "a",
					Operator:     []string{"user"},
					OperatedUser: nil,
				},
				stepConfig: map[string]*models.StepConfig{
					"a": {
						Disposal: models.Disposal{
							SignType:      models.AnyoneSign,
							JointSignRate: 0,
						},
						Next: []*models.NextStep{
							{
								Operation: "submit",
								Step:      "b",
							},
						},
					},
					"b": {
						Operator: []string{"user"},
						Next: []*models.NextStep{
							{
								Operation: "pass",
								Step:      "end",
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Helper{}
			got, err := h.Approval(tt.args.next, tt.args.operation, tt.args.operator, tt.args.admin, tt.args.endStep, tt.args.ticket, tt.args.stepConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("Approval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); len(diff) > 0 {
				t.Errorf("Approval() diff = %v", diff)
			}
		})
	}
}

func Test_jointlySignUpdater(t *testing.T) {
	type args struct {
		operator        string
		ticket          *models.Ticket
		jointlySignRate float32
		nextStep        *models.StepConfig
		endStep         []string
	}
	tests := []struct {
		name string
		args args
		want *models.Ticket
	}{
		{
			name: "all is ok",
			args: args{
				operator: "a",
				ticket: &models.Ticket{
					Operator:     []string{"a", "b"},
					OperatedUser: []string{"c"},
				},
				jointlySignRate: 0.60,
				nextStep: &models.StepConfig{
					Step: "end",
				},
				endStep: []string{"end"},
			},
			want: &models.Ticket{
				Status: models.Passed,
				Step:   "end",
			},
		},
		{
			name: "spin operate",
			args: args{
				operator: "b",
				ticket: &models.Ticket{
					Operator:     []string{"a", "b"},
					OperatedUser: []string{"c"},
				},
				jointlySignRate: 0.70,
				nextStep: &models.StepConfig{
					Step: "end",
				},
				endStep: []string{"end"},
			},
			want: &models.Ticket{
				Operator:     []string{"a"},
				OperatedUser: []string{"c", "b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jointlySignUpdater(tt.args.operator, tt.args.ticket, tt.args.jointlySignRate, tt.args.nextStep, tt.args.endStep)
			if diff := cmp.Diff(got, tt.want); len(diff) > 0 {
				t.Errorf("jointlySignUpdater() diff = %v", diff)
			}
		})
	}
}

func Test_serialSignUpdater(t *testing.T) {
	type args struct {
		operator string
		ticket   *models.Ticket
		in2      float32
		nextStep *models.StepConfig
		endStep  []string
	}
	tests := []struct {
		name string
		args args
		want *models.Ticket
	}{
		{
			name: "all is ok",
			args: args{
				operator: "a",
				ticket: &models.Ticket{
					Operator:     []string{"a"},
					OperatedUser: []string{"b"},
				},
				nextStep: &models.StepConfig{
					Step: "next",
				},
			},
			want: &models.Ticket{
				Step: "next",
			},
		},
		{
			name: "spin approval",
			args: args{
				operator: "a",
				ticket: &models.Ticket{
					Operator:     []string{"a", "c"},
					OperatedUser: []string{"b"},
				},
			},
			want: &models.Ticket{
				Operator:     []string{"c"},
				OperatedUser: []string{"b", "a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := serialSignUpdater(tt.args.operator, tt.args.ticket, tt.args.in2, tt.args.nextStep, tt.args.endStep)
			if diff := cmp.Diff(got, tt.want); len(diff) > 0 {
				t.Errorf("serialSignUpdater() diff = %v", diff)
			}
		})
	}
}
