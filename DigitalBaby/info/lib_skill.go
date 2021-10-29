/*******
* @Author:qingmeng
* @Description:技能库
* @File:lib_skill
* @Date2021/10/26
 */

package info

import (
	"math/rand"
	"time"
)

var AllSkills  []Skill

type LibSkills struct {
}

func (l LibSkills) ShowSomeSkills(skill []Skill) int {
	n:=0
	for i, s := range skill {
		s.ShowSkill()
		n=i
	}
	return n
}

func (l LibSkills) GetSkill(i int) Skill {
	return AllSkills[i]
}

func (l LibSkills) GetRandomSkill() Skill {
	return AllSkills[GetRandom(len(AllSkills)-1)]
}

func (l LibSkills) add(skill Skill) {
	AllSkills =append(AllSkills,skill)
}

func (l LibSkills) InitSkills() {
	s0:= NewSkill(0, "小型火焰", 20, 20, Specials[2])
	l.add(s0)
	s1:= NewSkill(1, "小火焰弹", 20, 20, Specials[2])
	l.add(s1)
	l.add(NewSkill(2, "空气炮", 15, 40, Specials[0]))
	l.add(NewSkill(3, "魔法火焰", 15, 40, Specials[0]))
	l.add(NewSkill(4, "V仔兽头椎", 15, 40, Specials[0]))
	l.add(NewSkill(5, "狐叶楔", 15, 40, Specials[0]))
	l.add(NewSkill(6, "破岩爪", 15, 40, Specials[0]))
	l.add(NewSkill(7, "摇滚魂", 15, 40, Specials[0]))
	l.add(NewSkill(8, "抓击", 15, 40, Specials[0]))
	l.add(NewSkill(9, "撞   击", 10, 30, Specials[0]))
	l.add(NewSkill(10, "治疗术", 15, 0, Specials[0]))
	l.add(NewSkill(11, "巨型火焰", 15, 60, Specials[2]))
	l.add(NewSkill(12, "妖狐火焰", 15, 60, Specials[2]))
	l.add(NewSkill(13, "天堂之拳", 15, 60, Specials[0]))
	l.add(NewSkill(14, "狂抓", 15, 60, Specials[0]))
	l.add(NewSkill(15, "V龙拳", 15, 60, Specials[0]))
	l.add(NewSkill(16, "摇滚魂", 15, 40, Specials[0]))
	l.add(NewSkill(17, "催眠曲", 15, 40, Specials[1]))
}

type libSkillsAction interface {
	add(skill Skill)
	InitSkills()
	GetSkill(i int)Skill
	GetRandomSkill()Skill			//随机获取个技能
	ShowSomeSkills(skill []Skill)int	//
}

func GetRandom(i int)int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(i)
}



