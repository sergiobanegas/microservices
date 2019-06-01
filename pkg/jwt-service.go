package pkg

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

var mySigningKey = []byte("secret")

var clientIdKey = "clientId"

var authorizationHeader = "Authorization"
var accessTokenHeader = "AccessToken"

type TokenService struct {
}

func (TokenService) GenerateToken(clientId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[clientIdKey] = clientId
	return token.SignedString(mySigningKey)
}

func (s *TokenService) GetClientFromContext(ctx context.Context) (string, error) {
	token, err := s.GetAuthToken(ctx)
	if err != nil {
		return "", err
	}
	return s.GetClientId(token)
}

func (s TokenService) GetClientId(tokenString string) (string, error) {
	token, err := s.decodeToken(tokenString)
	if err != nil {
		return "", status.Error(codes.PermissionDenied, "Invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims[clientIdKey].(string), nil
}

func (s *TokenService) GetAuthToken(ctx context.Context) (string, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	authHeaders := headers.Get(authorizationHeader)
	if len(authHeaders) == 0 {
		authHeaders = headers.Get(accessTokenHeader)
		if len(authHeaders) == 0 {
			return "", status.Error(codes.PermissionDenied, "Authorization header missing")
		}
		return authHeaders[0], nil
	}
	authorizationHeader := strings.Split(authHeaders[0], " ")
	if len(authorizationHeader) != 2 || authorizationHeader[0] != "Bearer" {
		return "", status.Error(codes.PermissionDenied, "Invalid token, must be format: Bearer + token")
	}
	return authorizationHeader[1], nil
}

func (TokenService) decodeToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
}
