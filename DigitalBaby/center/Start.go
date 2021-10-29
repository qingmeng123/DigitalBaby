/*******
* @Author:qingmeng
* @Description:
* @File:Start
* @Date2021/10/26
 */

package center

import (
	"fmt"
	"homework2/DigitalBaby/info"
)


func Start()  {
	//初始化库
	BabyLib:= info.LibDigitalBabys{}
	SkillsLib:= info.LibSkills{}
	info.InitBabies(&BabyLib)
	SkillsLib.InitSkills()
	Map:=MapInfo{}

	//初始化我的数码兽库
	myDigitalBabies:=info.MyDigitalBabyGroupInfo{}

	n:=0
	for  {
		fmt.Println("欢迎进入数码世界！")
		fmt.Println("亚古兽：太一~~！\n太一：亚古兽！！")
		fmt.Println("亚古兽：太一，好久不见，我来找你了")
		fmt.Println("太一：我可想你了！转眼间我已经上大学了")
		fmt.Println("亚古兽：我也好想来找你，这次能来，就是因为数码世界出现了时空兽，打破了各个时空的平衡。\n又需要咱们去拯救世界了")
		fmt.Println("太一：这样呀！怪不得地球最近也怪事频频，疫情连连!")
		fmt.Println("亚古兽：那让我们重启旅行吧！时空异动，我们可以遇到各个时期的数码兽哦")
		fmt.Println("太一：那让我们出发吧！")
		fmt.Println("**************************")
		fmt.Println("您获得了初始数码兽:亚古兽（亚古兽是信仰，不给你选择初始兽的权力~)")

		//给你初始兽
		yagushou:=info.DigitalBabies[0]
		myDigitalBabies.AddNewBaby(yagushou)
		myDigitalBabies.ShowBabies()
		fmt.Println("hhhh~偷偷送你一直奥米加(绝版)")
		myDigitalBabies.AddNewBaby(info.DigitalBabies[20])

		myDigitalBaby:=myDigitalBabies.ReadyBaby()
		fmt.Println("1.开始游戏\n2.保存游戏\n0.退出游戏")
		fmt.Scan(&n)
		switch n {
		case 0:
			fmt.Println("游戏结束")
			return
		case 1:
			for{
				fmt.Println("1.选择地图\n2.随机事件\n3.调整阵容\n0.返回主菜单")
				fmt.Scan(&n)
				switch n {
				case 1:
						Map.ChooseMap(&myDigitalBaby)
				case 2:
					fmt.Println("野生的",info.DigitalBabies[info.GetRandom(len(info.DigitalBabies)-1)].Name,"跳出来了！")
					fight(&myDigitalBaby,&info.DigitalBabies[info.GetRandom(len(info.DigitalBabies)-1)])
				case 3:
					for{
						myDigitalBabies.ShowBabies()
						fmt.Println("1.替换战斗宝贝\n2.放生一个宝贝\n3.增加宠物\n4.给数码兽升级\n其他返回上一级")
						fmt.Scan(&n)
						switch n {
						case 1:
							myDigitalBaby=myDigitalBabies.ReadyBaby()
						case 2:
							myDigitalBabies.DeleteDigitalBaby()
						case 3:
							fmt.Println("哈哈哈，外挂到了,免费抽宠！")
							fmt.Println("请输入想要获得宠物的只数,不超过10只")
							i:=0
							fmt.Scan(&i)
							for i := 0; i < 10; i++ {
								myDigitalBabies.AddNewBaby(info.DigitalBabies[info.GetRandom(i)])
							}
						case 4:
							fmt.Println("请选择您想升级的数码兽(id)")
							i:=0
							fmt.Scan(&i)
							myDigitalBabies.DigitalBabies[i].Upgrade()
						default:
							break
						}
					}
				case 0:
					break
				}
			}
		case 2:	//保存游戏。。。不会呀。。
			break
		}
	}
}