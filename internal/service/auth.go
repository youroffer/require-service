package service

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/himmel520/uoffer/require/roles"
)

func (s *Service) GetUserRoleFromToken(jwtToken string, publicKey *rsa.PublicKey) (string, error) {
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token claims are not of type jwt.MapClaims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", errors.New("role claim is not a string")
	}

	return role, err
}

func (s *Service) IsUserAuthorized(requiredRole, userRole string) bool {
	rolesHierarchy := map[string]int{
		roles.Anonym: 0,
		roles.Admin:  1,
		roles.User:   2,
	}

	return rolesHierarchy[userRole] >= rolesHierarchy[requiredRole]
}
