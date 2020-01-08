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

##两种实现的区别
值实现
var a animal 
a=dog{}
指针实现
var a animal
a=&dog{}

#接口组合

type IGet interface{
    Get(url string)string
}

type IPut interface{
    Put(url string)
}

组合接口
type IGetAndPut interface{
    IGet
    IPut
}

#obj.(type)返回的值
例如
a:=map[string]interface{}{
		"name":"xishilong",
		"age":20,
	}

c,ok:=a["name"].(string)
c是string类型的字符串"xishilong",ok是bool类型的，如果转化成功就返回true
