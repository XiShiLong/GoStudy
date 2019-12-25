package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*------------------------------------*/
func adder() func(int)int{
	sum:=0
	return func(v int)int {
		sum+=v
		return sum
	}
}
/*------------------------------------*/
//斐波拉且数列
func fibonaqie() func() int{
	a,b:=0,1
	return func()int{
		a,b=b,a+b
		return a
	}
}


/*------------------------------------*/
//接口编程
func fibonaqie2() intGen{
	a,b:=0,1
	return func()int{
		a,b=b,a+b
		return a
	}
}
//定义返回类型
type intGen func()int
//实现Read接口
func (obj intGen)Read(p []byte)(n int,err error){
	next:=obj()
	if next>1000{
		return 0,io.EOF
	}
	s:=fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	scanner:=bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main(){
	/*a:=adder()
	for i:=0;i<10;i++{
		fmt.Printf("0+1+...+ %d = %d \n",i,a(i))
	}*/
	/*f:=fibonaqie()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())*/
	f:=fibonaqie2()
	printFileContents(f)

}




