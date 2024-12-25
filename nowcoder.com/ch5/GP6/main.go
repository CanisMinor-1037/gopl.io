package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param s string字符串一维数组 
 * @return string字符串
*/
func join( s []string ) string {
    // write code here
    var result string
    for _, str := range s {
        result += str
    }
    return result
}