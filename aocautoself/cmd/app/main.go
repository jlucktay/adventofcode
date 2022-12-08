// This app will create some Advent of Code templates for you.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/browser/firefox"

	"go.jlucktay.dev/adventofcode/aocautoself/pkg/fetchaocday"
)

func main() {
	session := ""

	for _, cookieStore := range kooky.FindAllCookieStores() {
		if !cookieStore.IsDefaultProfile() {
			continue
		}

		fmt.Printf("%s -> %s\n", cookieStore.Browser(), cookieStore.FilePath())

		cookies, err := cookieStore.ReadCookies(kooky.Valid, kooky.Domain("adventofcode.com"), kooky.Name("session"))
		if err != nil {
			log.Fatal(err)
		} else if len(cookies) != 1 {
			log.Fatal("wrong number of session cookies")
		}

		fmt.Printf("Expires: %s\n", cookies[0].Expires.Format(time.RFC3339))
		session = cookies[0].Value
	}

	day, err := fetchaocday.Fetch(session, 2017, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(day.String())
}
