package random

import (
	"bytes"
	"math/rand"
	"sync"
	"time"
)

type RandomTool struct{}

// 获取区间内的随机数
func (s *RandomTool) GetRandomInt(start int, end int) int {
	if start == end {
		return start
	}
	a := int(start)
	e := int(end)
	rand.Seed(time.Now().UnixNano())
	g := e - a
	r := rand.Intn(g) + a
	return r
}

var (
	randSeek = int64(1)
	l        sync.Mutex
)

func (s *RandomTool) GetRandomString(num int, str ...string) string {
	ss := "0123456789"
	if len(str) > 0 {
		ss = str[0]
	}
	l := len(ss)
	r := rand.New(rand.NewSource(getRandSeek()))
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		x := r.Intn(l)
		buf.WriteString(ss[x : x+1])
	}
	return buf.String()
}

func (s *RandomTool) GetRandomChars(num int) string {
	ss := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return s.GetRandomString(num, ss)
}

func getRandSeek() int64 {
	l.Lock()
	if randSeek >= 100000000 {
		randSeek = 1
	}
	randSeek++
	l.Unlock()
	return time.Now().UnixNano() + randSeek
}
