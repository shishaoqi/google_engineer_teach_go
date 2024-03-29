package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	var rlim syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim)
	if err != nil {
		fmt.Errorf("获取Rlimit报错 %s", err)
		os.Exit(1)
	}
	fmt.Printf("ENV RLIMIT_NOFILE : %+v\n", rlim)

	rlim.Max = 65535
	rlim.Cur = 24576
	fmt.Println("syscall.RLIMIT_NOFILE: ", syscall.RLIMIT_NOFILE);
	//err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
	if err != nil {
		fmt.Printf("设置Rlimit报错 %s\n", err)
		os.Exit(1)
	}
	//fmt.Printf("ENV RLIMIT_NOFILE : %+v\n", rlim)
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim)
	if err != nil {
		fmt.Printf("获取Rlimit报错 %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("222 ENV RLIMIT_NOFILE : %+v\n", rlim)
}
