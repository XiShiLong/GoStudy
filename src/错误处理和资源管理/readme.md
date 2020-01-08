#1.defer
defer后面跟的语句是方法结束后必须执行的，不管是方法正常执行还是异常报错

#2.defer执行的顺序
func tryDefer(){
	defer fmt.Println(1) 
	defer fmt.Println(2)
	defer fmt.Println(3)
}
执行顺序是按栈的规则执行的即从后往前执行。3 2 1 

#自定义error
1.err=errors.New("there occored an unknow error")
2.type myerror struct {//实现error接口
  
  }
  
  func (myerror) Error() string {
  	return "there occored an unknow error"
  }
  
#panic 和 recover
 
##panic
*停止当前函数执行
*一直向上返回，执行每一层的defer
*如果没有遇见recover,程序退出

##recover
*尽在defer中调用
*获取panic的值
*如果无法处理，可以重新panic