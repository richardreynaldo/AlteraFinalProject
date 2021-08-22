package tests

import (
	"context"
	"errors"
	"finalProject/business/reviews"
	mock "finalProject/mocks/reviews"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	caseError     = errors.New("error")
	stringTesting = "testing"
	reviewUsecase reviews.Usecase
)

func TestArticleUsecase_GetById(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		reviewRepository mock.MockRepository
	}

	type args struct {
		cid       string
		articleID int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(mock *mock.MockRepository)
		want    reviews.Domain
		wantErr error
	}{
		{
			name: "error from repo",
			args: args{
				cid:       "test",
				articleID: 0,
			},
			doMock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(ctx, 0).Return(0, caseError)
			},
			want:    reviews.Domain{},
			wantErr: caseError,
		},
		{
			name: "flow success",
			args: args{
				cid:       "test",
				articleID: 1,
			},
			doMock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(ctx, 1).
					Return(1, nil)
			},
			want: reviews.Domain{
				Id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.articleID)
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := tt.fields.reviewRepository

			tt.doMock(&mockRepo)
			got, err := reviewUsecase.GetByID(ctx, tt.args.articleID)
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
