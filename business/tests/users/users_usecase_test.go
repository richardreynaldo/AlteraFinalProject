package tests

import (
	"context"
	"errors"
	"finalProject/business/users"
	mock "finalProject/mocks/users"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	caseError     = errors.New("error")
	stringTesting = "testing"
	userUsecase   users.UseCase
)

func TestArticleUsecase_GetById(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		userRepo mock.MockRepository
	}

	type args struct {
		cid    string
		userID int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(mock *mock.MockRepository)
		want    users.Domain
		wantErr error
	}{
		{
			name: "error from repo",
			args: args{
				cid:    "test",
				userID: 0,
			},
			doMock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(ctx, 0).Return(0, caseError)
			},
			want:    users.Domain{},
			wantErr: caseError,
		},
		{
			name: "flow success",
			args: args{
				cid:    "test",
				userID: 1,
			},
			doMock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(ctx, 1).
					Return(1, nil)
			},
			want: users.Domain{
				Id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.userID)
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := tt.fields.userRepo

			tt.doMock(&mockRepo)
			got, err := userUsecase.GetByID(ctx, tt.args.userID)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("articleUsecase.getById error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleUsecase.getById  = %v, want %v", got, tt.want)
			}
		})
	}
}
