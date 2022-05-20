package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		stat, err := os.Stat("src/files/a.txt")
		if err != nil {
			fmt.Println("错误信息：",err)
		}
		fmt.Println(stat.Name())
		fmt.Println(stat.Size())
		fmt.Println(stat.IsDir())
		fmt.Println(stat.ModTime())
		fmt.Println(stat.Mode())

		fmt.Println("################")

		fmt.Println(filepath.Abs("a.txt"))*/
	/*create, err := os.Create("src/files/b.txt")
	if err!=nil {
		fmt.Println("创建文件失败")

	}*/
	//fmt.Println(create)
	/*file, err := os.OpenFile("src/files/b.txt",os.O_RDWR, os.ModePerm)
	var b []byte

	file.Read(b)
	fmt.Println(string(b))
	fmt.Println(file,err)

	err= os.Remove("src/demo")
	fmt.Println(err)

	file.Close()*/
	/*
		fileName := "src/pic/pic01.jpg"
		file, err := os.Open(fileName)
		outfile, err := os.Create("src/pic/pic02.jpg")
		defer outfile.Close();file.Close()
		if err!=nil {
			fmt.Println("读取文件失败")
		}
		buf :=make([]byte,1024,1024)
		for  {
			_, err := file.Read(buf)
			if io.EOF==err {
				break
			}
			outfile.Write(buf)*/

	create, _ := os.Create("src/files/c.txt")
	fmt.Println(create)
}
