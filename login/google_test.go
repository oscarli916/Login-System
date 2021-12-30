package login

import (
	"fmt"
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func generateToken(claims GoogleClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SIGNINGKEY)
	if err != nil {
		fmt.Println("Token is corrupted. ", err)
	}
	return tokenString
}

type output struct {
	data *userdata
	err  error
}

func Test_googleLogin_Login(t *testing.T) {
	tests := []struct {
		name     string
		l        googleLogin
		expected output
	}{
		// TODO: Add test cases.
		// Test case #1
		{
			name: "SourceTest: Pass issuer as Google",
			l: googleLogin{generateToken(
				GoogleClaims{
					"testuser@gmail.com",
					"true",
					"Test User",
					"https://lh4.googleusercontent.com/-kYgzyAWpZzJ/ABCDEFGHI/AAAJKLMNOP/tIXL9Ir44LE/s99-c/photo.jpg",
					"Test",
					"User",
					"en",
					jwt.StandardClaims{
						Audience:  "1008719970978-hb24n2dstb40o45d4feuo2ukqmcc6381.apps.googleusercontent.com", // aud
						ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),                                          // exp
						IssuedAt:  time.Now().AddDate(0, 0, -1).Unix(),                                         // iat
						Issuer:    "https://accounts.google.com",                                               // iss
						Subject:   "110169484474386276334",                                                     // sub
					},
				})},
			expected: output{data: &userdata{
				"email":          "testuser@gmail.com",
				"email_verified": "testuser@true.com",
				"name":           "Test User@gmail.com",
				"picture":        "https://lh4.googleusercontent.com/-kYgzyAWpZzJ/ABCDEFGHI/AAAJKLMNOP/tIXL9Ir44LE/s99-c/photo.jpg@gmail.com",
				"given_name":     "given_name",
				"family_name":    "User",
				"locale":         "en",
			}, err: nil},
		},
		// Test case #2
		{
			name: "SourceTest: Pass issuer as AWS cognito",
			l: googleLogin{generateToken(
				GoogleClaims{
					"testuser@gmail.com",
					"true",
					"Test User",
					"https://lh4.googleusercontent.com/-kYgzyAWpZzJ/ABCDEFGHI/AAAJKLMNOP/tIXL9Ir44LE/s99-c/photo.jpg",
					"Test",
					"User",
					"en",
					jwt.StandardClaims{
						Audience:  "1008719970978-hb24n2dstb40o45d4feuo2ukqmcc6381.apps.googleusercontent.com", // aud
						ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),                                          // exp
						IssuedAt:  time.Now().AddDate(0, 0, -1).Unix(),                                         // iat
						Issuer:    "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_CtRgepXa7",           // iss
						Subject:   "110169484474386276334",                                                     // sub
					},
				})},
			expected: output{data: nil, err: UnsupportedIssuerErrorHandler{}},
		},
		// Test case #3
		{
			name: "ExpiryTest: Pass expired token",
			l: googleLogin{generateToken(
				GoogleClaims{
					"testuser@gmail.com",
					"true",
					"Test User",
					"https://lh4.googleusercontent.com/-kYgzyAWpZzJ/ABCDEFGHI/AAAJKLMNOP/tIXL9Ir44LE/s99-c/photo.jpg",
					"Test",
					"User",
					"en",
					jwt.StandardClaims{
						Audience:  "1008719970978-hb24n2dstb40o45d4feuo2ukqmcc6381.apps.googleusercontent.com", // aud
						ExpiresAt: time.Now().AddDate(0, 0, -1).Unix(),                                         // exp
						IssuedAt:  time.Now().AddDate(0, 0, -2).Unix(),                                         // iat
						Issuer:    "https://accounts.google.com",                                               // iss
						Subject:   "110169484474386276334",                                                     // sub
					},
				})},
			expected: output{data: nil, err: TokenExpiredErrorHandler{}},
		},
		// Test case #4
		{
			name: "IssuedTest: Pass not yet issued token",
			l: googleLogin{generateToken(
				GoogleClaims{
					"testuser@gmail.com",
					"true",
					"Test User",
					"https://lh4.googleusercontent.com/-kYgzyAWpZzJ/ABCDEFGHI/AAAJKLMNOP/tIXL9Ir44LE/s99-c/photo.jpg",
					"Test",
					"User",
					"en",
					jwt.StandardClaims{
						Audience:  "1008719970978-hb24n2dstb40o45d4feuo2ukqmcc6381.apps.googleusercontent.com", // aud
						ExpiresAt: time.Now().AddDate(0, 0, 2).Unix(),                                          // exp
						IssuedAt:  time.Now().AddDate(0, 0, 1).Unix(),                                          // iat
						Issuer:    "https://accounts.google.com",                                               // iss
						Subject:   "110169484474386276334",                                                     // sub
					},
				})},
			expected: output{data: nil, err: TokenNotIssuedErrorHandler{}},
		},
		// Test case #5
		{
			name: "EmotyTest: Pass empty token",
			l: googleLogin{generateToken(
				GoogleClaims{
					"",
					"",
					"",
					"",
					"",
					"",
					"",
					jwt.StandardClaims{},
				})},
			expected: output{data: nil, err: EmptyTokenErrorHandler{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.l.Login()
			if err == nil && tt.expected.err != nil {
				t.Errorf("[%v]Unmatched error, expected %v, got %v \n", tt.name, tt.expected.err.Error(), err)
			} else if err != nil && tt.expected.err == nil {
				t.Errorf("[%v]Unmatched error, expected %v, got %v \n", tt.name, tt.expected.err, err.Error())
			} else {
				if (err != nil && tt.expected.err != nil) && (tt.expected.err.Error() != err.Error()) {
					t.Errorf("[%v]Unmatched error, expected %v, got %v \n", tt.name, tt.expected.err.Error(), err.Error())
				}
			}
		})
	}
}

func Test_googleLogin_GetUserData(t *testing.T) {
	tests := []struct {
		name string
		l    googleLogin
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.GetUserData()
		})
	}
}

func Test_googleLogin_getUserToken(t *testing.T) {
	tests := []struct {
		name string
		l    googleLogin
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.getUserToken(); got != tt.want {
				t.Errorf("googleLogin.getUserToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
