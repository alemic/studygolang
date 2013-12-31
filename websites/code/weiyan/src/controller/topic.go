// Copyright 2013 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	studygolang@gmail.com

package controller

import (
	"bytes"
	"filter"
	"fmt"
	"github.com/studygolang/mux"
	"html/template"
	"logger"
	"math/rand"
	"model"
	"net/http"
	"service"
	"strconv"
	"strings"
	"time"
	"util"
)

func randomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// 在需要评论且要回调的地方注册评论对象
func init() {
	// 注册评论对象
	service.RegisterCommentObject("topic", service.TopicComment{})
}

func isMobile(req *http.Request) (ismobile bool) {
	ismobile = false
	if strings.Contains(req.Header["User-Agent"][0], "Android") {
		ismobile = true
	}
	return
}

func TopicsBriefHandler(rw http.ResponseWriter, req *http.Request) {
	ismobile := isMobile(req)
	fmt.Println("ismobile is", ismobile)
	TopicsHandler(rw, req, ismobile)
	//ismobile = !ismobile
	//nodes["ismobile"] = ismobile
}

func TopicsListHandler(rw http.ResponseWriter, req *http.Request) {
	//ismobile := isMobile(req)
	ismobile := true
	fmt.Println("ismobile is", ismobile)

	TopicsHandler(rw, req, ismobile)
	//ismobile = !ismobile
	//nodes["ismobile"] = ismobile
}

// 社区帖子列表页
// uri: /topics{view:(|/popular|/no_reply|/last)}
func TopicsHandler(rw http.ResponseWriter, req *http.Request, ismobile bool) {
	fmt.Println("User-Agent")
	fmt.Println(req.Header["User-Agent"][0])

	nodes := genNodes()

	// 设置内容模板
	page, _ := strconv.Atoi(req.FormValue("p"))
	if page == 0 {
		page = 1
	}
	vars := mux.Vars(req)
	order := ""
	where := ""
	switch vars["view"] {
	case "/popular":
		where = "like>0"
	case "/last":
		order = "ctime DESC"
	}

	var PAGE_NUM int
	if ismobile {
		PAGE_NUM = 30
	} else {
		PAGE_NUM = 30
	}

	topics, total := service.FindTopics(page, PAGE_NUM, where, order)
	fmt.Println(total)
	fmt.Println(len(topics))

	fmt.Println(PAGE_NUM)
	pageHtml := service.GetPageHtml(page, total, PAGE_NUM)
	if ismobile {
		req.Form.Set(filter.CONTENT_TPL_KEY, "/template/topics/list_mobile.html")
	} else {
		req.Form.Set(filter.CONTENT_TPL_KEY, "/template/topics/list.html")
	}

	// 设置模板数据
	switch vars["view"] {
	case "/popular":
		filter.SetData(req, map[string]interface{}{"popular": 1, "topics": topics, "page": template.HTML(pageHtml), "nodes": nodes})
	case "/last":
		filter.SetData(req, map[string]interface{}{"last": 1, "topics": topics, "page": template.HTML(pageHtml), "nodes": nodes})
	default:
		filter.SetData(req, map[string]interface{}{"active": 1, "topics": topics, "page": template.HTML(pageHtml), "nodes": nodes})
	}
}

// 某节点下的帖子列表
// uri: /topics/node{nid:[0-9]+}
func NodesHandler(rw http.ResponseWriter, req *http.Request) {
	page, _ := strconv.Atoi(req.FormValue("p"))
	if page == 0 {
		page = 1
	}
	vars := mux.Vars(req)
	topics, total := service.FindTopics(page, 0, "nid="+vars["nid"])
	pageHtml := service.GetPageHtml(page, total, 30)
	// 当前节点信息
	node := model.GetNode(util.MustInt(vars["nid"]))
	req.Form.Set(filter.CONTENT_TPL_KEY, "/template/topics/node.html")
	// 设置模板数据
	filter.SetData(req, map[string]interface{}{"activeTopics": "active", "topics": topics, "page": template.HTML(pageHtml), "total": total, "node": node})
}

// 社区帖子详细页
// uri: /topics/{tid:[0-9]+}
func TopicDetailHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uid := 0
	user, ok := filter.CurrentUser(req)
	if ok {
		uid = user["uid"].(int)
	}
	// TODO:刷屏暂时不处理

	topic, replies, err := service.FindTopicByTid(vars["tid"])
	if err != nil || topic == nil || topic["tid"] == nil {
		fmt.Println("------")
		fmt.Println(vars["tid"])
		i, _ := strconv.Atoi(vars["tid"])
		total := service.TopicsTotal()
		if i >= total {
			i = 0
		}
		if i <= 0 {
			i = total - 1
		}

		for ; i <= total; i++ {
			fmt.Println(i)
			topic, replies, err = service.FindTopicByTid(strconv.Itoa(i))
			if err == nil && topic != nil && topic["tid"] != nil {
				break
			}
		}
	}
	fmt.Println("------end..........")
	if err != nil || topic == nil || topic["tid"] == nil {
		NotFoundHandler(rw, req)
		return
	}

	// 增加浏览量
	service.IncrTopicView(vars["tid"], uid)

	topic["prev_tid"] = topic["tid"].(int) - 1
	topic["next_tid"] = topic["tid"].(int) + 1
	// 设置内容模板
	req.Form.Set(filter.CONTENT_TPL_KEY, "/template/topics/detail.html")
	// 设置模板数据
	filter.SetData(req, map[string]interface{}{"activeTopics": "active", "topic": topic, "replies": replies})
}

