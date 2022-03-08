package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

// ini hanya sementara
var SECRET_KEY = []byte("GIN_s3cr3T_k3Y")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// claim atau payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// HEADER:ALGORITHM & TOKEN TYPE and PAYLOAD:DATA
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// VERIFY SIGNATURE
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
