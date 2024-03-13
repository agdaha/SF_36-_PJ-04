package memdb

import (
	"reflect"
	"skillfactory/SF_36-_PJ-04/internal/app/model"
	"testing"
)

func TestStore_Posts(t *testing.T) {
	type fields struct {
		db []model.Post
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Post
		wantErr bool
	}{
		{
			name:   "memdb test",
			fields: fields{},
			args:   args{n: 1},
			want: []model.Post{
				{
					Id:          0,
					Title:       "Title 0",
					Description: "Description 0",
					Author:      "Author 0",
					Link:        "Link 0",
					Guid:        "Guid 0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				db: tt.fields.db,
			}
			got, err := s.Posts(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.Posts() = %v, want %v", got, tt.want)
			}
		})
	}
}
