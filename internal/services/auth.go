package services

type AuthService struct {
	jwtSecret []byte
}

func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{
		jwtSecret: []byte(jwtSecret),
	}
}

// func (s *AuthService) GenerateToken(user *models.User) (string, error) {
//     tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//         "user_id": user.ID,
//         "exp": time.Now().Add(time.Hour *24).Unix(),
//     })
//
//     return token.SignedString(s.jwtSecret)
// }
