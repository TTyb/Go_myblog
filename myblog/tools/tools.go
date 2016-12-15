package tools

import (
	"github.com/astaxie/beego"
	"os"
	"fmt"
	"strings"
	"path/filepath"
)

//创建递归文件夹
func Createdir(path, dirname string) {
	//获取当前路径的上一级路径
	parent, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Error(err)
	}
	//构造文件夹
	//C:/GOPATH/src/myblog/path/filename
	dir := strings.Replace(parent, "\\", "/", -1) + "/" + path + "/" + dirname

	//创建文件夹
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		beego.Error(err)
	} else {
		beego.Error("创建" + dir + "文件夹成功！")
	}
}

//写入md文件和创建json文件 http://www.cnblogs.com/fengbohello/p/4665883.html http://studygolang.com/articles/4973
//分类、文件名、内容
func Writefile(category, filename, content string) {
	//获取当前路径的上一级路径
	parent, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Error(err)
	}
	//构造文件夹
	//C:/GOPATH/src/myblog/topic/category/filename.md
	filepath := strings.Replace(parent, "\\", "/", -1) + "/topic/" + category + "/" + filename
	fmt.Println(filepath)
	//创建文件
	file, err := os.Create(filepath)
	if err != nil {
		beego.Error(err)
	}
	//最后关闭
	defer file.Close()
	//写入内容
	file.WriteString(content)
}


//生成json结构文件


//检查是否存在这个文件
func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}