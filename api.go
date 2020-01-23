package main

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/bubble", func(ctx iris.Context) {
		ctx.Write(parseAndSort([]byte("[" + ctx.URLParam("array") + "]")))
	})
	app.Post("/bubble", func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		ctx.Write(parseAndSort(body))
	})
	app.Run(iris.Addr(":80"))
}

func parseAndSort(bytes []byte) []byte {
	var array []int
	json.Unmarshal(bytes, &array)

	b, _ := json.Marshal(map[string][]int{"result": sort(array)})

	return b
}

func sort(array []int) []int {
	length := len(array)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if array[j] > array[j+1] {
				buf := array[j]
				array[j] = array[j+1]
				array[j+1] = buf
			}
		}
	}
	return array
}
