// 这个示例程序展示如何将一个类型嵌入另一个类型，以及
// 内部类型和外部类型之间的关系
package main

import (
	"fmt"
)

// notifier 是一个定义了
// 通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// 通过 user 类型值的指针
// 调用的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin 代表一个拥有权限的管理员用户
type admin struct {
	user  // 嵌入类型
	level string
}

// 通过 admin 类型值的指针
// 调用的方法
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// main 是应用程序的入口
func main() {
	// 创建一个 admin 用户
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 给 admin 用户发送一个通知，就近原则
	sendNotification(&ad)
	// 我们可以直接访问内部类型的方法
	ad.user.notify()

	// 内部类型的方法没有被提升
	ad.notify()
}

// sendNotification 接受一个实现了 notifier 接口的值
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}
