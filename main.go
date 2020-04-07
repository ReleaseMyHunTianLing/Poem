package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"

	"github.com/db"
	"github.com/PuerkitoBio/goquery"
)

const (
	rootPage = "https://so.gushiwen.org"
	basePage = "https://so.gushiwen.org/authors"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
)

func GetPome (url string) {
	request, err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil{
		panic(err)
	}

	request.Header.Add("user-agent",userAgent)
	request.Header.Add("Content-Type","text/plain;charset=UTF-8")

	res, err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error : %d %s",res.StatusCode,res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	doc.Find(".cont").											//古诗页
		Each(func(i int,s *goquery.Selection){
			title := ""
			authorAndchaodai := ""
			content := ""

			title = s.Find("p").Find("a").Find("b").Text()							//解析题目
			authorAndchaodai = s.Find(".source").Text()								//解析作者和朝代
			s.Find(".contson").Each(func (k int,s *goquery.Selection){             //解析出古诗文本
				content = s.Text()
				content = strings.Replace(content," ","",-1)					  //去除空格
				content = strings.Replace(content,"\n","",-1)                     //去除换行符
				//fmt.Printf("内容：%s\n",content)
			})
			if title!="" && authorAndchaodai!="" && content!="" {
				p := db.Pome{}
				p.Content = content
				p.Title = title
				p.AuthorAndchaodai = authorAndchaodai

				p.Save()                                                           //保存、输出Peom
			}
		})

}

func ParseUrl (url string, page int) []string {
	urls := make([]string,0)
	urlEach := strings.Replace(url,"A1.aspx","A%d.aspx",1)

	for i:= 1;i <= page;i++ {
		urls = append(urls,fmt.Sprintf(urlEach,i))
	}

	//fmt.Println(urls)
	return urls
}

func GetAuthorsPome(link string) {
	var PomePage string
	PomePage = rootPage + link

	request, err := http.NewRequest(http.MethodGet,PomePage,nil)
	if err != nil{
		panic(err)
	}

	request.Header.Add("user-agent",userAgent)
	request.Header.Add("Content-Type","text/plain;charset=UTF-8")

	res, err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error : %d %s",res.StatusCode,res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	doc.Find(".sonspic").Find(".cont").Find("p").Find("a").
		Each(func(i int,s *goquery.Selection){
			pLink, _ := s.Attr("href")
			fmt.Printf("作者作品：%s\n",rootPage+pLink)
			strs := ParseUrl(rootPage+pLink,10)                       //网站可以访问最多十页古诗
			for _, str := range strs {                                     
				GetPome(str)
			}
		})

}

func main() {
	db.Init()

	request, err := http.NewRequest(http.MethodGet,basePage,nil)
	if err != nil{
		panic(err)
	}

	request.Header.Add("user-agent",userAgent)
	request.Header.Add("Content-Type","text/plain;charset=UTF-8")

	res, err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error : %d %s",res.StatusCode,res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	
	doc.Find(".sons").Find(".cont").Find("a").                            //作者页
		Each(func(i int,s *goquery.Selection){
			authors := s.Text()
			fmt.Printf("%d authors= %s\n",i,authors)
			link ,_:= s.Attr("href")

			GetAuthorsPome(link)
	})
	var a int
	fmt.Scanf("%d",&a)
}