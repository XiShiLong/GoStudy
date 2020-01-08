# 表格测试的语法
表格测试实现了将测试数据和测试方法分离  
假如我们要测试方法:  
```
func calcTriangle(a, b int) int {  
	var c int  
	c = int(math.Sqrt(float64(a*a + b*b)))  
	return c  
}  
```  
测试代码:  
```    
func TestTriangle(t *testing.T){  
	tests:=[]struct{a,b,c int} {  
		{3,4,5},  
		{5,12,13},  
		{8,15,17},  
		{12,35,37},  
		{30000,40000,50000},  
	}  

	for _,tt:=range tests{  
		if actual:=calcTriangle(tt.b,tt.b);actual!=tt.c{  
			t.Errorf("calcTriangle(%d, %d); got %d ; expected %d",tt.a,tt.b,actual,tt.c)  
		}  
	}  
}  
```  


# 测试步骤
1.在被测试方法所在的同级目录下建一个go文件，命名格式为 ```测试方法名称_test```  
2.建立测试方发，方法名称为必须以  Test开头，且传入参数是  *testing.T 类型  
3.建立测试数据集,注意第一个{}声明数据类型，第二个{}才是与数据类型对应的数据
4.循环便利数据集对测试方法进行测试  

# 测试命令
cd 到测试文件所在目录
go test 运行测试代码
go test -coverprofile=c.out 将代码覆盖率用打印到c.out文件中
go tool cover -html=c.out 查看c.out文件

# 性能测试
文件命名和测试文件一样，函数命名开头必须以  Benchmark 开头  
```
func BenchmarkSubstr(b *testing.B){
	s:="黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans:=8
    //b.ResetTimer()
	for i:=0;i<b.N;i++{
		actual:=lengthOfNonRepeatingSubStr(s)
		if actual!=ans{
			b.Errorf("got %d for input %s; expected %d",actual,s,ans)
		}
	}
}
```
使用 go test -bench .  可以命令执行性能测试,运行结果如下  
```
goos: windows
goarch: 386
pkg: 测试与性能调优/nonrepeatingsubstr
BenchmarkSubstr-8        1000000              1579 ns/op
PASS
ok      测试与性能调优/nonrepeatingsubstr       2.843
```
结果表示运行了1000000次，每次耗时1579纳秒，共计花费2.843秒

`如果你不想从头就开始计算运行时间，那么你可以使用b.ResetTimer()加到你不想计算时间的代码后面`

#cpu使用率测试
go test -bench . -cpuprofile cpu.out 会生成一个cpu.out的文件
go tool pprof cpu.out 会进入命令行，输入help可以查看所有命令。输入web就可以看到cpu消耗图，方块面积越大表示占cpu越多。
