package tests

import (
	"context"
	"errors"
	"finalProject/business/articles"
	mock "finalProject/mocks/articles"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	caseError      = errors.New("error")
	stringTesting  = "testing"
	articleUsecase articles.Usecase
)

func TestArticleUsecase_GetById(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		articleRepositories mock.MockRepository
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
		want    articles.Domain
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
			want:    articles.Domain{},
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
			want: articles.Domain{
				Id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.articleID)
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := tt.fields.articleRepositories

			tt.doMock(&mockRepo)
			got, err := articleUsecase.GetByID(ctx, tt.args.articleID)
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
