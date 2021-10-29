/*******
* @Author:qingmeng
* @Description:
* @File:state
* @Date2021/10/25
 */

package info

var (
	Forms=[]Form{{0,"成长期"},{1,"成熟期"},{2,"完全体"},{3,"究极体"},{4,"超究极体"}}
	Specials=[]Special{{0,"正常"},{1,"睡眠"},{2,"灼烧"}}
)

type State interface {
	ChangeForm(n int) //改变形态
	ChangeState()     //改变状态
	GetForm()Form     //获取形态
	ChangeSpecial(n int)		//改变特殊状态
	GetSpecial()int 	//获取特殊值
}

//形态
type Form struct {
	id	int
	name string
}

//特殊状态
type Special struct {
	id int
	name string
}

//状态
type StateInfo struct {
	form    Form
	special Special
}

func (s *StateInfo) GetSpecial() int {
	return s.special.id
}

func (s *StateInfo) ChangeSpecial(n int ) {
	s.special =Specials[n]
}

func (s *StateInfo) GetForm()Form {
	return s.form
}

func (s *StateInfo) ChangeForm(n int) {
	s.form=Forms[n]

}

func (s *StateInfo) ChangeState() {

}

