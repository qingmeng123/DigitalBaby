/*******
* @Author:qingmeng
* @Description:
* @File:BilibiliTest
* @Date2021/10/24
 */

package main

import (
	"fmt"
	"homework2/BasicWork/bilibili"
)

func main() {
	var (
		up bilibili.Uploader
	)
	//声明推荐视频
	ups:=make([]bilibili.Uploader,5)

	//给用户赋值
	user:=bilibili.User{
		BaseInfo: bilibili.BaseInfo{
			Name: "qingmeng",
			VIP:true,
			Grade: 4,
			Icon:"cat.png",
			Signature: "这家伙很懒，没有留下签名",
		},
		Coin: 100,
		RecommendPart:ups,
	}

	//给推荐视频赋值...
	for i:=0;i<5;i++{
		ups[i].Video.VideoName="某视频"
	}

	//给up主赋值
	up.BaseInfo.Name="老师你好我叫何同学"
	up.BaseInfo.Grade=6
	up.BaseInfo.VIP=true
	up.BaseInfo.Signature="xx.png"
	up.BaseInfo.Signature="喜欢做贼有意思的视频..."
	up.Focus=8168000
	up.Video.VideoName="【何同学】我做了苹果放弃的产品..."
	up.Video.VideoState="pause"
	up.Video.VideoLikes=1727000
	up.Video.VideoCount=9040000
	up.Video.VideoCoin=100000
	up.Video.VideoColletion=100000

	//调用方法
	fmt.Println("用户动作：")
	user.PlayorPause(&up)
	user.Like(&up)
	user.GivePoint(2,&up)
	user.Collect(&up)
	user.Sanlian(&up,2)

	//界面信息
	fmt.Println("用户界面信息:\n"+"用户基本信息：(姓名 vip lv 头像 签名)")

	fmt.Println(user.BaseInfo)
	fmt.Println("推荐的视频:")
	fmt.Println(user.RecommendPart)
	fmt.Println("观看的up主视频信息：(视频名称，点赞数，收藏数，硬币数，播放量，播放状态)")
	fmt.Println(up.Video)
	fmt.Println("该视频up主基本信息和关注他的人数：")
	fmt.Println(up.BaseInfo,up.Focus)
}
