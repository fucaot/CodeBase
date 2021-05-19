package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}

// output:
// [{Casablanca 1942 false [Humphrey Bogart Ingrid Bergman]} {Casablanca 1967 true [Paul Newman]}]
// [{"Title":"Casablanca","Year":1942,"Color":false,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Casablanca","Year":1967,"Color":true,"Actors":["Paul Newman"]}]
// [
//     {
//         "Title": "Casablanca",
//         "Year": 1942,
//         "Color": false,
//         "Actors": [
//             "Humphrey Bogart",
//             "Ingrid Bergman"
//         ]
//     },
//     {
//         "Title": "Casablanca",
//         "Year": 1967,
//         "Color": true,
//         "Actors": [
//             "Paul Newman"
//         ]
//     }
// ]
func main() {
	fmt.Println("vim-go")
	var movies = []Movie{
		{
			Title: "Casablanca", Year: 1942, Color: false, Actors: []string{
				"Humphrey Bogart", "Ingrid Bergman",
			},
		},
		{
			Title: "Casablanca", Year: 1967, Color: true, Actors: []string{
				"Paul Newman",
			},
		},
	}
	// 直接打印
	fmt.Println(movies)

	// 通过json.Marshal 将切片（Slice）转化为json格式, 返回格式为byte数组
	jsonByte, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// 打印json
	fmt.Println(string(jsonByte))

	// 进行格式化
	jsonByte, err = json.MarshalIndent(movies, "", "    ")

	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println(string(jsonByte))
}
