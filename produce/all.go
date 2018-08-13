package main

import (
	"math/rand"
	"time"
	"fmt"
	"encoding/json"
)

/* 如何将100随机的分成30份?
	1.取 29 个(0, 100)的坐标
	2.随便生成 30 个随机数 按照总和与 100 的比例 所有数字等比例缩放
	3.微信红包的算法
*/
type elem struct {
	Gold int `json:"gold"`
	Wood int `json:"wood"`
	Lake int `json:"lake"`
	Fire int `json:"fire"`
	Earth int `json:"earth"`
}
type coordinate struct {
	x int
	y int
}



var elems []elem = []elem{{0,0,100,0,0},{0,0,100,0,0},{0,100,0,100,0},{0,100,0,0,0},{0,0,100,0,100},{0,0,0,100,100},{100,0,0,100,0},{0,0,100,0,0},{0,100,0,0,100},{0,0,0,0,100},{0,100,0,0,0},{0,0,0,100,133},{0,0,100,0,133},{0,0,0,0,133},{0,0,100,100,133},{0,133,100,100,0},{0,0,100,0,0},{0,0,100,100,0},{0,0,0,100,100},{0,0,100,0,0},{0,0,100,0,100},{0,100,0,0,0},{0,100,0,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,100},{100,0,0,0,100},{100,0,0,100,0},{0,0,100,100,0},{100,0,100,0,0},{0,0,100,0,0},{0,100,0,0,0},{0,0,0,100,0},{0,0,100,0,100},{0,100,0,0,0},{100,0,0,0,100},{0,0,133,0,0},{0,0,100,0,100},{0,0,0,100,0},{100,0,0,100,0},{100,0,0,100,0},{0,100,100,0,0},{0,0,0,0,100},{0,0,0,0,100},{100,0,100,0,0},{0,100,0,0,100},{100,100,0,0,0},{0,0,100,0,0},{100,0,0,0,0},{0,100,0,0,100},{0,0,0,0,100},{0,100,0,100,0},{0,0,0,0,100},{0,0,100,0,133},{0,0,100,0,135},{0,0,100,100,144},{100,0,0,0,150},{0,0,0,0,150},{0,135,0,100,150},{0,144,0,0,100},{0,150,0,0,144},{0,144,0,100,135},{100,135,0,100,0},{0,0,0,100,0},{100,0,0,0,100},{0,0,100,0,100},{0,0,100,0,100},{0,100,100,0,0},{100,0,0,0,0},{0,0,100,0,100},{100,0,0,0,100},{100,0,0,0,0},{0,0,100,0,0},{0,0,0,0,100},{0,0,100,100,0},{0,0,100,0,100},{0,0,0,100,0},{0,100,0,0,100},{100,133,0,0,100},{0,133,135,0,0},{100,0,144,0,0},{0,0,150,100,0},{0,0,144,100,100},{100,0,135,0,100},{0,100,0,0,0},{100,0,0,100,0},{100,0,0,100,0},{100,0,0,0,0},{100,0,0,0,0},{100,100,0,0,0},{100,0,0,0,100},{0,100,0,0,0},{0,0,0,100,0},{0,0,100,0,0},{0,100,0,0,0},{0,0,0,0,100},{0,0,0,0,135},{100,0,0,100,144},{100,0,100,0,150},{0,0,0,100,150},{100,0,0,100,170},{0,133,0,0,100},{0,135,100,0,200},{100,144,0,0,200},{0,170,0,0,200},{0,200,100,0,170},{0,170,0,100,144},{100,150,100,0,0},{0,144,100,100,0},{100,135,0,0,0},{0,0,0,0,100},{0,0,100,0,0},{0,100,0,0,100},{100,0,0,0,0},{0,100,0,0,0},{100,0,0,0,0},{0,0,0,100,0},{100,0,0,100,0},{0,100,0,0,0},{0,100,100,0,0},{0,0,0,0,100},{100,135,100,0,0},{100,144,0,100,0},{0,150,135,0,0},{0,150,144,0,100},{0,144,100,0,0},{0,135,200,0,100},{0,0,170,0,100},{0,0,144,0,100},{0,0,135,100,0},{0,0,0,100,100},{0,100,0,0,100},{100,0,100,0,0},{0,0,100,100,0},{100,0,0,100,0},{100,0,0,0,0},{100,100,0,0,0},{0,100,0,0,0},{0,0,0,100,100},{0,0,0,100,135},{0,0,0,100,144},{100,0,0,0,150},{0,0,100,0,170},{0,0,100,100,200},{100,135,0,0,200},{0,144,0,0,200},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,400,0,0,0},{0,200,100,0,150},{0,200,100,0,133},{100,170,0,100,0},{0,144,100,0,100},{0,0,0,100,0},{100,0,0,0,0},{0,100,0,0,100},{0,0,0,0,100},{100,0,0,100,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,100,0,100},{0,0,100,100,0},{0,0,100,0,0},{100,135,0,100,0},{100,144,133,0,0},{0,170,133,0,100},{0,200,144,0,133},{0,200,170,0,133},{0,170,200,0,133},{0,0,400,0,0},{0,135,200,0,133},{100,0,170,0,0},{100,0,144,100,0},{100,0,100,0,0},{0,0,0,100,0},{0,100,100,0,0},{0,100,0,0,0},{0,100,0,0,100},{100,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,100,0,135},{0,0,100,100,144},{100,0,0,0,170},{100,0,0,100,200},{0,0,100,100,200},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,100,0,0},{100,150,0,100,0},{0,133,100,0,100},{0,100,0,0,0},{0,100,0,0,100},{0,0,0,100,0},{100,0,100,0,0},{0,0,100,0,0},{100,0,0,100,0},{0,0,0,0,100},{100,0,0,100,0},{0,0,135,0,100},{0,144,144,0,0},{0,170,150,0,135},{0,200,150,0,144},{0,400,0,0,0},{0,400,0,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,200,0,144},{0,0,150,100,135},{0,0,133,0,0},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,100,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,0,100,0},{100,0,0,0,135},{100,0,0,0,144},{0,0,0,100,170},{0,0,100,100,200},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,0,100},{100,200,0,100,0},{0,170,100,100,0},{0,144,100,0,100},{0,0,0,100,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,100,100,0},{0,0,0,100,0},{0,133,144,0,100},{0,150,170,0,100},{0,200,200,0,144},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,133,200,0,170},{0,0,150,0,144},{0,0,133,100,100},{0,0,100,0,0},{0,0,0,100,100},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,100,0},{0,0,0,100,0},{0,0,0,100,144},{0,0,0,100,170},{0,0,0,100,200},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,150,0,0,100},{0,200,0,0,200},{0,400,0,0,0},{0,400,0,0,0},{0,200,100,0,150},{0,200,0,100,150},{0,200,0,0,144},{0,200,0,100,135},{0,170,0,100,0},{100,150,0,0,100},{0,144,100,0,0},{100,135,0,0,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,100,0,100},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,0,0,100},{100,0,0,0,100},{100,0,133,0,100},{100,0,150,0,100},{0,144,200,0,133},{0,0,400,0,0},{0,0,400,0,0},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,100,0,200},{0,0,144,0,150},{0,0,100,0,133},{0,0,0,0,100},{0,0,0,0,100},{0,0,0,0,100},{100,0,0,100,0},{100,0,0,100,0},{0,0,100,0,133},{100,0,0,0,150},{0,0,100,100,200},{0,0,0,0,400},{0,0,0,0,400},{100,0,0,100,200},{100,0,0,100,200},{0,144,0,100,170},{100,170,0,0,150},{100,200,0,0,150},{100,200,0,0,144},{0,170,100,0,135},{100,150,0,0,133},{0,150,100,0,100},{0,150,100,0,100},{0,144,100,0,0},{0,135,100,0,0},{100,0,0,100,0},{100,0,0,100,0},{100,0,0,0,100},{0,0,100,0,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,100,100},{100,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,133,0,0},{100,0,150,0,0},{0,135,100,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{0,0,0,0,400},{100,0,144,0,200},{0,0,100,0,150},{100,0,0,0,133},{100,0,0,0,0},{0,0,0,0,100},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,0,100},{0,0,0,100,133},{0,0,100,0,150},{100,0,0,100,200},{0,0,0,0,400},{0,0,0,0,200},{0,0,100,0,170},{100,0,0,100,150},{100,135,0,0,144},{100,144,0,0,135},{100,150,0,0,133},{100,150,0,0,100},{0,144,0,0,100},{100,135,0,100,0},{0,133,100,0,100},{0,133,100,0,100},{0,0,0,0,100},{0,0,0,100,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,0,100},{100,0,0,100,0},{0,0,0,0,100},{0,0,0,100,100},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,100,0},{0,0,144,100,100},{0,0,170,0,0},{0,0,100,0,135},{0,0,400,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,0,0,400},{0,0,0,0,400},{100,0,144,0,200},{100,0,0,0,170},{0,0,100,0,144},{100,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,100},{0,0,0,100,100},{0,0,0,0,100},{0,0,0,100,0},{0,0,0,0,144},{0,0,0,100,170},{100,0,100,0,200},{0,0,0,100,170},{0,0,0,100,144},{0,0,100,0,135},{0,0,0,0,100},{0,0,100,0,0},{0,133,100,0,100},{0,133,0,100,100},{0,0,100,100,0},{0,0,0,100,100},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,100},{0,0,0,0,100},{0,0,0,0,100},{0,0,100,0,0},{100,0,0,0,0},{0,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,0,0},{100,0,0,100,0},{100,0,100,0,0},{100,0,0,100,0},{100,0,135,100,0},{100,0,144,0,100},{100,0,170,0,0},{0,0,200,100,135},{0,0,400,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,200,0,200},{0,0,150,0,170},{100,0,133,0,144},{0,0,100,0,135},{100,0,0,0,100},{100,0,100,0,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,100,0,135},{0,0,100,100,144},{0,0,100,0,150},{100,0,0,0,144},{100,0,0,100,135},{0,0,100,0,0},{0,0,0,0,100},{0,0,0,0,100},{100,0,100,0,0},{0,0,0,100,100},{100,0,0,0,0},{0,0,133,100,0},{0,0,133,100,100},{100,0,0,0,0},{0,0,100,0,0},{0,0,0,100,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,100,100,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,100,0,100},{0,0,100,0,0},{0,0,0,100,100},{100,0,100,0,0},{0,0,0,100,100},{0,0,100,0,0},{100,0,0,0,0},{0,0,135,100,100},{0,0,144,0,100},{0,0,170,100,0},{0,0,200,0,100},{0,0,100,0,144},{0,0,200,100,150},{100,0,170,0,150},{100,0,144,0,144},{0,0,100,0,135},{0,0,0,0,100},{0,0,0,100,100},{100,0,0,0,0},{100,0,0,100,0},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,0,100,133},{100,0,0,100,0},{100,0,0,0,0},{0,0,100,100,0},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,135,0,100},{0,0,144,0,100},{0,0,150,100,0},{100,0,150,0,0},{0,0,144,0,0},{0,0,135,100,100},{100,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,100,100,0},{100,0,0,0,0},{0,0,0,100,100},{0,0,100,0,0},{0,0,100,100,0},{0,0,0,100,0},{100,0,0,0,0},{0,0,100,100,0},{0,0,0,100,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,135,0,100},{100,0,144,0,0},{0,0,150,0,100},{0,0,150,100,100},{100,0,150,0,133},{0,0,144,0,133},{100,0,135,0,0},{0,0,0,100,100},{100,0,100,0,0},{100,0,0,0,0},{0,0,0,0,100},{0,0,100,100,0},{0,0,100,100,0},{0,0,0,0,100},{0,0,0,100,0},{100,0,100,0,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,100,100},{0,0,0,0,100},{0,0,100,0,0},{100,0,0,0,100},{100,0,0,0,0},{0,0,135,0,100},{100,0,144,0,0},{0,0,170,0,100},{0,0,200,100,0},{0,0,200,100,100},{0,0,170,0,100},{100,0,144,0,0},{100,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,100,0,0},{0,0,100,100,0},{0,0,100,0,100},{0,0,0,0,100},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,100,100},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{100,0,0,100,0},{100,0,133,0,100},{100,0,133,0,0},{0,0,133,100,100},{0,0,0,100,0},{100,0,0,0,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{100,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,100,0,100},{0,0,0,0,100},{100,0,0,0,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,144,100,0},{100,0,170,100,0},{0,0,200,100,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,200,0,100},{100,0,150,100,0},{0,0,133,0,100},{0,0,0,100,0},{0,0,0,0,100},{0,0,100,0,0},{0,0,0,0,100},{0,0,100,0,0},{100,0,100,0,0},{0,0,0,100,0},{100,0,0,100,0},{100,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,100,0,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,100,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,100,100,0},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,100},{0,0,100,100,0},{0,0,0,100,100},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,0,100,100},{100,0,0,0,0},{100,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{100,0,0,100,0},{100,0,133,0,0},{0,0,150,100,100},{0,0,200,100,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,200,0,100},{0,0,150,0,100},{0,0,133,100,0},{100,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,0,0,100},{0,0,100,100,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,100,0,100},{0,0,100,0,0},{100,0,0,100,0},{100,0,0,100,0},{100,0,0,0,100},{0,0,100,0,100},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,100},{0,0,100,0,100},{0,0,100,0,0},{0,0,100,100,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,100,0,0},{0,0,0,100,100},{100,0,0,100,0},{0,0,0,0,100},{100,0,100,0,0},{100,0,0,0,0},{100,0,133,0,0},{100,0,150,0,100},{0,0,200,0,100},{0,0,400,0,0},{0,0,400,0,0},{0,0,400,0,0},{0,0,200,0,100},{100,0,150,0,0},{100,0,133,0,100},{100,0,0,0,0},{100,0,0,0,0},{0,0,0,0,100},{0,0,0,0,100},{100,0,0,0,100},{100,0,100,0,0},{100,0,0,0,0},{0,0,100,0,0},{0,0,0,100,0},{100,0,0,100,0},{100,0,0,0,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,100,100,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,100,100,0},{100,0,0,0,100},{0,0,0,100,0},{100,0,0,100,0},{133,0,100,0,0},{0,0,0,100,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,100},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,100,0,100},{100,0,0,0,100},{100,0,0,0,100},{100,0,0,0,100},{0,0,0,100,0},{0,0,144,0,100},{100,0,170,0,0},{100,0,200,0,100},{0,0,400,0,0},{100,0,200,0,0},{0,0,170,100,0},{0,0,144,0,100},{0,0,0,100,0},{100,0,100,0,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,100,100,0},{100,0,0,0,100},{0,0,100,0,100},{0,0,100,0,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,0,100,0},{100,0,100,0,0},{135,0,100,0,100},{144,0,0,0,0},{150,0,100,0,0},{144,0,100,0,0},{135,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,100,100,0},{100,0,0,100,0},{100,0,0,100,0},{0,0,0,0,100},{100,0,0,100,0},{0,0,100,100,0},{100,0,0,100,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,100,0,0},{0,0,100,0,0},{0,0,135,100,0},{0,0,144,0,100},{0,0,170,100,100},{0,0,200,100,0},{100,0,170,0,0},{0,0,144,100,0},{0,0,135,0,100},{100,0,100,0,0},{100,0,0,0,0},{0,0,100,0,100},{0,0,100,0,0},{100,0,0,100,0},{100,0,0,0,100},{0,0,100,100,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,100,0,100},{100,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{135,0,0,0,0},{144,0,0,100,0},{170,0,100,0,0},{200,0,0,0,100},{170,0,0,0,100},{144,0,0,100,100},{135,133,0,0,100},{0,0,0,100,0},{0,0,100,0,100},{0,0,0,0,100},{0,0,0,0,100},{100,0,0,0,100},{0,0,0,100,0},{0,0,100,0,100},{0,0,0,100,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,0,0,100},{0,0,100,0,0},{0,0,0,0,100},{100,0,0,0,0},{100,0,135,0,0},{0,0,144,100,100},{0,0,150,0,100},{0,0,144,100,0},{100,0,135,0,0},{0,0,100,0,0},{0,0,0,100,100},{100,200,100,0,0},{0,100,100,200,0},{200,100,100,0,0},{0,100,200,0,100},{100,200,0,0,100},{100,100,200,0,0},{0,100,100,0,200},{0,100,0,200,100},{100,100,200,0,0},{100,0,0,100,0},{0,0,100,0,0},{0,0,0,100,100},{0,0,100,0,0},{0,0,100,100,0},{0,0,100,0,100},{100,0,0,100,0},{0,0,100,0,100},{144,0,100,0,0},{170,0,100,0,100},{200,0,0,100,0},{400,0,0,0,0},{200,135,0,0,100},{170,144,100,0,0},{144,150,0,100,0},{135,144,0,0,0},{100,135,100,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,0,100},{0,0,100,100,0},{100,0,0,0,100},{0,0,0,0,100},{100,0,0,100,0},{0,0,100,100,0},{0,0,0,0,100},{0,0,0,0,100},{0,0,100,0,0},{0,0,0,100,0},{100,0,133,0,0},{0,0,100,100,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,200,100},{200,0,100,0,100},{100,0,100,200,0},{0,0,100,200,100},{100,200,100,0,0},{0,100,100,200,0},{100,100,0,200,0},{100,100,0,0,200},{100,100,0,200,0},{100,0,0,0,0},{100,0,0,100,0},{0,0,100,0,100},{0,0,0,0,100},{100,0,100,0,0},{0,0,0,0,100},{0,0,0,0,100},{133,0,0,100,100},{150,0,0,0,100},{200,0,100,0,0},{400,0,0,0,0},{400,0,0,0,0},{400,0,0,0,0},{200,170,0,0,100},{170,200,0,0,100},{144,170,100,0,0},{0,144,0,0,100},{0,0,0,100,0},{100,0,100,0,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,100,0,100},{0,0,100,0,0},{0,100,0,200,100},{100,0,0,200,100},{200,100,100,0,0},{100,200,0,0,100},{0,200,100,0,100},{100,100,0,200,0},{100,100,0,200,0},{100,100,0,200,0},{0,0,100,200,100},{100,0,100,0,0},{0,0,0,0,100},{0,0,0,0,100},{0,0,100,100,0},{100,0,0,0,0},{100,0,0,100,0},{100,0,0,100,0},{133,0,0,0,100},{150,0,0,100,0},{200,0,100,100,0},{400,0,0,0,0},{400,0,0,0,0},{400,0,0,0,0},{400,0,0,0,0},{0,400,0,0,0},{150,200,0,100,0},{133,150,0,0,0},{100,133,0,100,0},{0,0,100,0,100},{100,0,0,100,0},{100,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,100,0,100},{0,0,0,0,100},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,100,100},{0,0,100,0,100},{0,0,0,0,100},{0,0,0,100,100},{0,100,100,200,0},{100,100,0,200,0},{100,0,100,0,200},{0,100,100,200,0},{100,200,100,0,0},{0,0,100,200,100},{0,100,100,200,0},{100,200,0,0,100},{100,0,100,200,0},{0,0,0,100,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,0,100,100},{100,0,0,0,0},{0,0,0,100,0},{100,0,0,0,0},{0,0,100,0,0},{144,0,100,0,100},{170,0,0,100,0},{100,144,0,0,0},{400,0,0,0,0},{400,0,0,0,0},{0,400,0,0,0},{170,200,0,0,100},{144,170,0,0,0},{0,144,0,0,100},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,100,100,0},{0,0,0,100,0},{100,0,0,0,0},{133,0,0,0,0},{0,0,100,0,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,0,100},{0,0,0,0,100},{0,0,0,100,100},{100,0,0,100,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{100,0,100,200,0},{0,100,100,200,0},{100,0,0,200,100},{0,0,100,200,100},{100,200,0,0,100},{200,100,100,0,0},{200,0,100,0,100},{100,0,0,200,100},{0,0,100,200,100},{100,0,0,100,0},{0,0,0,100,0},{100,0,0,0,100},{0,0,100,0,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,0,100},{100,0,0,0,0},{135,0,0,100,0},{144,135,0,0,100},{170,150,0,0,0},{200,200,0,100,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{135,200,0,0,100},{0,150,100,0,100},{0,133,100,100,0},{100,0,100,0,0},{0,0,0,100,0},{100,0,0,0,0},{135,0,0,0,100},{144,0,0,0,0},{150,0,100,0,0},{144,0,0,0,100},{135,0,100,100,0},{133,0,0,100,0},{0,0,100,0,100},{0,0,0,100,100},{100,0,0,100,0},{0,0,0,100,0},{0,0,0,100,100},{0,0,0,100,0},{100,0,100,0,0},{0,0,100,0,100},{0,0,100,100,0},{100,0,100,200,0},{100,100,0,200,0},{100,100,0,0,200},{100,0,0,200,100},{100,100,200,0,0},{100,0,100,200,0},{100,100,0,200,0},{100,0,100,200,0},{200,100,0,0,100},{0,0,100,100,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,100,100},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,0,100},{0,0,100,0,0},{135,144,0,0,0},{144,170,100,0,0},{150,200,133,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,0,100},{100,150,100,0,0},{0,133,0,100,100},{100,0,0,100,0},{0,0,0,100,0},{135,0,0,100,100},{144,0,100,100,0},{170,0,100,0,100},{200,0,100,0,100},{170,0,0,100,0},{144,0,0,0,0},{150,0,100,0,100},{144,0,0,0,100},{135,0,0,0,0},{0,0,100,0,100},{0,0,100,0,0},{0,0,0,0,100},{100,0,0,100,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,0,100,0},{100,100,0,200,0},{100,0,200,0,100},{0,100,0,200,100},{100,0,100,200,0},{0,100,0,200,100},{200,0,100,0,100},{100,0,200,0,100},{100,0,0,200,100},{100,100,200,0,0},{100,0,100,0,0},{0,0,100,100,0},{0,0,0,100,0},{100,0,0,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,100},{100,133,0,0,0},{0,150,135,0,100},{0,200,100,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,100,0,100},{100,170,100,0,0},{100,144,0,0,100},{100,0,0,0,0},{0,0,0,100,0},{0,0,0,0,100},{144,0,100,0,100},{170,0,0,100,0},{200,0,0,100,0},{400,0,0,0,0},{200,0,100,0,100},{170,0,0,0,100},{200,0,0,0,100},{170,0,0,100,0},{144,0,100,0,0},{100,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,0,100,0},{100,0,0,100,0},{0,0,0,0,100},{0,0,0,100,100},{100,100,0,200,0},{0,100,100,200,0},{0,100,0,200,100},{0,100,0,200,100},{0,100,100,200,0},{0,0,100,200,100},{0,100,100,200,0},{0,0,100,200,100},{0,100,100,200,0},{0,0,0,0,100},{0,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{100,0,0,0,0},{100,0,100,0,0},{100,0,0,0,0},{0,0,0,100,100},{100,133,135,0,0},{0,150,144,0,100},{0,200,170,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,0,100},{100,150,0,0,100},{0,135,0,0,100},{100,0,0,0,0},{0,0,100,0,100},{133,0,100,0,100},{150,0,100,0,0},{200,0,0,100,100},{400,0,0,0,0},{400,0,0,0,0},{400,0,0,0,0},{200,0,100,0,0},{400,0,0,0,0},{200,0,100,0,100},{150,0,0,0,100},{133,0,0,0,100},{0,0,100,0,100},{0,0,0,0,100},{0,0,100,100,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,100,0},{200,100,100,0,0},{100,100,0,200,0},{0,100,100,200,0},{0,100,100,200,0},{0,100,100,200,0},{100,100,0,200,0},{100,100,0,0,200},{100,0,0,200,100},{100,100,200,0,0},{100,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,100,100,0},{100,0,0,0,0},{0,0,0,100,100},{0,0,0,0,100},{0,0,144,0,0},{0,144,100,0,0},{0,170,100,0,0},{0,0,400,0,0},{0,400,0,0,0},{100,200,150,0,0},{100,170,133,0,0},{100,144,0,0,0},{100,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,100,0,0},{144,0,0,100,0},{170,0,0,0,100},{200,0,0,100,0},{400,0,0,0,0},{400,0,0,0,0},{400,0,0,0,0},{400,0,0,0,0},{200,0,100,0,0},{150,0,100,0,100},{133,0,100,0,0},{100,0,0,100,0},{100,0,0,100,0},{0,0,100,0,0},{0,0,100,0,100},{100,0,0,0,0},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,100,0},{0,0,0,100,0},{100,0,0,0,100},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{100,0,0,0,100},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,0,0,100},{0,0,100,0,100},{0,0,133,100,0},{100,135,150,0,0},{0,144,100,0,0},{0,0,400,0,0},{0,0,400,0,0},{100,200,200,0,0},{100,170,150,0,0},{0,144,133,100,0},{100,135,0,0,100},{100,0,100,0,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,100,0},{135,0,0,0,100},{144,0,100,0,0},{170,0,100,0,0},{200,0,0,0,100},{400,0,0,0,0},{400,0,0,0,0},{400,0,0,0,0},{200,0,0,0,0},{150,0,100,100,0},{133,0,0,100,0},{0,0,100,0,0},{0,0,100,0,0},{100,0,100,0,0},{0,0,100,0,100},{0,0,100,0,100},{0,0,0,0,100},{100,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{100,0,100,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,100,0,100},{100,0,0,100,0},{100,0,0,0,0},{100,0,100,0,0},{100,0,0,0,0},{100,0,133,0,0},{0,144,150,0,100},{0,170,200,0,100},{0,0,400,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,144,0,0},{100,150,100,0,0},{0,133,0,0,100},{0,0,0,0,100},{0,0,0,100,0},{100,0,0,0,100},{0,0,0,100,0},{100,0,0,100,0},{135,0,100,0,100},{144,0,0,0,100},{170,0,0,0,100},{200,0,100,100,0},{400,0,0,0,0},{200,0,0,100,0},{170,0,0,100,0},{144,0,0,0,100},{100,0,0,0,0},{0,0,100,100,0},{0,0,0,100,0},{100,0,100,0,0},{0,0,0,100,0},{100,0,0,100,0},{100,0,0,0,100},{0,0,100,0,0},{0,0,100,100,0},{0,0,100,0,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,100,0,100},{0,0,0,100,100},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,0,100},{0,0,0,0,100},{0,0,0,0,100},{0,0,100,0,0},{0,0,0,0,100},{0,0,0,0,100},{0,0,0,100,0},{0,133,100,0,0},{100,150,144,0,0},{0,200,170,100,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,135,0,0},{0,150,100,100,0},{0,133,100,0,0},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,100,0,0},{0,0,100,0,100},{100,0,0,100,0},{135,0,100,0,0},{144,0,100,0,100},{170,0,0,0,100},{200,0,100,0,0},{170,0,0,100,100},{144,0,0,100,0},{135,0,100,0,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,0,100},{0,0,0,100,0},{100,0,0,0,0},{0,0,100,100,0},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,100,0,0},{0,0,0,0,100},{0,0,0,0,100},{0,0,100,0,0},{0,0,0,0,100},{0,0,0,0,100},{100,0,0,100,0},{0,0,0,100,100},{0,0,100,0,100},{100,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{100,0,0,0,0},{0,0,0,100,100},{0,0,100,100,0},{0,144,135,100,0},{100,170,144,0,0},{0,200,150,0,0},{0,400,0,0,0},{100,200,135,0,0},{0,170,100,0,100},{0,144,0,100,100},{0,0,100,100,0},{100,0,0,0,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,100,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,100,0,0},{135,0,100,100,0},{144,0,100,0,0},{150,0,0,0,100},{144,0,0,100,0},{135,0,0,0,100},{100,0,0,0,100},{0,0,0,100,100},{100,0,0,100,0},{100,0,0,100,0},{0,0,100,0,100},{100,0,0,0,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,100,0,100},{0,0,100,0,100},{0,0,100,0,0},{0,0,0,100,0},{0,0,100,0,100},{100,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,100,0},{0,0,0,100,0},{0,0,0,0,100},{100,135,0,100,0},{100,144,0,0,100},{100,170,133,0,0},{0,200,100,0,100},{0,170,100,100,0},{0,144,100,0,0},{0,135,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,0,100},{100,0,0,100,0},{0,0,100,100,0},{0,0,100,0,0},{0,0,100,0,100},{0,0,100,0,100},{0,0,0,100,0},{0,0,100,0,100},{133,0,0,100,0},{100,0,0,100,0},{0,133,0,100,100},{0,133,0,100,0},{0,0,0,100,0},{100,0,100,0,0},{100,0,0,100,0},{100,0,0,0,100},{100,0,0,100,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,100,0},{100,0,100,0,0},{0,0,100,0,0},{0,0,0,100,100},{100,0,0,100,0},{100,0,0,0,100},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,0,100,100},{0,0,100,100,0},{0,0,0,100,100},{0,0,0,100,100},{100,0,0,0,0},{100,0,0,0,0},{0,0,100,0,100},{0,0,100,100,0},{100,135,100,0,0},{0,144,100,0,0},{0,150,0,0,100},{100,144,0,0,0},{100,135,0,100,0},{100,0,100,0,0},{0,0,100,0,0},{0,0,0,100,100},{0,0,0,100,0},{0,0,0,100,100},{0,0,0,100,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{100,133,0,0,0},{0,135,0,0,100},{100,144,0,100,0},{0,150,0,100,100},{100,150,0,0,0},{100,144,0,0,0},{0,135,100,100,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,100,0,0},{0,0,0,100,0},{100,0,0,100,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,0,100,0},{100,0,100,0,0},{100,0,0,100,0},{0,0,100,0,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,100,0,0},{0,0,100,133,100},{100,0,0,133,0},{0,0,0,0,100},{100,0,0,0,0},{100,0,0,100,0},{100,0,0,0,0},{0,0,100,0,0},{100,0,100,0,0},{0,0,0,100,0},{100,0,100,0,0},{100,0,0,0,0},{0,133,0,0,100},{0,0,0,100,0},{0,0,100,100,0},{0,0,0,100,100},{0,0,100,0,100},{100,0,0,100,0},{100,0,0,0,100},{0,0,100,0,0},{0,0,0,0,100},{0,0,100,100,0},{100,0,0,100,0},{0,135,0,100,0},{100,144,0,0,0},{0,150,100,100,0},{0,150,0,100,0},{100,170,0,100,0},{100,200,0,100,0},{0,200,0,0,100},{100,170,100,0,0},{0,150,100,0,0},{0,144,0,100,100},{100,135,0,0,100},{0,0,0,0,100},{0,0,0,0,100},{0,0,100,0,0},{100,0,0,0,0},{0,0,100,0,0},{0,0,0,0,100},{100,0,0,0,100},{100,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,135,100},{0,0,100,144,0},{100,0,0,150,0},{0,0,100,150,0},{0,0,0,144,0},{100,0,0,135,100},{0,0,0,0,100},{0,0,0,0,100},{0,0,0,100,100},{100,0,0,100,0},{100,0,0,100,0},{100,0,0,100,0},{0,0,0,0,100},{100,0,0,0,0},{100,0,0,100,0},{0,133,0,100,100},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,0,0,100},{100,0,100,0,0},{100,0,0,0,0},{0,0,0,0,100},{0,135,0,100,0},{0,144,100,100,0},{0,170,0,100,0},{0,200,100,0,100},{0,200,100,100,0},{0,200,0,0,100},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,100,0},{100,200,100,0,0},{100,170,0,100,0},{0,144,0,100,0},{0,0,100,100,0},{0,0,0,0,100},{0,0,0,100,100},{0,0,100,100,0},{100,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{100,0,0,0,0},{0,0,0,100,0},{100,0,0,135,100},{0,0,100,144,0},{0,0,100,170,0},{100,0,0,200,100},{0,0,100,200,0},{0,0,0,170,0},{100,0,100,144,0},{0,0,0,135,0},{100,0,0,0,0},{0,0,100,0,0},{100,0,0,0,100},{100,0,0,100,0},{100,0,0,100,0},{100,0,100,0,0},{0,135,0,0,100},{0,144,0,100,0},{100,150,0,0,100},{0,144,100,100,0},{0,135,0,100,100},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,0,0},{0,0,100,0,0},{0,0,100,0,0},{0,144,100,0,100},{0,170,100,0,100},{100,200,100,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{100,200,100,0,0},{0,150,100,0,0},{100,133,0,0,100},{0,0,100,0,0},{0,133,0,100,100},{100,133,0,100,0},{0,0,100,0,0},{0,0,0,0,100},{0,0,0,100,0},{100,0,0,100,0},{100,0,0,135,0},{0,0,0,144,100},{0,0,0,170,0},{100,0,0,200,0},{0,0,0,400,0},{0,0,0,400,0},{100,0,0,200,0},{0,0,100,170,100},{0,0,0,144,100},{0,0,0,100,0},{100,0,0,100,0},{100,0,0,100,0},{0,0,100,0,0},{100,0,0,100,0},{0,135,0,100,100},{100,144,100,0,0},{0,170,100,100,0},{100,200,0,0,100},{0,170,100,0,100},{100,144,0,100,0},{0,0,0,100,0},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,0,0},{100,133,0,100,0},{0,150,100,100,0},{100,200,0,0,100},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,100,100,0},{100,150,0,100,0},{0,135,0,0,100},{0,144,100,100,0},{0,150,0,100,100},{0,150,0,100,100},{100,144,0,0,100},{0,135,0,100,0},{100,0,0,100,0},{0,0,0,0,100},{100,0,0,144,0},{0,0,0,170,100},{100,0,0,200,100},{0,0,0,400,0},{0,0,0,400,0},{0,0,0,400,0},{0,0,0,400,0},{0,0,100,200,100},{100,0,0,150,0},{0,0,100,133,0},{0,0,0,100,100},{100,0,0,0,0},{100,0,0,0,0},{100,135,0,0,0},{0,144,0,100,0},{0,170,100,0,100},{0,200,0,100,100},{0,400,0,0,0},{100,200,100,0,0},{0,150,100,0,100},{100,133,0,100,0},{0,0,0,100,0},{0,0,100,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,144,100,0,100},{0,170,100,100,0},{100,200,0,0,100},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,100,0,100},{100,150,0,100,0},{0,144,0,100,0},{100,170,0,100,0},{100,200,0,100,0},{0,200,100,0,0},{0,170,100,100,0},{0,150,100,0,0},{0,144,100,0,100},{0,135,100,133,0},{0,0,100,150,100},{100,0,100,200,0},{0,0,0,400,0},{0,0,0,400,0},{0,0,0,400,0},{0,0,0,400,0},{100,0,0,200,0},{0,0,0,170,0},{0,0,0,144,100},{0,0,100,0,0},{100,0,0,100,0},{0,0,100,100,0},{100,0,0,0,100},{100,144,0,0,100},{0,170,100,100,0},{100,200,0,0,100},{0,400,0,0,0},{0,400,0,0,0},{100,200,0,0,100},{0,150,100,100,0},{100,133,0,100,0},{100,0,0,0,0},{100,0,0,0,0},{0,0,100,100,0},{100,0,0,100,0},{0,135,100,100,0},{0,144,0,100,0},{0,170,0,100,100},{0,200,0,0,100},{100,200,0,0,0},{100,200,0,100,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,100,0},{0,170,100,100,0},{0,144,100,100,0},{0,150,0,100,100},{100,200,0,0,100},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,100,0},{100,200,0,0,0},{0,170,100,100,0},{100,144,0,100,0},{100,0,0,144,100},{100,0,0,170,100},{0,0,100,200,0},{0,0,0,400,0},{0,0,0,400,0},{100,0,0,200,0},{0,0,0,170,100},{0,0,100,144,100},{0,0,100,135,0},{0,0,0,100,0},{100,0,0,0,0},{0,0,100,0,0},{100,135,0,0,100},{100,150,0,0,0},{100,200,0,100,0},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{100,200,0,100,0},{100,150,0,0,0},{0,133,100,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{100,0,0,100,0},{100,135,0,0,100},{100,150,100,0,0},{0,200,0,100,100},{0,400,0,0,0},{100,200,0,100,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,0,100},{0,170,0,100,0},{0,144,100,0,0},{100,135,100,0,0},{0,144,100,0,100},{0,170,100,0,100},{0,200,100,0,100},{0,400,0,0,0},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,0,100},{0,150,100,100,0},{0,133,0,100,0},{0,0,0,144,100},{100,0,0,170,0},{0,0,100,200,100},{100,0,0,200,100},{100,0,0,170,0},{100,0,0,144,100},{100,0,0,135,0},{100,0,0,0,100},{100,0,0,0,0},{0,0,100,0,0},{100,0,0,0,0},{100,144,0,100,0},{0,170,0,0,100},{0,200,0,100,100},{0,400,0,0,0},{0,400,0,0,0},{100,200,100,0,0},{0,170,100,100,0},{0,144,100,0,0},{0,0,100,0,0},{0,0,0,100,100},{0,0,0,0,100},{100,0,0,0,0},{100,0,100,0,0},{100,0,0,0,100},{100,0,0,0,100},{0,144,100,100,0},{0,170,0,100,0},{0,200,0,100,100},{0,170,0,0,100},{100,200,100,0,0},{0,200,100,0,100},{100,170,0,0,100},{100,144,0,0,0},{0,135,100,100,0},{100,0,0,0,100},{0,135,100,100,0},{100,144,0,0,100},{100,170,0,0,0},{100,200,0,0,0},{0,200,0,0,100},{0,400,0,0,0},{100,200,0,100,0},{100,150,0,0,0},{100,133,0,0,100},{0,0,100,135,0},{0,0,0,144,100},{0,0,100,150,0},{0,0,0,150,100},{0,0,100,144,0},{0,0,100,135,0},{100,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{100,0,0,0,0},{100,133,0,0,100},{100,150,0,0,100},{0,200,0,0,100},{0,400,0,0,0},{0,400,0,0,0},{0,200,0,100,0},{0,170,100,0,100},{0,144,0,100,0},{0,135,0,100,100},{0,0,0,0,100},{0,0,100,0,0},{0,0,100,0,0},{100,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,0,100},{0,135,0,0,100},{0,144,100,0,100},{0,150,100,100,0},{100,144,100,0,0},{0,150,100,0,100},{100,150,0,100,0},{0,144,100,0,100},{0,135,0,0,100},{0,0,0,0,100},{0,0,100,0,0},{100,0,0,0,0},{100,135,100,0,0},{100,144,0,100,0},{0,150,100,100,0},{100,170,0,100,0},{0,200,100,100,0},{0,170,0,100,100},{0,144,0,0,100},{0,0,0,0,100},{0,0,0,100,100},{0,0,0,0,100},{100,0,0,133,0},{100,0,0,133,100},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,0,0,100},{0,0,100,100,0},{0,0,100,0,0},{100,144,0,0,100},{0,170,100,100,0},{100,200,0,0,100},{0,200,100,0,100},{0,170,0,0,100},{0,144,100,100,0},{0,135,100,100,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,100,0,100},{100,0,0,0,0},{0,0,100,0,0},{0,0,0,100,100},{100,0,100,0,0},{0,0,100,0,0},{100,0,0,0,0},{100,0,0,100,0},{100,133,0,0,100},{0,0,0,100,0},{0,133,0,100,100},{0,133,0,100,100},{0,0,100,0,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,0,100,0},{100,0,0,100,0},{100,0,0,0,0},{100,135,100,0,0},{0,144,0,100,100},{0,150,0,100,100},{0,144,0,0,100},{100,135,0,100,0},{0,0,0,100,100},{0,0,0,100,100},{0,0,100,100,0},{100,0,0,0,0},{100,0,0,0,100},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,0,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,100},{0,135,0,100,100},{100,144,0,0,100},{0,150,100,0,0},{0,150,100,0,100},{100,144,0,0,0},{0,135,100,100,0},{0,0,0,100,100},{0,0,0,0,100},{0,0,0,100,0},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{100,0,0,0,0},{0,0,100,0,100},{0,0,100,0,0},{100,0,0,0,0},{0,0,0,100,0},{0,0,0,0,100},{100,0,0,100,0},{0,0,0,100,100},{0,0,0,100,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,100,0},{0,0,100,0,0},{100,0,0,0,100},{100,0,100,0,0},{0,0,0,100,100},{100,0,0,0,0},{0,133,100,0,100},{0,0,0,100,100},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,0,100},{0,0,0,100,0},{0,0,0,100,0},{100,0,0,100,0},{0,0,0,0,100},{0,0,100,0,0},{0,0,100,0,0},{0,0,0,100,0},{0,0,0,100,0},{0,0,100,0,0},{0,0,100,0,100},{0,0,0,0,100},{100,0,100,0,0},{0,133,100,0,100},{0,133,0,100,100},{100,0,0,0,0},{100,0,0,0,0},{100,0,0,100,0},{0,0,100,0,0},{0,0,0,100,0},{100,0,0,0,100}}
//var elems []elem = []elem{{0,0,100,0,0},{100,100,0,0,0}}
func setRand()  {
	 d := rand.New(rand.NewSource(time.Now().UnixNano()))
/*	for i := 0; i < len(elems); i++ {
		switch {
		case elems[i].Gold == 0:
			elems[i].Gold = r[0]
			fallthrough
		case elems[i].Gold == 100 || elems[i].Gold == 200 || elems[i].Gold == 400:
			elems[i].Gold = elems[i].Gold + r[1]
			fallthrough

		case elems[i].Wood == 0:
			elems[i].Wood = r[2]
			fallthrough
		case elems[i].Wood == 100 || elems[i].Wood == 200 || elems[i].Wood == 400:
			elems[i].Wood = elems[i].Wood + r[3]
			fallthrough

		case elems[i].Lake == 0:
			elems[i].Lake = r[4]
			fallthrough
		case elems[i].Lake == 100 || elems[i].Lake == 200 || elems[i].Lake == 400:
			elems[i].Lake = elems[i].Lake + r[5]
			fallthrough

		case elems[i].Fire == 0:
			elems[i].Fire = r[6]
			fallthrough
		case elems[i].Fire == 100 || elems[i].Fire == 200 || elems[i].Fire == 400:
			elems[i].Fire = elems[i].Fire + r[7]
			fallthrough

		case elems[i].Earth == 0:
			elems[i].Earth = r[8]
			fallthrough
		case elems[i].Earth == 100 || elems[i].Earth == 200 || elems[i].Earth == 400:
			elems[i].Earth = elems[i].Earth + r[9]
		}
	}*/
	// 先随机了，然后在大的里面减去
		for i := 0; i < len(elems); i++{
			r := proRandOne()
			sum := 0
			if elems[i].Gold  == 0 {
				elems[i].Gold = r[0]
				sum = sum + r[0]
			}

			if elems[i].Wood == 0 {
				elems[i].Wood = d.Intn(3)
				sum = sum + elems[i].Wood
			}

			if elems[i].Lake  == 0 {
				elems[i].Lake = r[2]
				sum = sum + r[2]
			}

			if elems[i].Fire == 0 {
				elems[i].Fire = r[3]
				sum = sum + r[3]
			}

			if elems[i].Earth == 0 {
				elems[i].Earth = r[4]
				sum = sum + r[4]
			}


			sub := [5]int{-1,-2,-3,-4,-5}
			num := 0
			if  elems[i].Gold == 100 || elems[i].Gold == 200 || elems[i].Gold == 400 {
				sub[0] = 0
				num++
			}
			if  elems[i].Wood == 100 || elems[i].Wood == 200 || elems[i].Wood == 400 {
				sub[1] = 1
				num++
			}
			if  elems[i].Lake == 100 || elems[i].Lake == 200 || elems[i].Lake == 400 {
				sub[2] = 2
				num++
			}
			if  elems[i].Fire == 100 || elems[i].Fire == 200 || elems[i].Fire == 400 {
				sub[3] = 3
				num++
			}
			if  elems[i].Earth == 100 || elems[i].Earth == 200 || elems[i].Earth == 400 {
				sub[4] = 4
				num++
			}

			eachNum := 0
			if num != 0{
				eachNum = sum / num

			}

			tag := true
			if  elems[i].Gold == 100 || elems[i].Gold == 200 || elems[i].Gold == 400 {
				if tag {
					elems[i].Gold = elems[i].Gold - eachNum - d.Intn(5)
				}
				//tag = false
			}

			if  elems[i].Wood == 100 || elems[i].Wood == 200 || elems[i].Wood == 400  {
				if tag {
					elems[i].Wood = elems[i].Wood - eachNum- d.Intn(2)
				}
				//tag = false
			}

			if  elems[i].Lake == 100 || elems[i].Lake == 200 || elems[i].Lake == 400 {
				if tag {
					elems[i].Lake = elems[i].Lake - eachNum- d.Intn(2)
				}
				//tag = false
			}

			if elems[i].Fire == 100  || elems[i].Fire == 200 || elems[i].Fire == 400 {
				if tag {
					elems[i].Fire = elems[i].Fire - eachNum- d.Intn(6)
				}
				//tag = false
			}

			if elems[i].Earth == 100 || elems[i].Earth == 200 || elems[i].Earth == 400{
				if tag {
					elems[i].Earth = elems[i].Earth - eachNum- d.Intn(5)
				}
				//tag = false
			}

		}

		n := d.Intn(2) + 1
		for i := 0; i < len(elems); i++{
			if elems[i].Wood == 0{
				elems[i].Wood = n
			}
			if elems[i].Wood > 10 {
				elems[i].Wood = elems[i].Wood - n
			}
		}

}

