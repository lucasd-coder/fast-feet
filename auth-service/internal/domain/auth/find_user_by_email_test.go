package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lucasd-coder/fast-feet/auth-service/internal/domain/auth"
	"github.com/lucasd-coder/fast-feet/auth-service/internal/mocks"
	"github.com/lucasd-coder/fast-feet/auth-service/internal/provider/validator"
	"github.com/lucasd-coder/fast-feet/auth-service/internal/shared"
	"github.com/lucasd-coder/fast-feet/auth-service/pkg/pb"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindUserByEmailSuite struct {
	suite.Suite
	svc  auth.Service
	repo *mocks.Repository_internal_domain_auth
	ctx  context.Context
}

func (suite *FindUserByEmailSuite) SetupTest() {
	val := validator.NewValidation()
	repo := new(mocks.Repository_internal_domain_auth)

	suite.repo = repo
	suite.svc = auth.NewService(val, repo)
	suite.ctx = context.Background()
}

func (suite *FindUserByEmailSuite) TestFindUserByEmailValidateFailure() {
	pld := auth.FindUserByEmail{}
	_, err := suite.svc.FindUserByEmail(suite.ctx, &pld)
	st, ok := status.FromError(err)
	suite.True(ok)
	suite.Equal(st.Code(), codes.InvalidArgument)
}

func (suite *FindUserByEmailSuite) TestFindUserByEmailFailure() {
	pld := &auth.FindUserByEmail{
		Email: "maria123@gmail.com",
	}

	errUserNotFound := fmt.Errorf("fail called FindUserByEmail %w", shared.ErrUserNotFound)

	suite.repo.On("FindUserByEmail", suite.ctx, pld).Return(nil, errUserNotFound)

	_, err := suite.svc.FindUserByEmail(suite.ctx, pld)
	suite.NotNil(err)
	st, ok := status.FromError(err)
	suite.True(ok)
	suite.Equal(st.Code(), codes.NotFound)
}

func (suite *FindUserByEmailSuite) TestFindUserByEmailSuccess() {
	email := "maria@gmail.com"

	pld := &auth.FindUserByEmail{
		Email: email,
	}

	userRepresentation := &auth.UserRepresentation{
		ID:       "123456",
		Username: "maria",
		Enabled:  true,
		Email:    email,
	}

	rp := &pb.GetUserResponse{
		Id:       "123456",
		Username: "maria",
		Enabled:  true,
		Email:    email,
	}

	suite.repo.On("FindUserByEmail", suite.ctx, pld).Return(userRepresentation, nil)
	resp, err := suite.svc.FindUserByEmail(suite.ctx, pld)
	suite.Nil(err)
	suite.Equal(resp.GetEmail(), rp.GetEmail())
	suite.Equal(resp.GetId(), rp.GetId())
	suite.Equal(resp.GetUsername(), rp.GetUsername())
}

func TestFindUserByEmailSuite(t *testing.T) {
	suite.Run(t, new(FindUserByEmailSuite))
}
