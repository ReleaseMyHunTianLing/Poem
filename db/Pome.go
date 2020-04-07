package db

import (
	"fmt"
	"encoding/json"
)

type Pome struct {
	ID int                  
	Title string			
	AuthorAndchaodai string			
	Content string
}

func (p *Pome) Insert() bool{
	stmt, err := db.Prepare("INSERT into poem (title,dynastyAndauthor,content) values (?,?,?)")
	if err != nil{
		return false
	}
	_,err = stmt.Exec(&p.Title,&p.AuthorAndchaodai,&p.Content)
	if err != nil {
		return false
	}
	return true
}

func (p *Pome) Save() {
	data, err := json.Marshal(p)
	if err !=nil {
		fmt.Printf("生成json出错\n")
	}

	res := p.Insert()                            //向数据库插入记录项
	if res==false {
		fmt.Printf("insert failed\n")
	}
	fmt.Println(string(data))
}