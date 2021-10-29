/*******
* @Author:qingmeng
* @Description:
* @File:ReceiverDemo
* @Date2021/10/24
 */

package main

import "fmt"

func Receivers(v interface{}){
	switch v.(type) {
	case byte:
		fmt.Println("This is byte")
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case float64:
		fmt.Println("This is float")
	case bool:
		fmt.Println("This is bool")
	default:
		fmt.Println("Sorry,unkown Type")
	}
}

func main() {
	Receivers("你好")
	Receivers(72183)
	Receivers(false)
}