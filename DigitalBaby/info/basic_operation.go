/*******
* @Author:qingmeng
* @Description:
* @File:basic_operation
* @Date2021/10/25
 */

package info

import (
	"fmt"
)

//基础操作
type Operation interface {
	Choose(target *DigitalBaby)   //选择目标
	FeedBack(target *DigitalBaby)bool //目标反馈	有控制返回true,无控制返回false
	Action
}

//操作
func Operate(operator Operation,target *DigitalBaby)bool  {
	n:=0
	fmt.Println("请选择您想进行的操作(id):")
	fmt.Println("1.攻击","2.治疗","3.捕捉","4.逃跑")
	fmt.Scan(&n)
	switch n {
		case 1:
			operator.Attack(target)
		case 2:
			operator.Cure(target)
		case 3:
			operator.Catch(target)
		case 4:
			operator.Escape()
			return false					//逃跑，不进行操作了
		default:
			Operate(operator,target)
	}
	return true								//继续操作
}

//基本行动
type Action interface {
	Attack(target *DigitalBaby) //攻击
	Cure(target *DigitalBaby)   //治疗
	Catch(target *DigitalBaby)  //捕捉
	Escape()                    //逃跑
}

//结束攻击，交换回合
func (d *DigitalBaby) Return(other *DigitalBaby)  {
	fmt.Println("现在轮到对方的回合")
	if !d.FeedBack(other){
		fmt.Printf("%s释放了技能:\n",other.Name)
		other.Skills[GetRandom(len(other.Skills)-1)].Attack(d)				//敌方只会攻击,且技能随机
	}
}

//回合
type Turn interface {
	InitTurn()
	AddTurn()
	GetTurn()int
}

type fightTurn struct {
	turn int	//回合
}

func (f* fightTurn) InitTurn() {
	f.turn=0
}

func (f* fightTurn) AddTurn() {
	f.turn++
}

func (f* fightTurn) GetTurn()int {
	return f.turn
}

//目标
type Target interface {
	ChooseAble(target *DigitalBaby)bool
}





