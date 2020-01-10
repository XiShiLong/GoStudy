# 接收channel方式  
c:=make(chan int)
c<-2
c<-3

1.n,ok:=<-c 
n是取出来的数据，ok如果取出来值就是true否则就是false
2. range
for n:=range(c){

}

注意！ 
和for n:=range(c){}同一个函数的在其后面的代码将不会执行  
例如  
func test(c chan int){  
	for n:=range c{  
		fmt.Printf("%d",n)  
	}  
	fmt.Println("hahahha")  
}  
fmt.Println("hahahha")将不再被执行  


n就是value,如果c中的数据取完了for循环结束  

# 关闭chanl

如果发送方不再发送数据可以调用close函数，这时接收方街道的是 对应类型的空类型，比如int是0，string 是 ""


# 利用chanl等待

func work(c chan bool){
    /*
    do something....
    */
    c<-true
}

func main(){
    chs:=make([]chan bool,10)
    for i:=0;i，10；i++{
        chs[i]=make(chan bool,1)
    }
    for i:=0;i<10;i++{
        go work(chs[i])
    }
    //等待10个共工作做完
    for i:=0;i<10;i++{
        <-chs[i]
    }
}

# 利用 sync.WaitGroup

func work(wg *sync.WaitGroup){
    /*
    do something....
    */
    //每个携程做完就Done操作后wg中的add初始值-1
    wg.Done()
}

func main(){
    chs:=make([]chan bool,10)
    
    //开启10个携程
    for i:=0;i<10;i++{
        go work(chs[i])
    }
    wg.Add(10)
    //10个携程做完后waite函数就不再阻塞主线程
    wg.Wait()
}