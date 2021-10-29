/*******
* @Author:qingmeng
* @Description:
* @File:digitalBabyGroup
* @Date2021/10/29
 */

package info

import "fmt"

type MyDigitalBabyGroup interface {
	AddNewBaby(baby DigitalBaby)		//增加数码兽
	ShowBabies()						//展示数码兽
	DeleteDigitalBaby()					//放生数码兽
	ReadyBaby()DigitalBaby
}
//我的数码宝贝们
type MyDigitalBabyGroupInfo struct {
	DigitalBabies []DigitalBaby

}

func (m *MyDigitalBabyGroupInfo) ReadyBaby() DigitalBaby {
	i:=0
	fmt.Println("您的数码宝贝有：")
	m.ShowBabies()
	if len(m.DigitalBabies)>1{
		fmt.Println("请输入您想用来战斗的数码宝贝(id)：")
		fmt.Scan(&i)
		if i< len(m.DigitalBabies) {
			fmt.Printf("选择成功，您选择了%s作为战斗数码兽\n",m.DigitalBabies[i].Name)
			return m.DigitalBabies[i]
		}else {
			m.ReadyBaby()
	}
	}
	fmt.Printf("选择失败，您只有%s作为战斗数码兽\n",m.DigitalBabies[0].Name)
	return m.DigitalBabies[0]
}

func (m *MyDigitalBabyGroupInfo) AddNewBaby(baby DigitalBaby) {
	m.DigitalBabies=append(m.DigitalBabies,baby)
}

func (m *MyDigitalBabyGroupInfo) ShowBabies() {
	for i, baby := range m.DigitalBabies {
		fmt.Printf("id: %d ",i)
		baby.ShowState()
		baby.ShowsKills()
	}
}

func (m *MyDigitalBabyGroupInfo) DeleteDigitalBaby() {
	fmt.Println("您的数码宝贝有：")
	m.ShowBabies()
	fmt.Println("请选择您想放生的数码宝贝：")
	i:=0
	fmt.Scan(i)
	if i< len(m.DigitalBabies){
		a:=m.DigitalBabies[:i]
		b:=m.DigitalBabies[i+1:]
		m.DigitalBabies=append(a,b...)
	}else {
		m.DeleteDigitalBaby()
	}
	fmt.Println("放生成功，您现在的数码宝贝有：")
	m.ShowBabies()
}

