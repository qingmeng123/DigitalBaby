/*******
* @Author:qingmeng
* @Description:
* @File:digitalBaby
* @Date2021/10/25
 */

package info

import (
	"fmt"
)

var lib = LibDigitalBabys{}

type Baby interface {
	ShowState()				//显示状态信息
	ShowsKills()			//显示技能基本信息
	IsDead()bool
	Promote

}

type Promote interface {
	GetLv()		int                	   	//获取等级
	Upgrade()                         //升级
	Learn(skill Skill)                //技能学习
	LearnFromAll(lib libSkillsAction) //从技能库选择技能学习
	Forget()                          //技能遗忘
	IsEvolute()bool                   //判断是否可进化
	LearnRandom(lib libSkillsAction)  //从技能库随机学习技能
	State
}

type DigitalBaby struct  {
	Id				int
	Name 			string
	Lv 				int
	Hp				int		//血量
	Exp   			int
	Skills			[]Skill			//这个地方导致后面改了又改，考虑到数码兽技能不满4个，避免空指针，还是用动态的吧
	StateInfo
	PropertyInfo
	fightTurn
}

func (d *DigitalBaby) Learn(skill Skill) {
	i:= len(d.Skills)-1
	if i<4{
		d.Skills=append(d.Skills,skill)
		fmt.Printf("%s学习了新技能:\n",d.Name)
		skill.ShowSkill()
		return
		}
	d.Forget()
	d.Learn(skill)
}

func (d *DigitalBaby) LearnRandom(lib libSkillsAction) {
	d.Learn(lib.GetRandomSkill())
}

func (d *DigitalBaby) IsEvolute() bool {
	if d.GetLv()==15||d.GetLv()==30||d.GetLv()==50{
		return true
	}
	return false
}

func (d *DigitalBaby)Evolute(promote Promote)  {
	fmt.Println("*************************")
	fmt.Printf("%s~~进化！！！\n",d.Name)
	c:=lib.GetOneBaby(d.Id+5)
	d.Name=c.Name
	d.Id=c.Id
	d.dp=c.dp
	d.Hp=c.Hp
	fmt.Println(d.Name+"~~!!!!!!!")
	d.Learn(c.Skills[0])
}

func (d *DigitalBaby) GetLv() int {
	return d.Lv
}

func (d *DigitalBaby) Forget() {
	i:=3
	fmt.Println("请输入需要遗忘的技能（0~3）(非技能番号返回)")
	d.ShowsKills()
	fmt.Scan(&i)
	if i>=0&&i<4{
		fmt.Printf("遗忘%s成功\n",d.Skills[i].SkillName())
		d.Skills[i]=nil
	}else {
		fmt.Println("取消遗忘")
	}
}

func (d *DigitalBaby) LearnFromAll(lib libSkillsAction) {
	fmt.Println("有如下技能:")
	n:=lib.ShowSomeSkills(AllSkills)
	i:=0
	fmt.Println("请选择你想学习的技能：(非技能番号退出学习技能)")
	fmt.Scan(&i)
	if i<=n&&i>=0{
		d.Learn(lib.GetSkill(i))
	}else{
		fmt.Println("退出技能学习")
	}
}

func (d *DigitalBaby) Upgrade() {
	fmt.Printf("%s升级了！\n",d.Name)
	d.Lv+=1
	if d.IsEvolute(){					//等级达到要求触发进化
		fmt.Printf("%s达到了进化的等级！\n",d.Name)
		d.Evolute(d)
		d.ShowState()

	}else{
		d.Hp+=10
		d.Exp=0
		d.dp+=1
		d.ShowState()
	}
	d.LearnRandom(LibSkills{})
}

func (d *DigitalBaby) ShowsKills() {
	for i, skill := range d.Skills {
		fmt.Println(i,skill.SkillName(),"剩余次数",skill.SkillMp(),"伤害",skill.SkillHarm())
	}
}

func (d *DigitalBaby) ShowState() {
	fmt.Println(d.Name,d.form.name,"Lv:",d.Lv,"Hp:",d.Hp,"Dp:",d.dp,"状态：",d.special.name)
}

func (d *DigitalBaby) IsDead() bool {
	return d.Hp<=0
}

//攻击
func (d *DigitalBaby) Attack(target *DigitalBaby) {
	if target.FeedBack(d){
		return
	}
	fmt.Println(d.Name,"的技能情况：")
	d.ShowsKills()
	fmt.Println("请输入需要释放的技能(id):")
	var i int
	fmt.Scan(&i)
	if i< len(d.Skills)&&d.Skills[i].SkillMp()>0{
		if d.ChooseAble(target){
			fmt.Println(d.Name,"释放了技能：\n",d.Skills[i].SkillName(),"!!!!!")
			d.Skills[i].Choose(target)
		}
	}else{
		d.Attack(target)
	}
}

func (d *DigitalBaby) Cure(target *DigitalBaby)  {
	d.Hp+=100
	fmt.Printf("给%s加了100滴血\n现在的Hp为:%d\n",d.Name,d.Hp)
}

func (d *DigitalBaby) FeedBack(target *DigitalBaby) bool {
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


func (d *DigitalBaby) Catch(target *DigitalBaby)  {		//捕捉

	fmt.Println("非捕捉数码兽，禁止捕捉！")
}

func (d *DigitalBaby) Escape()  { //逃跑
	if d.ChooseAble(d){
		fmt.Println("主动逃跑成功了！")
	}
}

//是否选择
func (d *DigitalBaby) ChooseAble(target *DigitalBaby)bool  {
	return true
}

func (d *DigitalBaby) Choose(target *DigitalBaby)  {
}



