package main

import (
	"fmt"
)

func income(balance float64) (float64 ,float64, string) {
	fmt.Print("本次收入金额：")
	num := 0.0
	fmt.Scanln(&num)
	fmt.Print("本次收入说明：")
	str := ""
	fmt.Scanln(&str)
	fmt.Println("----------登记完成----------")
	return balance+num,num,str
}

func expense(balance float64) (float64 ,float64, string){
	fmt.Print("本次支出金额：")
	num := 0.0
	fmt.Scanln(&num)
	if num>balance{
		fmt.Println("余额不足")
		return balance,0.0,"余额不足，无效支出"
	}
	fmt.Print("本次支出说明：")
	str := ""
	fmt.Scanln(&str)
	fmt.Println("----------登记完成----------")
	return balance-num,num,str
}
func main(){
	//余额
	balance := 10000.0
	//收支的金额
	money := 0.0
	//收支详情
	details := "收支\t账户金额\t收支金额\t说 明"
	var str string
	//选择
	key := 0
	flag := false
	lable1:
	for{
		fmt.Println("--------家庭收支记账软件--------")
		fmt.Println("         1 收支明细            ")
		fmt.Println("         2 登记收入            ")
		fmt.Println("         3 登记支出            ")
		fmt.Println("         4 退    出            ")
		fmt.Print("         请选择(1-4):")
		fmt.Scanln(&key)
		switch key{
		case 1:
			fmt.Println("---------当前收支明细---------")
			if flag{
				fmt.Println(details)
			}else{
				fmt.Println("当前没有收支明细……")
			}
		case 2:
			balance ,money, str = income(balance)
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",balance,money,str)
			flag = true
		case 3:
			balance ,money, str = expense(balance)
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v",balance,money,str)
			flag = true
		case 4:
			fmt.Println("你确定要退出吗?y/n")
			choice := ""
			for{
				fmt.Scanln(&choice)
				if choice =="y"||choice == "n"{
					break
				}
				fmt.Println("输入有误，请重新输入 y/n")
			}
			if choice == "y"{
				break lable1
			}
		default:
			fmt.Println("请输入正确的选项")
			continue
		}
	}
	fmt.Println("已退出")
}