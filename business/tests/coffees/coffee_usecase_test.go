package tests

import (
	"context"
	"errors"
	"finalProject/business/coffees"
	mock "finalProject/mocks/coffees"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	caseError     = errors.New("error")
	stringTesting = "testing"
	coffeeUsecase coffees.Usecase
)

func TestCoffeeeUsecase_GetById(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		coffeeRepositories mock.MockRepository
	}

	type args struct {
		cid      string
		coffeeID int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(mock *mock.MockRepository)
		want    coffees.Domain
		wantErr error
	}{
		{
			name: "error from repo",
			args: args{
				cid:      "test",
				coffeeID: 0,
			},
			doMock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(ctx, 0).Return(0, caseError)
			},
			want:    coffees.Domain{},
			wantErr: caseError,
		},
		{
			name: "flow success",
			args: args{
				cid:      "test",
				coffeeID: 1,
			},
			doMock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(ctx, 1).
					Return(1, nil)
			},
			want: coffees.Domain{
				Id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.coffeeID)
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := tt.fields.coffeeRepositories

			tt.doMock(&mockRepo)
			got, err := coffeeUsecase.GetByID(ctx, tt.args.coffeeID)
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
