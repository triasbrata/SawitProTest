package handler

import (
	"testing"

	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/stretchr/testify/assert"
)

func Test_calculateDronePlan(t *testing.T) {
	type args struct {
		estate       repository.Estate
		trees        []repository.Tree
		max_distance int
	}
	tests := []struct {
		name string
		args args
		want model.DronePlanResponseSuccess
	}{
		{
			name: "success",
			args: args{
				estate: repository.Estate{Width: 3, Length: 6, ID: "a"},
				trees: []repository.Tree{
					{
						EstateID: "",
						X:        3,
						Y:        1,
						Height:   10,
					},
					{
						EstateID: "",
						X:        3,
						Y:        2,
						Height:   10,
					},
					{
						EstateID: "",
						X:        4,
						Y:        2,
						Height:   10,
					},
					{
						EstateID: "",
						X:        6,
						Y:        2,
						Height:   10,
					},
					{
						EstateID: "",
						X:        5,
						Y:        3,
						Height:   10,
					},
				},
				max_distance: 0,
			},
			want: model.DronePlanResponseSuccess{
				Distance: 252,
			},
		},
		{
			name: "success",
			args: args{
				estate: repository.Estate{Width: 1, Length: 5, ID: "a"},
				trees: []repository.Tree{
					{
						EstateID: "",
						X:        2,
						Y:        1,
						Height:   5,
					},
					{
						EstateID: "",
						X:        3,
						Y:        1,
						Height:   3,
					},
					{
						EstateID: "",
						X:        4,
						Y:        1,
						Height:   4,
					},
				},
				max_distance: 0,
			},
			want: model.DronePlanResponseSuccess{
				Distance: 54,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateDronePlan(tt.args.estate, tt.args.trees, tt.args.max_distance)
			assert.Equal(t, tt.want, got)
		})
	}
}
