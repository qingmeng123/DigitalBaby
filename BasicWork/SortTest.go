/*******
* @Author:qingmeng
* @Description:
* @File:SortTest
* @Date2021/10/24
 */

package main

import (
	"fmt"

)

type Interface interface {
	Len()int
	Less(i,j int)bool
	Swap(i,j int)
}

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	a:=[]byte(m[i])
	b:=[]byte(m[j])
	flag:=false
	var res bool
	for k := 0; k < len(a)&&k< len(b); k++ {
		if a[k]<b[k]{
			res= true
			break
		}else if a[k]>b[k]{
			res= false
			break
		} else if k==len(a)-1||k==len(b)-1{
			flag=true
		}
	}
	if flag{
		res= len(a)< len(b)
	}
		return res
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func Sort(s Interface){
	lengh:=s.Len()
	for i := 0; i < lengh; i++ {

		for j :=0 ; j < lengh-1; j++ {
			if s.Less(j+1,j){
				s.Swap(j+1,j)
			}
		}

	}
}

func main() {
	strs := MyStringList{
		"hello",
		"thanks",
		"asleep",
		"hello hello",
	}

	Sort(strs)
	for _, v := range strs {
		fmt.Printf("%s\n", v)
	}
}