func proRandOne()  []int{
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	//存放结果的slice
	nums := make([]int, 0)
	for i := 0; i < 10; i++{
		// r 是生成0~10的随机数
		r := d.Intn(5)
		// num 是对r进行取整 , 产生1，2。。。。10
		num := int(r) + 1

			nums = append(nums, num)
	}
	return nums
}

func proRandTwo()  []int{
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	//存放结果的slice
	nums := make([]int, 0)
	for i := 0; i < 10; i++{
		// r 是生成0~30的随机数
		r := d.Intn(30)
		// num 是对r进行取整 , 产生1，2。。。。10
		num := int(r) + 1

		nums = append(nums, num)
	}
	return nums
}

func main() {
	setRand()
	js,_ := json.Marshal(&elems)
	jsIndent,_ := json.MarshalIndent(&elems, "", "\t")
	fmt.Println("\njs:\n",string(js), "\n\njsIndent:\n",string(jsIndent))


	allNums := [5]int{100000,100000,100000,100000,100000}

	for i := 0; i < len(elems); i++{
		allNums[0] = allNums[0] - elems[i].Gold
		allNums[1] = allNums[1] - elems[i].Wood
		allNums[2] = allNums[2] - elems[i].Lake
		allNums[3] = allNums[3] - elems[i].Fire
		allNums[4] = allNums[4] - elems[i].Earth
	}

	fmt.Println(allNums)

}

