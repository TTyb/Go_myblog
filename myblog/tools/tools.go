package tools

import (
	"github.com/astaxie/beego"
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"
	"bufio"
	"fmt"
)


// 文章列表
type Topic struct {
	Title       string
	Createdtime string
	Content     string
	Category    string
}

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


//检查是否存在这个文件
//func checkFileIsExist(filename string) (bool) {
//	var exist = true;
//	if _, err := os.Stat(filename); os.IsNotExist(err) {
//		exist = false;
//	}
//	return exist;
//}

//遍历dirPth目录下，后缀为suffix的文件
//path := "C:/GOPATH/src/myremember"，suffix := "md"
//path:="C:/GOPATH/src/myblog/topic/百哥"--[C:/GOPATH/src/myblog/topic/百哥/2016-12-15#卧槽.md] <nil>
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 1000)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() {
			// 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			//匹配文件
			//将strings.Replace(替换的string, "\\", "/", -1)
			files = append(files, strings.Replace(dirPth + PthSep + fi.Name(), "\\", "/", -1))
		}
	}
	return files, nil
}

//遍历文件夹下的所有文件，返回的是所有文件的名字，包括文件和文件夹的名字
//path="C:/GOPATH/src/myblog"
func GetList(path string) (files []string) {
	files = make([]string, 0, 1000)
	dir_list, err := ioutil.ReadDir(path)
	if err != nil {
		beego.Error(err)
	}

	for _, v := range dir_list {
		files = append(files, v.Name())
	}
	return files
}

//获取所有的文章列表
func GetAllTopics() ([]Topic, error) {

	var Title string
	var Createdtime string
	var Content string
	var Category string

	var cate []Topic
	//获取当前路径的上一级路径
	parent, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Error(err)
	}
	//C:/GOPATH/src/myblog/topic/
	filetopic := strings.Replace(parent, "\\", "/", -1) + "/topic/"
	//获取所有分类--[百哥]
	categories := GetList(filetopic)
	for _, n := range categories {

		//获得分类
		Category = n
		//每个分类的文件夹目录为
		//C:/GOPATH/src/myblog/topic/百哥
		catefilepath := filetopic + n
		//打开每个分类获取里面的md信息
		mdname := GetList(catefilepath)
		//遍历每个md的名字
		for _, mname := range mdname {
			//获得创建时间
			Createdtime = strings.Split(mname, "#")[0]
			//获得文章标题
			Title = strings.Split(strings.Split(mname, "#")[1], ".")[0]
			//获得每个md的内容
			Content = readfile(catefilepath + "/" + mname)
			cates := Topic{Title:Title, Createdtime:Createdtime, Category:Category, Content:Content}
			cate = append(cate, cates)
		}
	}
	return cate, err
}

//读取文件的内容
func readfile(path string) (content string) {

	content = ""
	file, err := os.Open(path)
	if err != nil {
		beego.Error(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = content + strings.TrimSpace(scanner.Text()) + "\n"
	}

	return content
}

//获取文章的详情
func GetTopicsContent(title string) ([]Topic, error) {

	var cate []Topic
	topics, err := GetAllTopics()
	if err != nil {
		beego.Error(err)
	}

	for _, maps := range topics {
		//如果title一样就返回这个文件的内容
		if maps.Title == title {
			cate = append(cate, maps)
		}
	}
	return cate, err
}

//删除文章
func DeleteTopic(title, createdtime, category string) error {
	//获取当前路径的上一级路径
	parent, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Error(err)
	}
	//C:/GOPATH/src/myblog/topic/
	topicname := strings.Replace(parent, "\\", "/", -1) + "/topic/" + category + "/" + createdtime + "#" + title + ".md"
	fmt.Println(topicname)
	err = os.Remove(topicname)

	//读取这个分类是否还有md文件
	juegefile := strings.Replace(parent, "\\", "/", -1) + "/topic/" + category + "/"
	juege := GetList(juegefile)
	//如果没有就删掉这个分类
	if len(juege) == 0 {
		err = os.Remove(juegefile)
	}
	return err
}

//获取某个单一分类的列表
func GetCategoryContent(category string) ([]Topic, error) {

	var Title string
	var Createdtime string
	var Content string

	var cate []Topic

	//获取当前路径的上一级路径
	parent, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Error(err)
	}
	//C:/GOPATH/src/myblog/topic/
	categorypath := strings.Replace(parent, "\\", "/", -1) + "/topic/" + category + "/"
	mdname := GetList(categorypath)
	//遍历每个md的名字
	for _, mname := range mdname {
		//获得创建时间
		Createdtime = strings.Split(mname, "#")[0]
		//获得文章标题
		Title = strings.Split(strings.Split(mname, "#")[1], ".")[0]
		//获得每个md的内容
		Content = readfile(categorypath + "/" + mname)
		cates := Topic{Title:Title, Createdtime:Createdtime, Category:category, Content:Content}
		cate = append(cate, cates)
	}
	return cate, err
}

//获得所有分类名字
func GetCategoryName() (string,error) {
	//获取当前路径的上一级路径
	parent, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Error(err)
	}
	//C:/GOPATH/src/myblog/topic/
	categorypath := strings.Replace(parent, "\\", "/", -1) + "/topic/"
	categorynames := GetList(categorypath)

	//自己构造html
	html := ""
	for _, n := range categorynames {
		//<li><a href="/category?category={{.Category}}">{{.Category}}</a></li>
		html = html + "<li><a href=\"/category?category=" + n + "\">" + n + "</a></li>"
	}
	return html, err
}

//首页的获取所有文章
func GetAllTopicsTopic() ([]Topic, error) {

	var Title string
	var Createdtime string
	var Content string
	var Category string

	var cate []Topic
	//获取当前路径的上一级路径
	parent, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Error(err)
	}
	//C:/GOPATH/src/myblog/topic/
	filetopic := strings.Replace(parent, "\\", "/", -1) + "/topic/"
	//获取所有分类--[百哥]
	categories := GetList(filetopic)
	for _, n := range categories {

		//获得分类
		Category = n
		//每个分类的文件夹目录为
		//C:/GOPATH/src/myblog/topic/百哥
		catefilepath := filetopic + n
		//打开每个分类获取里面的md信息
		mdname := GetList(catefilepath)
		//遍历每个md的名字
		for _, mname := range mdname {
			//获得创建时间
			Createdtime = strings.Split(mname, "#")[0]
			//获得文章标题
			Title = strings.Split(strings.Split(mname, "#")[1], ".")[0]
			//获得每个md的内容
			Content = readfile(catefilepath + "/" + mname)
			cates := Topic{Title:Title, Createdtime:Createdtime, Category:Category, Content:Content}
			cate = append(cate, cates)
		}
	}
	return cate, err
}