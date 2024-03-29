package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lucasd-coder/fast-feet/auth-service/internal/domain/auth"
	"github.com/lucasd-coder/fast-feet/auth-service/internal/mocks"
	"github.com/lucasd-coder/fast-feet/auth-service/internal/provider/validator"
	"github.com/lucasd-coder/fast-feet/auth-service/internal/shared"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateUserSuite struct {
	suite.Suite
	svc  auth.Service
	repo *mocks.Repository_internal_domain_auth
	ctx  context.Context
}

func (suite *CreateUserSuite) SetupTest() {
	val := validator.NewValidation()
	repo := new(mocks.Repository_internal_domain_auth)

	suite.repo = repo
	suite.svc = auth.NewService(val, repo)
	suite.ctx = context.Background()
}

func (suite *CreateUserSuite) TestCreateUserValidateFailure() {
	pld := auth.Register{}

	_, err := suite.svc.CreateUser(suite.ctx, &pld)
	st, ok := status.FromError(err)
	suite.True(ok)
	suite.Equal(st.Code(), codes.InvalidArgument)
}

func (suite *CreateUserSuite) TestCreateUserRegisterFailure() {
	pld := &auth.Register{
		FirstName: "test",
		LastName:  "123",
		Username:  "maria@gmail.com",
		Password:  "12345@*$",
		Roles:     "USER",
	}

	errUserAlreadyExist := fmt.Errorf("fail called Register %w", shared.ErrUserAlreadyExist)

	suite.repo.On("Register", suite.ctx, pld).Return("", errUserAlreadyExist)

	_, err := suite.svc.CreateUser(suite.ctx, pld)
	suite.NotNil(err)
	suite.ErrorIs(errUserAlreadyExist, err)
}

func (suite *CreateUserSuite) TestCreateUserSuccess() {
	pld := &auth.Register{
		FirstName: "test",
		LastName:  "123",
		Username:  "maria@gmail.com",
		Password:  "12345@*$",
		Roles:     "USER",
	}

	id := "1234567"

	suite.repo.On("Register", suite.ctx, pld).
		Return(id, nil)

	resp, err := suite.svc.CreateUser(suite.ctx, pld)
	suite.Nil(err)
	suite.Equal(resp.GetId(), id)
}

func TestCreateUserSuite(t *testing.T) {
	suite.Run(t, new(CreateUserSuite))
}
