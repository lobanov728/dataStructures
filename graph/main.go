package main

import (
	"fmt"
)

type Node struct {
	Nodes []*Node
	Name  string
}

func main() {
	fmt.Println("Hello")
	graph := map[string][]string{}

	//graph["you"] = []string{"Alice", "Bob", "Claire"}
	//graph["Bob"] = []string{"Anuj", "Peggy"}
	//graph["Alice"] = []string{"Peggy"}
	//graph["Claire"] = []string{"Tom", "Jonny"}
	//graph["Anuj"] = []string{}
	//graph["Peggy"] = []string{}
	//graph["Jonny"] = []string{}
	//graph["Tom"] = []string{}
	//
	//var queue []string
	//for _, name := range graph["you"] {
	//	queue = append(queue, name)
	//}
	//fmt.Println(queue)
	//var person string
	//for ; len(queue) > 0; {
	//	fmt.Println(len(queue))
	//	person, queue = queue[0], queue[1:]
	//	fmt.Println(len(queue))
	//	if personIsSeller(person) {
	//		fmt.Println(person, "is a mango seller")
	//	} else {
	//		personSlice, ok := graph[person]
	//		if ok {
	//			for _, name := range personSlice {
	//				queue = append(queue, name)
	//			}
	//		}
	//	}
	//}

	//graph := &Node{
	//	Nodes: []*Node{
	//		&Node{Name: "Alice"},
	//		&Node{Name: "Bob"},
	//		&Node{Name: "Claire"},
	//	},
	//	Name: "Iam",
	//}

	graph["jacket"] = []string{"belt", "tie"}
	graph["tie"] = []string{"shirt"}
	graph["shirt"] = []string{}
	graph["belt"] = []string{"pants"}
	graph["pants"] = []string{"underwear", "shirt"}
	graph["underwear"] = []string{}
	graph["right sock"] = []string{}
	graph["left sock"] = []string{}
	graph["right shoe"] = []string{"pants", "right sock"}
	graph["left shoe"] = []string{"pants", "left sock"}
	var queue []string
	for _, name := range graph["right shoe"] {
		queue = append(queue, name)
	}
	var person string
	for ; len(queue) > 0; {
		fmt.Println(len(queue))
		person, queue = queue[0], queue[1:]
		fmt.Println(person)
		personSlice, ok := graph[person]
		if ok {
			for _, name := range personSlice {
				queue = append(queue, name)
			}
		}
	}
}

func personIsSeller(s string) bool {
	return string([]rune(s)[len(s)-1]) == "m"
}
