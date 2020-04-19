package main

import (
	"encoding/json"
	"fmt"
)

var a = `{"code":200,"msg":"success","data":{"recommend_card_num":0,"recommend_card_expert":"","count":{"all":33,"enable":30,"draft":0,"deleted":0,"private":0,"audit":3,"original":0},"list":[{"ArticleId":"105595102","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(10)\u57fa\u7840\u77e5\u8bc6-\u7b97\u6cd5\u590d\u6742\u5ea6\u4e3b\u65b9\u6cd5","PostTime":"2020\u5e7404\u670818\u65e5 11:01:27","ViewCount":"5","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105595090","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(9)\u57fa\u7840\u77e5\u8bc6-\u7b97\u6cd5\u590d\u6742\u5ea6\u53ca\u6e10\u8fdb\u7b26\u53f7","PostTime":"2020\u5e7404\u670818\u65e5 11:00:56","ViewCount":"5","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105595077","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(8.2)\u57fa\u7840\u77e5\u8bc6-\u5206\u6cbb\u6cd5\u548c\u9012\u5f52","PostTime":"2020\u5e7404\u670818\u65e5 11:00:09","ViewCount":"6","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105595061","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(8.1)\u57fa\u7840\u77e5\u8bc6-\u524d\u8a00","PostTime":"2020\u5e7404\u670818\u65e5 10:59:33","ViewCount":"8","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105595033","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(7)\u7b80\u5355\u5165\u95e8Golang-\u6807\u51c6\u5e93","PostTime":"2020\u5e7404\u670818\u65e5 10:58:56","ViewCount":"7","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105563314","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(6)\u7b80\u5355\u5165\u95e8Golang-\u5e76\u53d1\u3001\u534f\u7a0b\u548c\u4fe1\u9053","PostTime":"2020\u5e7404\u670816\u65e5 17:39:19","ViewCount":"19","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105563303","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(5)\u7b80\u5355\u5165\u95e8Golang-\u63a5\u53e3","PostTime":"2020\u5e7404\u670816\u65e5 17:38:39","ViewCount":"23","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105563289","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(4)\u7b80\u5355\u5165\u95e8Golang-\u7ed3\u6784\u4f53\u548c\u65b9\u6cd5","PostTime":"2020\u5e7404\u670816\u65e5 17:38:06","ViewCount":"21","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105563270","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(3)\u7b80\u5355\u5165\u95e8Golang-\u6d41\u7a0b\u63a7\u5236\u8bed\u53e5","PostTime":"2020\u5e7404\u670816\u65e5 17:37:22","ViewCount":"7","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105563253","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(2)\u7b80\u5355\u5165\u95e8Golang-\u5305\u3001\u53d8\u91cf\u548c\u51fd\u6570","PostTime":"2020\u5e7404\u670816\u65e5 17:36:30","ViewCount":"21","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105563142","Title":"\u6570\u636e\u7ed3\u6784\u548c\u7b97\u6cd5(Golang\u5b9e\u73b0)(1)\u7b80\u5355\u5165\u95e8Golang-\u524d\u8a00","PostTime":"2020\u5e7404\u670816\u65e5 17:34:30","ViewCount":"18","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":1,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105563501","Title":"Golang1.5\u5230Golang1.12\u5305\u7ba1\u7406\uff1agolang vendor \u5230 go mod","PostTime":"2019\u5e7403\u670805\u65e5 00:00:00","ViewCount":"11","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":0,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"},{"ArticleId":"105563315","Title":"\u516b\u7687\u540e\uff0c\u56de\u6eaf\u4e0e\u9012\u5f52\uff08Python\u5b9e\u73b0\uff09","PostTime":"2019\u5e7403\u670805\u65e5 00:00:00","ViewCount":"3","CommentCount":"0","CommentAuth":"2","IsTop":"0","Status":"1","UserName":"m0_46803965","Type":"1","is_vip_article":false,"editor_type":0,"is_recommend":false,"title_repeat_num":0,"collect_count":"0"}],"total":33,"list_status":"all","page":2,"size":20}}`
type XX struct {
	Data XX1 `json:"data"`
}

type XX1 struct {
	L []XX2 `json:"list"`
}

type XX2 struct {
	Title     string `json:"Title"`
	UserName  string `json:"UserName"`
	ArticleId string `json:"ArticleId"`
}

func main() {
	x := new(XX)
	err := json.Unmarshal([]byte(a), x)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	l := len(x.Data.L)
	for i := l - 1; i >= 0; i-- {
		v := x.Data.L[i]
		fmt.Printf("- [%s](https://blog.csdn.net/%s/article/details/%s)\n", v.Title, v.UserName, v.ArticleId)
	}
}