// 社区帖子详细页
// uri: /topics/{tid:[0-9]+}
func InspectTopicHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)

	vars := mux.Vars(req)
	uid, _ := strconv.Atoi(Config["auid"])
	user, ok := filter.CurrentUser(req)
	if ok {
		uid = user["uid"].(int)
	}

	fmt.Println("------------")
	fmt.Println(vars["tid"])
	fmt.Println(vars["val"])
	fmt.Println("------------")
	total := service.TopicsTotal()
	if vars["tid"] == "" {
		vars["tid"] = strconv.Itoa(randInt(1, total))
	} else {
		tid, _ := strconv.Atoi(vars["tid"])

		client := strings.Split(req.RemoteAddr, ":")

		fmt.Println(tid, uid, client[0])
		if service.InsertVote(tid, uid, client[0]) == true {
			valInt, err := strconv.Atoi(vars["val"])
			fmt.Println(valInt)
			if err != nil {
				fmt.Fprint(rw, `{"errno": 1, "error": "invalid request data"}`)
				return
			}

			if valInt == 0 {
				service.IncrTopicHate(vars["tid"], uid)
			} else {
				// 增加like量
				service.IncrTopicLike(vars["tid"], uid)
			}
		}
	}
	// TODO:刷屏暂时不处理
	i, _ := strconv.Atoi(vars["tid"])
	i++
	if i >= total {
		i = 0
	} else if i <= 0 {
		i = total - 1
	}

	topic, _, err := service.FindTopicByTid(strconv.Itoa(i))
	if err != nil || topic == nil || topic["tid"] == nil {
		for ; i <= total; i++ {
			fmt.Println(i)
			topic, _, err = service.FindTopicByTid(strconv.Itoa(i))
			if err == nil && topic != nil && topic["tid"] != nil {
				break
			}
		}
	}

	fmt.Println("------end..........")
	if err != nil || topic == nil || topic["tid"] == nil {
		NotFoundHandler(rw, req)
		return
	}

	// 增加浏览量
	service.IncrTopicView(vars["tid"], uid)

	// 设置内容模板
	req.Form.Set(filter.CONTENT_TPL_KEY, "/template/topics/inspect.html")
	// 设置模板数据
	filter.SetData(req, map[string]interface{}{"activeTopics": "active", "topic": topic})
}

// 新建帖子
// uri: /topics/new{json:(|.json)}
func NewTopicHandler(rw http.ResponseWriter, req *http.Request) {
	nodes := genNodes()
	vars := mux.Vars(req)
	content := req.FormValue("content")
	// 请求新建帖子页面
	if content == "" || req.Method != "POST" || vars["json"] == "" {
		req.Form.Set(filter.CONTENT_TPL_KEY, "/template/topics/new.html")
		filter.SetData(req, map[string]interface{}{"nodes": nodes})
		return
	}

	// 入库
	topic := model.NewTopic()
	fmt.Println("anonymous")
	fmt.Println(req.FormValue("anonymous"))
	if req.FormValue("anonymous") == "1" {
		topic.Uid, _ = strconv.Atoi(Config["auid"])
	} else {
		user, _ := filter.CurrentUser(req)
		if user != nil {
			topic.Uid = user["uid"].(int)
		} else {
			topic.Uid, _ = strconv.Atoi(Config["auid"])
		}
	}

	fmt.Println(topic)
	topic.Nid = util.MustInt(req.FormValue("nid"))
	topic.Title = req.FormValue("title")
	topic.Content = req.FormValue("content")
	fmt.Println(topic)
	errMsg, err := service.PublishTopic(topic)
	fmt.Println("PublishTopic end")
	fmt.Println(errMsg)
	if err != nil {
		fmt.Fprint(rw, `{"errno": 1, "error":"`, errMsg, `"}`)
		return
	}
	fmt.Fprint(rw, `{"errno": 0, "error":""}`)
}

// 将node组织成一定结构，方便前端展示
func genNodes() []map[string][]map[string]interface{} {
	sameParent := make(map[string][]map[string]interface{})
	allParentNodes := make([]string, 0)
	for _, node := range model.AllNode {
		if node["pid"].(int) != 0 {
			if len(sameParent[node["parent"].(string)]) == 0 {
				sameParent[node["parent"].(string)] = []map[string]interface{}{node}
			} else {
				sameParent[node["parent"].(string)] = append(sameParent[node["parent"].(string)], node)
			}
		} else {
			allParentNodes = append(allParentNodes, node["name"].(string))
		}
	}
	nodes := make([]map[string][]map[string]interface{}, 0)
	for _, parent := range allParentNodes {
		tmpMap := make(map[string][]map[string]interface{})
		tmpMap[parent] = sameParent[parent]
		nodes = append(nodes, tmpMap)
	}
	logger.Debugf("%v\n", nodes)
	return nodes
}
