/*******
* @Author:qingmeng
* @Description:数码宝贝库
* @File:lib_digitalBaby
* @Date2021/10/26
 */

package info

var DigitalBabies []DigitalBaby

type LibDigitalBabies interface {
	GetOneBaby(i int)*DigitalBaby
	add(new DigitalBaby)
}

type LibDigitalBabys struct {
	GrowthDigitalBabies 	[]DigitalBaby		//成长期
	 MatureDigitalBabies []DigitalBaby			//成熟期
	 CompleteDigitalBabies []DigitalBaby		//完全体
	 UltimateDigitalBabies []DigitalBaby		//究极体
}

func (b *LibDigitalBabys) GetOneBaby(i int) *DigitalBaby {
	return &DigitalBabies[i]
}

//增加数码兽
func (b *LibDigitalBabys)add(new DigitalBaby) {
	DigitalBabies =append(DigitalBabies,new)
}

//初始化数码库
func InitBabies(b *LibDigitalBabys)  {
	b0:= DigitalBaby{
		Id:   0,
		Name: "亚古兽",
		Lv:   1,
		Hp:   100,
		Exp:  0,
		Skills: []Skill{
			NewSkill(0, "小型火焰", 10, 40,Specials[2]),
			NewSkill(1, "撞   击", 10, 30, Specials[0]),
			NewSkill(1, "撞   晕", 10, 30, Specials[1]),
		},
		StateInfo:    StateInfo{Forms[0],Specials[0]},
		PropertyInfo: PropertyInfo{10},
		fightTurn:    fightTurn{},
	}
	b.add(b0)

	b1:= DigitalBaby{
		Id:           1,
		Name:         "加布兽",
		Lv:           1,
		Hp:           100,
		Skills:       []Skill{
			NewSkill(0, "小火焰球", 10, 40, Specials[2]),
			NewSkill(1, "撞   击", 10, 30, Specials[0]),
		},
		PropertyInfo: PropertyInfo{10},
		StateInfo:    StateInfo{Forms[0],Specials[0]},
	}
	b.add(b1)

	b2:=DigitalBaby{
		Id:           2,
		Name:         "巴达兽",
		Lv: 			1,
		Hp:           100,
		Skills:       []Skill{
			NewSkill(0, "空气炮", 15, 40, Specials[0]),NewSkill(1, "撞   击", 10, 30, Specials[0]),
		},
		PropertyInfo: PropertyInfo{10},
		StateInfo:    StateInfo{Forms[0],Specials[0]},
	}

	b.add(b2)
	b3:=DigitalBaby{
		Id:           3,
		Name:         "小狗兽",
		Lv: 			1,
		Hp:           100,
		Skills:       []Skill{
			NewSkill(0, "抓击", 15, 40, Specials[0]),NewSkill(1, "撞   击", 10, 30, Specials[0]),
		},
		PropertyInfo: PropertyInfo{10},
		StateInfo:    StateInfo{Forms[0],Specials[0]},
	}
	b.add(b3)

	b4:=DigitalBaby{
		Id:           4,
		Name:         "V仔兽",
		Lv: 			1,
		Hp:           100,
		Skills:       []Skill{
			NewSkill(0, "V仔兽头椎", 15, 40, Specials[0]),NewSkill(1, "撞   击", 10, 30, Specials[0]),

		},
		PropertyInfo: PropertyInfo{
			dp: 10,
		},
		StateInfo:    StateInfo{Forms[0],Specials[0]},
	}
	b.add(b4)

	b5:=DigitalBaby{
		Id:   5,
		Name: "暴龙兽",
		Lv:   15,
		Hp:   300,
		Exp:  0,
		Skills: []Skill{
			NewSkill(0, "巨型火焰", 15, 60, Specials[2]),NewSkill(0, "巨型火焰", 15, 60, Specials[2]),
		},
		StateInfo:    StateInfo{Forms[1],Specials[0]},
		PropertyInfo: PropertyInfo{30},
	}
	b.add(b5)

	b6:=DigitalBaby{
		Id:           6,
		Name:         "加鲁鲁兽",
		Lv: 			15,
		Hp:           300,
		Skills:       []Skill{
			NewSkill(0, "妖狐火焰", 15, 60, Specials[2]),NewSkill(0, "妖狐火焰", 15, 60, Specials[2]),
		},
		PropertyInfo: PropertyInfo{30},
		StateInfo:    StateInfo{Forms[1],Specials[0]},
	}
	b.add(b6)

	b7:=DigitalBaby{
		Id:           7,
		Name:         "天使兽",
		Lv: 			15,
		Hp:           300,
		Skills:       []Skill{
			NewSkill(0, "天堂之拳", 15, 60, Specials[0]),
			NewSkill(0, "天堂之拳", 15, 60, Specials[0]),
		},
		PropertyInfo: PropertyInfo{30},
		StateInfo:    StateInfo{Forms[1],Specials[0]},
	}
	b.add(b7)

	b8:=DigitalBaby{
		Id:           8,
		Name:         "迪路兽",
		Lv: 			15,
		Hp:           300,
		Skills:       []Skill{
			NewSkill(0, "狂抓", 15, 60, Specials[0]),
			NewSkill(0, "狂抓", 15, 60, Specials[0]),
		},
		PropertyInfo: PropertyInfo{30},
		StateInfo:    StateInfo{Forms[1],Specials[0]},
	}
	b.add(b8)
	b9:=DigitalBaby{
		Id:           9,
		Name:         "V龙兽",
		Lv: 			15,
		Hp:           300,
		Skills:       []Skill{
			NewSkill(0, "V龙拳", 15, 60, Specials[0]),
			NewSkill(0, "V龙拳", 15, 60, Specials[0]),
		},
		PropertyInfo: PropertyInfo{30},
		StateInfo:    StateInfo{Forms[1],Specials[0]},
	}
	b.add(b9)
	b10:=DigitalBaby{
		Id:           10,
		Name:         "机械暴龙兽",
		Lv: 			30,
		Hp:           500,
		Skills:       []Skill{
			NewSkill(0, "鱼叉机关炮", 15, 100, Specials[0]),
			NewSkill(0, "鱼叉机关炮", 15, 100, Specials[0]),
		},
		PropertyInfo: PropertyInfo{50},
		StateInfo:    StateInfo{Forms[2],Specials[0]},
	}
	b.add(b10)
	b11:=DigitalBaby{
		Id:           11,
		Name:         "兽人加鲁鲁兽",
		Lv: 			30,
		Hp:           500,
		Skills:       []Skill{
			NewSkill(0, "凯撒锐爪", 15, 100, Specials[0]),
			NewSkill(0, "凯撒锐爪", 15, 100, Specials[0]),
		},
		PropertyInfo: PropertyInfo{50},
		StateInfo:    StateInfo{Forms[2],Specials[0]},
	}
	b.add(b11)
	b12:=DigitalBaby{
		Id:           12,
		Name:         "神圣天使兽",
		Lv: 			30,
		Hp:         	 500,
		Skills:       []Skill{
			NewSkill(0, "神圣光波", 10, 100, Specials[1]),
			NewSkill(0, "神圣光波", 10, 100, Specials[1]),
		},
		PropertyInfo: PropertyInfo{50},
		StateInfo:    StateInfo{Forms[2],Specials[0]},
	}
	b.add(b12)
	b13:=DigitalBaby{
		Id:           13,
		Name:         "天女兽",
		Lv: 			30,
		Hp:           500,
		Skills:       []Skill{
			NewSkill(0, "神圣之箭", 10, 100, Specials[0]),
			NewSkill(0, "神圣之箭", 10, 100, Specials[0]),

		},
		PropertyInfo: PropertyInfo{50},
		StateInfo:    StateInfo{Forms[2],Specials[0]},
	}
	b.add(b13)
	b14:=DigitalBaby{
		Id:           14,
		Name:         "天翔V龙兽",
		Lv: 			30,
		Hp:           100,
		Skills:       []Skill{
			NewSkill(0, "V翼斩", 10, 100, Specials[0]),
			NewSkill(0, "V翼斩", 10, 100, Specials[0]),

		},
		PropertyInfo: PropertyInfo{50},
		StateInfo:    StateInfo{Forms[2],Specials[0]},
	}
	b.add(b14)
	b15:=DigitalBaby{
		Id:           15,
		Name:         "战斗暴龙兽",
		Lv: 			50,
		Hp:           1000,
		Skills:       []Skill{
			NewSkill(0, "盖亚能量炮", 5, 200, Specials[0]),
			NewSkill(0, "盖亚能量炮", 5, 200, Specials[0]),
		},
		PropertyInfo: PropertyInfo{100},
		StateInfo:    StateInfo{Forms[3],Specials[0]},
	}
	b.add(b15)
	b16:=DigitalBaby{
		Id:           16,
		Name:         "钢铁加鲁鲁兽",
		Lv: 			50,
		Hp:           1000,
		Skills:       []Skill{
			NewSkill(0, "冰冻炸弹", 5, 200, Specials[0]),NewSkill(1, "优雅十字光线", 5, 200, Specials[0]),

		},
		PropertyInfo: PropertyInfo{100},
		StateInfo:    StateInfo{Forms[3],Specials[0]},
	}
	b.add(b16)
	b17:=DigitalBaby{
		Id:           17,
		Name:         "炽天使兽",
		Lv: 			50,
		Hp:           1000,
		Skills:       []Skill{
			NewSkill(0, "七重天堂", 5, 200, Specials[0]),
			NewSkill(0, "七重天堂", 5, 200, Specials[0]),
		},
		PropertyInfo: PropertyInfo{100},
		StateInfo:    StateInfo{Forms[3],Specials[0]},
	}
	b.add(b17)

	b18:=DigitalBaby{
		Id:           18,
		Name:         "座天使兽",
		Lv: 			50,
		Hp:           1000,
		Skills:       []Skill{
			NewSkill(0, "伊甸标枪", 10, 180, Specials[0]),NewSkill(0, "治疗光波", 10, 0, Specials[0]),

		},
		PropertyInfo: PropertyInfo{100},
		StateInfo:    StateInfo{Forms[3],Specials[0]},
	}
	b.add(b18)
	b19:=DigitalBaby{
		Id:           19,
		Name:         "究极V龙兽",
		Lv: 			50,
		Hp:           1000,
		Skills:       []Skill{
			NewSkill(0, "闪光V热线", 5, 200, Specials[2]),NewSkill(1, "究极力之剑", 5, 200, Specials[0]),

		},
		PropertyInfo: PropertyInfo{100},
		StateInfo:    StateInfo{Forms[3],Specials[0]},
	}
	b.add(b19)
	b20:=DigitalBaby{
		Id:           20,
		Name:         "奥米加兽",
		Lv:           100,
		Hp:           2000,
		Exp:          0,
		Skills:		[]Skill{
			NewSkill(0, "暴龙剑", 10, 300, Specials[2]),
			NewSkill(1, "加鲁鲁炮", 5, 400, Specials[1]),
			NewSkill(1, "冰冻弹", 10, 30, Specials[1]),
		},
		StateInfo:    StateInfo{Forms[4],Specials[0]},
		PropertyInfo: PropertyInfo{300},

	}
	b.add(b20)
	b.GrowthDigitalBabies=DigitalBabies[:5]
	b.MatureDigitalBabies=DigitalBabies[5:10]
	b.CompleteDigitalBabies=DigitalBabies[10:15]
	b.UltimateDigitalBabies=DigitalBabies[15:20]
}




