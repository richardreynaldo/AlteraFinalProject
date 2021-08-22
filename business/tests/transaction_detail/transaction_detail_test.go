package tests

import (
	"context"
	"errors"
	"finalProject/business/transaction_detail"
	mock "finalProject/mocks/articles"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	caseError          = errors.New("error")
	stringTesting      = "testing"
	transactionDetailUsecase transaction_detail.Usecase
)

func TestTransactionDetailUsecase_GetById(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		transactionDetailRepo mock.MockRepository
	}

	type args struct {
		cid       string
		transactionDetailID int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(mock *mock.MockRepository)
		want    transaction_detail.Domain
		wantErr error
	}{
		{
			name: "error from repo",
			args: args{
				cid:       "test",
				transactionDetailID: 0,
			},
			doMock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(ctx, 0).Return(0, caseError)
			},
			want:    transaction_detail.Domain{},
			wantErr: caseError,
		},
		{
			name: "flow success",
			args: args{
				cid:       "test",
				transactionDetailID: 1,
			},
			doMock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(ctx, 1).
					Return(1, nil)
			},
			want: transaction_detail.Domain{
				Id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.transactionDetailID)
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := tt.fields.transactionDetailRepo

			tt.doMock(&mockRepo)
			got, err := transactionDetailUsecase.GetByID(ctx, tt.args.transactionDetailID)
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
