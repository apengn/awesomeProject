package main

import "awesomeProject/cmd2"

type S struct {
}

func main() {
	cmd2.Execute()
}

// func main() {
// 	var x S
// 	y := &x
// 	_ = identity(y)
//
// }
// func identity(z *S) *S {
// 	return z
// }

//go  build  -ldflags "-X awesomeProject/version.gitCommit=`git rev-parse --short HEAD` -X awesomeProject/version.buildTime=`date +%Y%m%d-%H%M%S`"
//
// func main() {
// 	fmt.Println(Signal())
//
// 	time.Sleep(3 * time.Second)
// }
//
// func Signal() int {
//
// 	ch := make(chan os.Signal)
//
// 	// 监听 系统信号  如果是interrupt  kill    则退出 进程
// 	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
//
// 	// 停止向channel 转发信号  下面这个不会执行
// 	//signal.Stop(ch)
//
// 	go func() {
// 		s := <-ch
// 		fmt.Println("=====", s)
// 		s = <-ch
// 		fmt.Println("exit  :", s)
// 	}()
// 	return 9
// }

// func main() {
// 	//var validID = regexp.MustCompile(`(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?`)
//
// 	var validID = regexp.MustCompile("^[a-zA-Z0-9][-a-zA-Z0-9]*[a-zA-Z0-9]$")
//
// 	// var validID = regexp.MustCompile(`^[-A-Za-z0-9_.]{3,16}`)
//
// 	fmt.Println(validID.MatchString("fsTtT"))
//
// }

// func init() {
// 	//以时间作为初始化种子
// 	rand.Seed(time.Now().UnixNano())
// }
// func main() {
//
// 	for i := 0; i < 200; i++ {
// 		// rand.
// 		a := rand.Intn(3)
//
// 		fmt.Println(a)
// 	}
//
// }
