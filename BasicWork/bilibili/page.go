/*******
* @Author:qingmeng
* @Description:
* @File:page
* @Date2021/10/24
 */

package bilibili
import "fmt"

type Video struct {
	VideoName string
	VideoLikes int     							//视频点赞数
	VideoColletion int							//当前视频收藏量
	VideoCoin int								//该视频被投的硬币数
	VideoCount int								//播放量
	VideoState string 							//当前Video的状态：open or pause

}

type Uploader struct {
	BaseInfo BaseInfo
	Focus int 									//关注人数
	Video Video									//当前页面放着的视频
}

type BaseInfo struct {
	Name string
	VIP bool
	Grade int
	Icon string     							//头像
	Signature string							//签名

}

type User struct {
	BaseInfo BaseInfo							//用户基本信息
	RecommendPart []Uploader					//推荐视频部分
	Coin int									//用户的硬币数
}

type UserAction interface {			//一个类具有了一个接口的所记录的所有功能，该类型才算“实现”了这个接口
	Like(up *Uploader)							//点赞
	Collect(up *Uploader)						//收藏
	GivePoint(n int ,up *Uploader)				//投币
	PlayorPause(up *Uploader)					//播放
	Sanlian(up *Uploader,n int)
}

func(u *User)Like(up *Uploader){				//给视频点赞
	fmt.Println(u.BaseInfo.Name,"给视频:",up.Video.VideoName,"点了赞")
	up.Video.VideoLikes++
	fmt.Println("该视频的点赞量为：",up.Video.VideoLikes)
}

func(u *User)Collect(up *Uploader){				//收藏视频
	fmt.Println(u.BaseInfo.Name,"把up主",up.BaseInfo.Name,"的视频:",up.Video.VideoName,"收藏了！")
	fmt.Println("该视频的收藏量为：",up.Video.VideoColletion)
}

func(u *User)GivePoint(n int,up *Uploader){		//给up主投币
	if n>0&&n<=2{
		fmt.Println(u.BaseInfo.Name,"给视频:",up.Video.VideoName,"投了",n,"枚币")
		up.Video.VideoCoin+=n
		u.Coin-=n
	}else{
		fmt.Println("投币失败，投币数只能为1~2")
	}
	fmt.Println("该视频目前被投了",up.Video.VideoCoin,"枚硬币")
}


func(u *User)PlayorPause(up *Uploader){ 		//更改视频播放状态

	if up.Video.VideoState=="playing"{
		up.Video.VideoState="pause"
	}else{
		up.Video.VideoState="playing"
		up.Video.VideoCount++
	}
	fmt.Println(u.BaseInfo.Name,"更改了视频播放状态,现在的播放状态是：",up.Video.VideoState)

}

func (u *User)Sanlian(up *Uploader,n int){
	fmt.Println("您给该视频一件三连了！")
	u.GivePoint(n,up)
	u.Like(up)
	u.Collect(up)
}

