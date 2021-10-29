/*******
* @Author:qingmeng
* @Description:
* @File:fight
* @Date2021/10/27
 */

package center

import (
	"fmt"
	"homework2/DigitalBaby/info"
)



func fight(baby *info.DigitalBaby, digitalBaby *info.DigitalBaby)  {
	baby.InitTurn()							//战斗回合
	fmt.Println("战斗开始！\n我方信息:")
	baby.ShowState()
	fmt.Println("对方信息:")
	digitalBaby.ShowState()
	for{
		if baby.IsDead(){
			fmt.Println("数码兽没状态了，退出战斗")
			return
		}else{
			baby.AddTurn()
			fmt.Println("第",baby.GetTurn(),"回合")
			if info.Operate(baby, digitalBaby){
				if digitalBaby.IsDead() {
					fmt.Println(baby.Name,"你真棒！战斗胜利了！")
					return
				}
				baby.Return(digitalBaby)
				if baby.IsDead() {
					fmt.Println("数码兽没状态了，退出战斗")
					return
				}
			}else{
				return
			}
		}
	}
}
