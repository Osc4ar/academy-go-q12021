package presenter

import (
	"reflect"
	"taskmanager/domain/model"
	"testing"
)

func TestNewTaskPresenter(t *testing.T) {
	tests := []struct {
		name string
		want TaskPresenter
	}{
		{
			name: "Normal Behaviour", want: &taskPresenter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskPresenter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskPresenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskPresenter_ResponseTasks(t *testing.T) {
	type args struct {
		t []*model.Task
	}
	tests := []struct {
		name string
		tp   *taskPresenter
		args args
		want []*model.Task
	}{
		{
			name: "Done Task",
			tp:   &taskPresenter{},
			args: args{
				t: []*model.Task{
					{
						ID:          1,
						Content:     "Any Content",
						Completed:   true,
						DueDate:     "12/12/2021",
						WorkingTime: 0,
					},
				},
			},
			want: []*model.Task{
				{
					ID:          1,
					Content:     "Done: Any Content",
					Completed:   true,
					DueDate:     "12/12/2021",
					WorkingTime: 0,
				},
			},
		},
		{
			name: "ToDo Task",
			tp:   &taskPresenter{},
			args: args{
				t: []*model.Task{
					{
						ID:          1,
						Content:     "Any Content",
						Completed:   false,
						DueDate:     "12/12/2021",
						WorkingTime: 0,
					},
				},
			},
			want: []*model.Task{
				{
					ID:          1,
					Content:     "ToDo: Any Content",
					Completed:   false,
					DueDate:     "12/12/2021",
					WorkingTime: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := &taskPresenter{}
			if got := tp.ResponseTasks(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskPresenter.ResponseTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskPresenter_ResponseTask(t *testing.T) {
	type args struct {
		t *model.Task
	}
	tests := []struct {
		name string
		tp   *taskPresenter
		args args
		want *model.Task
	}{
		{
			name: "Done Task",
			tp:   &taskPresenter{},
			args: args{
				t: &model.Task{
					ID:          1,
					Content:     "Any Content",
					Completed:   true,
					DueDate:     "12/12/2021",
					WorkingTime: 0,
				},
			},
			want: &model.Task{
				ID:          1,
				Content:     "Done: Any Content",
				Completed:   true,
				DueDate:     "12/12/2021",
				WorkingTime: 0,
			},
		},
		{
			name: "ToDo Task",
			tp:   &taskPresenter{},
			args: args{
				t: &model.Task{
					ID:          1,
					Content:     "Any Content",
					Completed:   false,
					DueDate:     "12/12/2021",
					WorkingTime: 0,
				},
			},
			want: &model.Task{
				ID:          1,
				Content:     "ToDo: Any Content",
				Completed:   false,
				DueDate:     "12/12/2021",
				WorkingTime: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := &taskPresenter{}
			if got := tp.ResponseTask(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskPresenter.ResponseTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
