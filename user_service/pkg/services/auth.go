package services

import (
	"context"
	"net/http"

	"e-commerce/user_service/pkg/db"
	"e-commerce/user_service/pkg/models"
	"e-commerce/user_service/pkg/pb"
	"e-commerce/user_service/pkg/utils"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	err := s.H.DB.QueryRow(context.Background(), "select email, password from users where email=$1", req.Email).Scan(&user.Email, &user.Password)
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	if _, err := s.H.DB.Exec(context.Background(), "insert into users(email,password) values($1,$2)", user.Email, user.Password); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Cannot Create the User",
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	err := s.H.DB.QueryRow(context.Background(), "select email, password from users where email=$1", req.Email).Scan(&user.Email, &user.Password)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	err = s.H.DB.QueryRow(context.Background(), "select email, password from users where email=$1", claims.Email).Scan(&user.Email, &user.Password)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
