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

	fmt.Println("========================signed jwt key========================")
	fmt.Println(s)

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
	fmt.Println()
	fmt.Println("===========Test creating a new jwt token with custom claims ============")
	mycustomtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := mycustomtoken.SignedString([]byte("hihihi"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ss)

	// verifiedParsedToken, err := jwt.Parse(ss, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte("hihihi"), nil
	// })

	// testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYSI6ImhleWZvb2ZvbyIsImlzcyI6InRlc3QiLCJzdWIiOiJzb21lYm9keSIsImF1ZCI6WyJzb21lYm9keV9lbHNlIiwiam9yZGFuIl0sIm5iZiI6MTcwNjk1NTk0NywiaWF0IjoxNzA2OTU1OTQ3LCJqdGkiOiIxIn0.dZaozaQ6hr4oknRZfaPwVxC3u99IARm2AqH892oHvo0"
	// verifiedParsedToken, err := jwt.Parse(testToken, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte("hihihi"), nil
	// })

	verifiedParsedToken, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("hihihi"), nil
	})

	if err != nil {
		fmt.Println()
		fmt.Println("------verifiedParsedToken error------")
		log.Fatalln(err)
	}

	// fmt.Println(verifiedParsedToken)
	issuedAt, err := verifiedParsedToken.Claims.GetIssuedAt()
	if err != nil {
		fmt.Println()
		fmt.Println("------GetIssuedAt error------")
		log.Fatalln(err)
	}
	fmt.Println()
	fmt.Printf("test token issued at : %s", issuedAt)

	// d, err := json.Marshal(verifiedParsedToken)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("*******************s")
	// log.Println(string(d))
}
