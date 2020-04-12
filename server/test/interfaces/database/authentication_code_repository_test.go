package database_test

import (
	mock "github.com/kindaidensan/UMR/test/mock_interfaces/mock_database"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/database"
	"errors"
	"testing"
	"github.com/golang/mock/gomock"
)

func NewMockAuthenticationCodeRepositoryInstance(t *testing.T) (*database.AuthenticationCodeRepository, *mock.MockRedisHandler) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()	
	redis := mock.NewMockRedisHandler(ctrl) 
	repository := database.NewAuthenticationCodeRepository(redis)
	return repository, redis
}

func TestAuthenticationCodeStore(t *testing.T) {
	repo, redis := NewMockAuthenticationCodeRepositoryInstance(t)
	authenticationCode := domain.AuthenticationCode {
		ID: "id",
		Code: "3340",
	}
	redis.EXPECT().Set("id", "3340").Return(nil)
	err := repo.Store(authenticationCode)
	if err != nil {
		t.Errorf("faild: Expectation: return nil")
	}

	redis.EXPECT().Set("id", "3340").Return(errors.New("error"))
	err = repo.Store(authenticationCode)
	if err == nil {
		t.Errorf("faild: Expectation: return err")
	}
}

func TestAuthenticationCodeFindID(t *testing.T) {
	repo, redis := NewMockAuthenticationCodeRepositoryInstance(t)
	redis.EXPECT().Get("id").Return("3340", errors.New("error"))
	_, err := repo.FindID("id")
	if err == nil {
		t.Errorf("faild: Expectation: return err")	
	}

	redis.EXPECT().Get("id").Return("3340", nil)
	authenticationCode, err := repo.FindID("id")
	if err != nil {
		t.Errorf("faild: Expectation: return nil")
	}
	if authenticationCode.ID != "id" {
		t.Errorf("faild: Expectation: return id")
	}
	if authenticationCode.Code != "3340" {
		t.Errorf("faild: Expectation: return 3340")
	}
}