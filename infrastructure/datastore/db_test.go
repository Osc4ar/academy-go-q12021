package datastore

import (
	"reflect"
	"taskmanager/domain/model"
	"testing"
)

func Test_dbType_FindByID(t *testing.T) {
	type fields struct {
		tasks []*model.Task
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Task
		wantErr bool
	}{
		{
			name: "Not found ID",
			fields: fields{
				tasks: []*model.Task{},
			},
			args:    args{id: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found ID",
			fields: fields{
				tasks: []*model.Task{
					{
						ID: 1,
					},
				},
			},
			args: args{id: 1},
			want: &model.Task{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dbType{
				tasks: tt.fields.tasks,
			}
			got, err := d.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("dbType.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbType.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbType_FindAll(t *testing.T) {
	type fields struct {
		tasks []*model.Task
	}
	tests := []struct {
		name   string
		fields fields
		want   []*model.Task
	}{
		{
			name: "Empty DB",
			fields: fields{
				tasks: []*model.Task{},
			},
			want: []*model.Task{},
		},
		{
			name: "DB with records",
			fields: fields{
				tasks: []*model.Task{
					{
						ID: 1,
					},
				},
			},
			want: []*model.Task{
				{
					ID: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dbType{
				tasks: tt.fields.tasks,
			}
			if got := d.FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbType.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTaskToSlice(t *testing.T) {
	type args struct {
		t *model.Task
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Valid record",
			args: args{
				&model.Task{
					ID:          1,
					Content:     "Any Content",
					Completed:   true,
					DueDate:     "10/15/2021",
					WorkingTime: 0,
				},
			},
			want: []string{"1", "Any Content", "true", "10/15/2021", "0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTaskToSlice(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTaskToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recordToTask(t *testing.T) {
	type args struct {
		record []string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Task
		wantErr bool
	}{
		{
			name: "Valid Task",
			args: args{
				record: []string{"1", "Any Content", "false", "10/15/2021", "12"},
			},
			want:    model.Task{ID: 1, Content: "Any Content", Completed: false, DueDate: "10/15/2021", WorkingTime: 12},
			wantErr: false,
		},
		{
			name: "Invalid ID",
			args: args{
				record: []string{"a", "Any Content", "false", "10/15/2021", "12"},
			},
			want:    model.Task{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := recordToTask(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("recordToTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recordToTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
