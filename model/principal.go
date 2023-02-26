package model

import "github.com/golang-jwt/jwt"

type GrpcPrincipal struct {
	UserId           string   `json:"user_id"`
	ExternalId       string   `json:"external_id"`
	Email            string   `json:"email"`
	Username         string   `json:"username"`
	FirstName        string   `json:"first_name"`
	LastName         string   `json:"last_name"`
	Roles            []string `json:"roles"`
	Permissions      []string `json:"permissions"`
	Divisions        []int32  `json:"divisions"`
	MarketCodes      []string `json:"market_codes"`
	RoleIds          []string `json:"role_ids"`
	AssocMarketCodes []string `json:"assoc_market_codes"`
	jwt.StandardClaims
}
