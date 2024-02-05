package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// token with no claims
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	/* Load key from somewhere, for example an environment variable */
	key = []byte("hohoho")
	// t = jwt.New(jwt.SigningMethodHS256)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": "john",
			"foo": 2,
		})
	s, err := t.SignedString(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	// b, err := json.Marshal(t)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(b))

	// custom claims
	type MyCustomClaims struct {
		Ba  string `json:"ba"`
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		"blacksheep",
		"heyfoofoo",
		jwt.RegisteredClaims{
			// Issuer:  "test",
			// Subject: "mysubject",
			// A usual scenario is to set the expiration time relative to the current time
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else", "jordan"},
		},
	}

	// c, _ := json.Marshal(claims)
	fmt.Println("=======================")
	// fmt.Printf(string(c))
	// jwt.SigningMethodHS256
	mycustomtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := mycustomtoken.SignedString([]byte("hihihi"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ss)

	// parsedToken, err := jwt.Parse(ss, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte("hihihi"), nil
	// })

	// testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYSI6ImhleWZvb2ZvbyIsImlzcyI6InRlc3QiLCJzdWIiOiJzb21lYm9keSIsImF1ZCI6WyJzb21lYm9keV9lbHNlIiwiam9yZGFuIl0sIm5iZiI6MTcwNjk1NTk0NywiaWF0IjoxNzA2OTU1OTQ3LCJqdGkiOiIxIn0.dZaozaQ6hr4oknRZfaPwVxC3u99IARm2AqH892oHvo0"
	// parsedToken, err := jwt.Parse(testToken, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte("hihihi"), nil
	// })

	parsedToken, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("hihihi"), nil
	})

	if err != nil {
		log.Println(err)
	}

	// fmt.Println(parsedToken)

	fmt.Println(parsedToken.Claims.GetIssuedAt())

	// d, err := json.Marshal(parsedToken)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("*******************s")
	// log.Println(string(d))
}
