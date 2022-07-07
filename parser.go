package json_parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Msg struct {
	ID         string `json:"id"`
	Type       int64  `json:"type"`
	Content    string `json:"content"`
	Createtime int64  `json:"create_time"`
	HasRead    bool   `json:"has_read"`
}

func Parse(path string) []*Msg {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("read file err=%v \n", err)
		return nil
	}

	elements := []*Msg{}

	err = json.Unmarshal(bytes, &elements)
	if err != nil {
		log.Printf("unmarshal json err=%v \n", err)
		return nil
	}
	return elements
}

func StreamParse(path string) []*Msg {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("open file err:%v \n", err)
		return nil
	}

	decoder := json.NewDecoder(file)
	t, err := decoder.Token()
	if err != nil {
		log.Printf("get json token failed, err=%v \n", err)
		return nil
	}
	log.Printf("start token:%v", t)

	msgList := make([]*Msg, 0)
	count := 0
	for decoder.More() {
		var msg Msg
		err = decoder.Decode(&msg)
		if err != nil {
			log.Printf("decod err:%v", err)
			return nil
		}
		msgList = append(msgList, &msg)
		count++
	}

	t, err = decoder.Token()
	if err != nil {
		log.Printf("get json token failed, err=%v \n", err)
		return nil
	}
	log.Printf("end token:%v", t)
	log.Printf("count:%d", count)

	return msgList
}
