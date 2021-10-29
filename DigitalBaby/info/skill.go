/*******
* @Author:qingmeng
* @Description:
* @File:skill
* @Date2021/10/25
 */

package info

import (
	"fmt"
)

type Skill interface {
	SkillId()   int
	SkillName() string
	SkillMp()	int
	SkillHarm()	int
	ShowSkill()			//显示单个技能信息详细信息

	Operation
}

//技能影响
type SkillEffect interface {
	Hurt(target *DigitalBaby)   //技能伤害
	Effect(target *DigitalBaby) //技能效果
	State
}


//技能
type BaseSkill struct {
	id   		int
	name        string
	mp			int
	harm		int
	StateInfo

}

func (b *BaseSkill) ShowSkill() {
	fmt.Printf("id：%d,技能名：%s,剩余次数：%d,伤害:%d\n",b.id,b.name,b.mp,b.harm)
}

func (b *BaseSkill) Attack(target *DigitalBaby) {
	fmt.Println(b.SkillName())
	b.Choose(target)
}

func (b *BaseSkill) Cure(target *DigitalBaby) {
	target.Hp+=100
	fmt.Printf("使用技能%s给%s加了100滴血",b.SkillName(),target.Name)
}

func (b *BaseSkill) Catch(target *DigitalBaby) {
	panic("implement me")
}

func (b *BaseSkill) Escape() {
}

func (b *BaseSkill) SkillId() int {
	return b.id
}

func (b *BaseSkill) SkillName() string {
	return b.name
}

func (b *BaseSkill) SkillMp() int {
	return b.mp
}

func (b *BaseSkill) SkillHarm() int {
	return b.harm
}

func (b *BaseSkill) Hurt(target *DigitalBaby) {
	hurt:=b.SkillHarm()-target.PropertyInfo.BabyDp()
	if hurt<0{
		hurt=1
	}
	fmt.Println(target.Name,"受到了",hurt,"点伤害")
	target.Hp-=hurt
	fmt.Println(target.Name,"现在的Hp:",target.Hp)
}

func (b *BaseSkill) Effect(target *DigitalBaby) {
	target.ChangeSpecial(b.GetSpecial())
}

func (b *BaseSkill) Choose(target *DigitalBaby) {
		b.mp-=1
		b.Hurt(target)
		b.Effect(target)
}

func (b *BaseSkill) FeedBack(target *DigitalBaby) bool{
	if target.special.id==1{
		fmt.Printf("%s睡着了，不能行动\n",target.Name)
		target.special.id--
		return	true
	}
	if target.special.id==2{
		fmt.Printf("%s灼烧了，每回合扣%d滴血\n",target.Name,target.Hp/10)
		target.Hp-=20
	}
	return false
}

//是否选择
func (b *BaseSkill) ChooseAble(target *DigitalBaby) bool {
	fmt.Println("确定选择？\n1.确定\n任意键返回")
	i:=0
	fmt.Scan(&i)
	if i==1{
		return true
	}
	return false
}

func NewSkill(id int, name string, mp int, harm int, Special Special) Skill {
	return &BaseSkill{
		id:      id,
		name:    name,
		mp:      mp,
		harm:    harm,
		StateInfo:StateInfo{
			form:    Form{},
			special: Special,
		},
	}
}


