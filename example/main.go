package main

import (
	ape ".."
)

func main() {
	app := ape.NewApp()
	app.Get("/hello", func(req *ape.Request, res *ape.Response) (ape.Any, error) {
		return map[string]string{"hello":"world"}, nil
	})

	sub := ape.NewApp()
	sub.Get("/world", func(req *ape.Request, res *ape.Response) (ape.Any, error) {
		return map[string]string{"hello":"world is mine"}, nil
	})
	app.Mount("/sub", sub)

	app.ListenAndServe(":8080")
}
