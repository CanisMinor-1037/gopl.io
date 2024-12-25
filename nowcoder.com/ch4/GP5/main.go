package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param a int整型 变量a
 * @param b int整型 变量b
 * @return bool布尔型一维数组
*/
func equal( a int ,  b int ) []bool {
    // write code here
    var r []bool
    r1 := &a == &b
    r2 := a == b
    r = append(r, r1)
    r = append(r, r2)
    return r
}