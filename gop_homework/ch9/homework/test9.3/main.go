//练习 9.3： 扩展Func类型和(*Memo).Get方法，支持调用方提供一个可选的done channel，
// 使其具备通过该channel来取消整个操作的能力(§8.9)。一个被取消了的Func的调用结果不应该被缓存。
package main
