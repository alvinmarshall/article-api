package main

import "article_api/src/router"

func main() {
	r := router.NewRouter()
	r.Logger.Fatal(r.Start(":3000"))
}
