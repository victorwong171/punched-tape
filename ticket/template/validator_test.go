package template

import (
	"reflect"
	"testing"

	"github.com/victorwong171/punched-tape/models"

	"github.com/victorwong171/go-utils/desc/set"
)

func Test_canValidateReachability(t *testing.T) {
	type args struct {
		start      string
		stepMap    map[string]*models.StepConfig
		endStepSet set.Set[string]
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "all is ok",
			args: args{
				start: "1",
				stepMap: map[string]*models.StepConfig{
					"1": {
						Step: "1",
						Next: []*models.NextStep{
							{
								Step: "2",
							},
							{
								Step: "3",
							},
						},
					},
					"2": {
						Step: "2",
						Next: []*models.NextStep{
							{
								Step: "3",
							},
						},
					},
					"3": {
						Step: "3",
						Next: []*models.NextStep{
							{
								Step: "1",
							},
							{
								Step: "4",
							},
						},
					},
					"4": {
						Step: "4",
					},
				},
				endStepSet: set.Setify("4"),
			},
			wantErr: false,
		},
		{
			name: "bad next step",
			args: args{
				start: "1",
				stepMap: map[string]*models.StepConfig{
					"1": {
						Step: "1",
						Next: []*models.NextStep{
							{
								Step: "2",
							},
							{
								Step: "",
							},
						},
					},
					"2": {
						Step: "2",
						Next: []*models.NextStep{
							{
								Step: "3",
							},
						},
					},
					"3": {
						Step: "3",
						Next: []*models.NextStep{
							{
								Step: "1",
							},
							{
								Step: "4",
							},
						},
					},
					"4": {
						Step: "4",
					},
				},
				endStepSet: set.Setify("4"),
			},
			wantErr: true,
		},
		{
			name: "unreachable steps",
			args: args{
				start: "1",
				stepMap: map[string]*models.StepConfig{
					"1": {
						Step: "1",
						Next: []*models.NextStep{
							{
								Step: "2",
							},
							{
								Step: "3",
							},
						},
					},
					"2": {
						Step: "2",
						Next: []*models.NextStep{
							{
								Step: "3",
							},
						},
					},
					"3": {
						Step: "3",
						Next: []*models.NextStep{
							{
								Step: "1",
							},
							{
								Step: "4",
							},
						},
					},
					"4": {
						Step: "4",
					},
					"5": {
						Step: "5",
					},
				},
				endStepSet: set.Setify("4"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateReachability(tt.args.start, tt.args.stepMap, tt.args.endStepSet); (err != nil) != tt.wantErr {
				t.Errorf("canTraverse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validator_Validate(t *testing.T) {
	tests := []struct {
		name        string
		template    models.TicketTemplate
		signTypeSet set.Set[string]
		wantErr     error
	}{
		{
			name: "StartStep is empty",
			template: models.TicketTemplate{
				StartStep: "",
				Config:    []*models.StepConfig{},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrStartStepEmpty,
		},
		{
			name: "Config is empty",
			template: models.TicketTemplate{
				StartStep: "start",
				Config:    []*models.StepConfig{},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrConfigEmpty,
		},
		{
			name: "Bad step config",
			template: models.TicketTemplate{
				StartStep: "start",
				Config: []*models.StepConfig{
					{Step: "", Next: nil},
				},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrBadStepConfig,
		},
		{
			name: "Non-end step with no next steps",
			template: models.TicketTemplate{
				StartStep: "start",
				Config: []*models.StepConfig{
					{Step: "start", Next: nil},
				},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrNextStepEmpty,
		},
		{
			name: "Bad sign type",
			template: models.TicketTemplate{
				StartStep: "start",
				Config: []*models.StepConfig{
					{
						Step: "start",
						Disposal: models.Disposal{
							SignType: "invalidSignType",
						},
						Next: []*models.NextStep{
							{
								Step: "end",
							},
						},
					},
				},
			},
			signTypeSet: set.Setify("validSignType"),
			wantErr:     ErrBadSignType,
		},
		{
			name: "End step has next steps",
			template: models.TicketTemplate{
				StartStep: "start",
				EndStep:   []string{"end"},
				Config: []*models.StepConfig{
					{Step: "start", Next: []*models.NextStep{{Step: "end"}}},
					{Step: "end", Next: []*models.NextStep{{Step: "another"}}},
				},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrEndStepHasNext,
		},
		{
			name: "Duplicate step definition",
			template: models.TicketTemplate{
				StartStep: "start",
				EndStep:   []string{"next"},
				Config: []*models.StepConfig{
					{Step: "start", Next: []*models.NextStep{{Step: "next"}}},
					{Step: "next", Next: nil},
					{Step: "next", Next: nil}, // Duplicate
				},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrDuplicateStep,
		},
		{
			name: "Start step not found in configurations",
			template: models.TicketTemplate{
				StartStep: "notfound",
				EndStep:   []string{"start"},
				Config: []*models.StepConfig{
					{Step: "start", Next: nil},
				},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrStartStepNotFound,
		},
		{
			name: "unreachable",
			template: models.TicketTemplate{
				StartStep: "start",
				EndStep:   []string{"end"},
				Config: []*models.StepConfig{
					{
						Step: "start",
						Next: []*models.NextStep{
							{
								Step: "next",
							},
							{
								Step: "unreachable",
							},
						},
					},
					{
						Step: "next",
						Next: []*models.NextStep{
							{
								Step: "end",
							},
						},
					},
					{
						Step: "end",
						Next: nil,
					},
				},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrUnreachableSteps,
		},
		{
			name: "Some steps are unreachable",
			template: models.TicketTemplate{
				StartStep: "start",
				Config: []*models.StepConfig{
					{Step: "start", Next: []*models.NextStep{{Step: "next"}}},
					{Step: "next", Next: []*models.NextStep{{Step: "end"}}},
					{Step: "end", Next: nil},
				},
			},
			signTypeSet: set.Setify(""),
			wantErr:     ErrUnreachableSteps,
		},
		{
			name: "Valid template",
			template: models.TicketTemplate{
				StartStep: "start",
				EndStep:   []string{"end"},
				Config: []*models.StepConfig{
					{
						Step: "start",
						Next: []*models.NextStep{
							{
								Step: "next",
							},
						},
					},
					{
						Step: "next",
						Next: []*models.NextStep{
							{
								Step: "end",
							},
						},
					},
					{
						Step: "end",
						Next: nil,
					},
				},
			},
			signTypeSet: set.Setify(""),
			wantErr:     nil,
		},
		{
			name: "badJointSignRate",
			template: models.TicketTemplate{
				StartStep: "start",
				EndStep:   []string{"end"},
				Config: []*models.StepConfig{
					{
						Step: "start",
						Next: []*models.NextStep{
							{
								Step: "next",
							},
						},
						Disposal: models.Disposal{
							SignType:      models.JointlySign,
							JointSignRate: 2,
						},
					},
					{
						Step: "next",
						Next: []*models.NextStep{
							{
								Step: "end",
							},
						},
					},
					{
						Step: "end",
						Next: nil,
					},
				},
			},
			signTypeSet: set.Setify(models.JointlySign),
			wantErr:     ErrBadJointSignRate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validator{
				signTypeSet: tt.signTypeSet,
			}
			if err := v.Validate(tt.template); (err == nil) != (tt.wantErr == nil) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewValidator(t *testing.T) {
	tests := []struct {
		name string
		want Validator
	}{
		{
			name: "all is ok",
			want: &validator{
				signTypeSet: models.DisposalSignType,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
