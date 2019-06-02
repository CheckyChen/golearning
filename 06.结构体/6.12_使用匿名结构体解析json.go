package main

import (
	"encoding/json"
	"fmt"
)

// 手机拥有屏幕、电池、指纹识别等信息，将这些信息填充为 JSON 格式的数据。如果需要选择性地分离 JSON 中的数据则较为麻烦。Go 语言中的匿名结构体可以方便地完成这个操作

type Screen struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}

type Battery struct {
	Capacity int // 容量
}

func GenJsonData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool //// 序列化时添加的字段：是否有指纹识别
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			Capacity: 2910,
		},
		HasTouchID: true,
	}
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	jsonData := GenJsonData()
	fmt.Println(string(jsonData))

	//	 只需要屏幕和指纹识别信息的结构和是来
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Printf("%+v\n", screenAndTouch)

	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData, &batteryAndTouch)
	fmt.Printf("%+v\n", batteryAndTouch)

}
