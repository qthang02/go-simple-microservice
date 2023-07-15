package gapi

import (
	db "authentication-service/db/sqlc"
	"authentication-service/pb"
	"authentication-service/token"
	"authentication-service/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	pb.UnimplementedMicroGoServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
