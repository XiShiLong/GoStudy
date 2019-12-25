# 函数的地位
函数是一等公民；参数、变量、返回值都可以是函数
# 正统 函数式编程
不可变性：不能有状态，只有常量和函数
# 闭包
```
func adder() func(int)int{
	sum:=0
	return func(v int)int {
		sum+=v
		return sum
	}
}

func main(){
	a:=adder()
	for i:=0;i<10;i++{
		fmt.Printf("0+1+...+ %d = %d \n",i,a(i))
	}
}```
adder()函数在return时不仅仅返回return后面的function，而是联合function和function涉及的function引用的变量sum形成一个闭包一起返回。
![闭包](https://github.com/aaaGoodMan/GoStudy/src/函数式编程/bibao.png)
