package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"taskmanager/domain/model"
	"taskmanager/usecase/interactor"
	"testing"
)

func TestNewTaskController(t *testing.T) {
	type args struct {
		ti interactor.TaskInteractor
	}
	tests := []struct {
		name string
		args args
		want TaskController
	}{
		{
			name: "Default behaviour",
			args: args{
				ti: nil,
			},
			want: &taskController{nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskController(tt.args.ti); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskController_GetTasks(t *testing.T) {
	mockedTI := interactor.NewMockTaskInteractor([]*model.Task{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	})

	mockedWriter := httptest.NewRecorder()

	type fields struct {
		taskInteractor interactor.TaskInteractor
	}
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Get Tasks",
			fields: fields{
				taskInteractor: mockedTI,
			},
			args: args{
				w: mockedWriter,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := &taskController{
				taskInteractor: tt.fields.taskInteractor,
			}
			if err := tc.GetTasks(tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("taskController.GetTasks() error = %v, wantErr %v", err, tt.wantErr)
			}

			var expectedTasks []*model.Task
			expectedTasks, _ = mockedTI.GetAll(expectedTasks)

			var actualTasks []*model.Task
			if err := json.Unmarshal(mockedWriter.Body.Bytes(), &actualTasks); err != nil {
				t.Errorf("Cannot parse reponse into Tasks")
			}

			if *expectedTasks[0] != *actualTasks[0] {
				t.Errorf("Wrong response body\nExpected: %v\nActual: %v", expectedTasks, actualTasks)
			}
		})
	}
}

func Test_taskController_GetTask(t *testing.T) {
	mockedTI := interactor.NewMockTaskInteractor([]*model.Task{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	})

	mockedWriter := httptest.NewRecorder()

	type fields struct {
		taskInteractor interactor.TaskInteractor
	}
	type args struct {
		id uint
		w  http.ResponseWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Get Existing Task",
			fields: fields{
				taskInteractor: mockedTI,
			},
			args: args{
				id: 1,
				w: mockedWriter,
			},
			wantErr: false,
		},
		{
			name: "Get Non-Existing Task",
			fields: fields{
				taskInteractor: mockedTI,
			},
			args: args{
				id: 3,
				w: mockedWriter,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := &taskController{
				taskInteractor: tt.fields.taskInteractor,
			}
			if err := tc.GetTask(tt.args.id, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("taskController.GetTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
