/*******
* @Author:qingmeng
* @Description:
* @File:map
* @Date2021/10/27
 */

package center

import (
	"fmt"
	"homework2/DigitalBaby/info"
)

var M=make(map[string][]info.DigitalBaby)

type Map interface {
	ChooseMap(baby *info.DigitalBaby)
	InitMap()
}

type MapInfo struct {
	info.LibDigitalBabys
}

func (m MapInfo) ChooseMap(baby *info.DigitalBaby) {
	info.InitBabies(&m.LibDigitalBabys)
	m.InitMap()
	for{
		fmt.Println("请选择您想进入的图:")
		fmt.Println("1.新手村\n2.刷级区\n3.精英怪\n4.Danger\n0.返回上一级")
		i:=0
		fmt.Scan(&i)
		switch i {
		case 0:
			return
		case 1:
			fmt.Println("遇见了可爱的成长期数码兽：")
			fight(baby,&M["成长期兽"][info.GetRandom(4)])
		case 2:
			fmt.Println("刷到一直成熟期数码兽：")
			fight(baby,&M["成熟期兽"][info.GetRandom(4)])
		case 3:
			fmt.Println("oh!遭遇完全体数码兽袭击：")
			fight(baby,&M["完全体兽"][info.GetRandom(4)])
		case 4:
			fmt.Println("Danger!Danger!遭遇BOSS：")
			fight(baby,&M["究极体兽"][info.GetRandom(4)])
		default:
			m.ChooseMap(baby)
		}
	}
}

func (m MapInfo) InitMap() {
	M["成长期兽"]=m.GrowthDigitalBabies
	M["成熟期兽"]=m.MatureDigitalBabies
	M["完全体兽"]=m.CompleteDigitalBabies
	M["究极体兽"]=m.UltimateDigitalBabies
}

