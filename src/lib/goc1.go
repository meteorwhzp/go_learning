/*
 提供的常用库，有一些常用的方法，方便使用
*/
package lib

import "encoding/json"

//一个加法实现
//返回a+b的值
func Add(n1, n2 int) int {
	json.Marshal(nil)
	return n1 + n2
}
