package main

import "utils/folder"

// go run main.go add --id 2 --title 无敌风火轮 --auther 佚名
// go run main.go get --all true
func main() {
	// getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	// all := getCmd.Bool("all", false, "是否选择全部")
	// id := getCmd.String("id", "", "设置 id")

	// addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	// addId := addCmd.String("id", "", "book id")
	// addTitle := addCmd.String("title", "", "book title")
	// addAuther := addCmd.String("auther", "", "book auther")
	// switch os.Args[1] {
	// case "get":
	// 	list, _ := book.GetBook(getCmd, id, all)
	// 	fmt.Println(list)
	// case "add":
	// 	book.AddBook(addCmd, addId, addTitle, addAuther, true)
	// }
	// fmt.Println(*all, *id)
	// fmt.Println(os.Args[1])
	// #1 根据 ID 、all 获取 books
	// #2 添加 books id title  author price image_url

	// #3 更新 book id必填其它可选

	// #4 删除 book id
	// book.GetBook(addCmd, *id, *all)

	folder.Register()
}
