package repo_test

import (
	"testing"

	"github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/user/dto"
	mock_repo "github.com/agusheryanto182/go-health-record/module/feature/user/repo/mocks"
	"github.com/golang/mock/gomock"
)

func TestCheckUserByIdAndRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockUserRepo.EXPECT().CheckUserByIdAndRole("user123", "role123").Return(true, nil)

	ok, err := mockUserRepo.CheckUserByIdAndRole("user123", "role123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !ok {
		t.Error("Expected user to exist with the given role")
	}
}

func TestDeleteUserNurse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockUserRepo.EXPECT().DeleteUserNurse(gomock.Any()).Return(nil)

	err := mockUserRepo.DeleteUserNurse(nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockUser := &entities.User{}
	mockUserRepo.EXPECT().GetUser(int64(123), "someRole").Return(mockUser, nil)

	user, err := mockUserRepo.GetUser(123, "someRole")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if user == nil {
		t.Error("Expected user object, got nil")
	}
}

func TestGetUserByFilters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockFilter := &dto.UserFilter{}
	mockUserRepo.EXPECT().GetUserByFilters(mockFilter).Return(nil, nil)

	users, err := mockUserRepo.GetUserByFilters(mockFilter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if users != nil {
		t.Error("Expected nil users")
	}
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockUser := &entities.User{}
	mockUserRepo.EXPECT().GetUserByID("user123").Return(mockUser, nil)

	user, err := mockUserRepo.GetUserByID("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if user == nil {
		t.Error("Expected user object, got nil")
	}
}

func TestIsNipExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockUserRepo.EXPECT().IsNipExist(int64(123)).Return(true, nil)

	exists, err := mockUserRepo.IsNipExist(123)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !exists {
		t.Error("Expected NIP to exist")
	}
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockUserRepo.EXPECT().Register(gomock.Any()).Return("userID123", nil)

	userID, err := mockUserRepo.Register(&entities.User{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if userID != "userID123" {
		t.Errorf("Expected userID to be 'userID123', got '%s'", userID)
	}
}

func TestSetPasswordNurse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockUserRepo.EXPECT().SetPasswordNurse(gomock.Any()).Return(nil)

	err := mockUserRepo.SetPasswordNurse(nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestUpdateUserNurse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockUserRepoInterface(ctrl)
	mockUserRepo.EXPECT().UpdateUserNurse(gomock.Any()).Return(nil)

	err := mockUserRepo.UpdateUserNurse(nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
