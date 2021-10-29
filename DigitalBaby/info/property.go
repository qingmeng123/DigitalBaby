/*******
* @Author:qingmeng
* @Description:
* @File:property
* @Date2021/10/25
 */

package info


//属性
type Property interface {
	BabyDp()		int
}

type PropertyInfo struct {
	dp    int				//防御
}

func (p PropertyInfo) BabyDp()int {
	return p.dp
}


