#结构体实现interface
1.结构体只用实现interface中的方法即可实现interface
2.由两种实现方法
type animal interface{
    eat()
}

type dog struct{

}

func(d dog) eat(){//值实现
}
或
func(d *dog) eat(){//指针实现
}

#两种实现的区别
值实现
var a animal 
a=dog{}
指针实现
var a animal
a=&dog{}