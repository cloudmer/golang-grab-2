package main

import (
	"runtime"
	"time"
	"xmn_2/core/algorithm/ssc"
	_ "xmn_2/core/model"
	"log"
	"os"
	"xmn_2/core/algorithm/shishicai/CustomPackage"
	"xmn_2/core/mail"
)

func main(){
	log.Println("服务启动中．．．　进程ID:", os.Getpid())
	mail.SendMail("邮件测试", "邮件测试")
	os.Exit(0)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for {
		select {
		case <-time.After(1 * time.Minute):
			//todo
			// 时时彩　包含数据包　算法　邮件报警
			go ssc.Contain()
			// 时时彩 连号 算法 邮件报警
			go ssc.Consecutive()
			// 时时彩 连续AB表报警
			go ssc.ContailMultiple()
			// 时时彩 AB包 自定义A包周期 报警
			go CustomPackage.Calculation()
		}
	}
}
