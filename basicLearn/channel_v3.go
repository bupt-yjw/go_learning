package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}
func say2(s string, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	ch <- 0
	close(ch) //关闭通道并不会丢失里面的数据，只是让读取通道数据的时候不会读完之后一直阻塞等待新数据写入
}

func main() {
	ch := make(chan int)
	go say2("world", ch)
	say("hello")
	fmt.Println(<-ch) //主函数会等待信道的值才会结束。
}

//注释掉的方法，在执行中会发现，say2方法并没有执行完成。这是因为goroutine 还没来得及跑完 5 次的时候，主函数已经退出了。通过引入信道，可以解决这个问题。

/**
package main
import (
    "fmt"
    "time"
)
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s, (i+1)*100)
    }
}
func say2(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(150 * time.Millisecond)
        fmt.Println(s, (i+1)*150)
    }
}
func main() {
    go say2("world")
    say("hello")
}
*/
