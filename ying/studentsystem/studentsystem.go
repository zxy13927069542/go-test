package main

import "fmt"

//student结构体
type student struct {
	Id    int
	Name  string
	Age   int
	Class int
}

//student构造函数
func NewStudent(Id, Age, Class int, Name string) *student {
	return &student{
		Id:    Id,
		Name:  Name,
		Age:   Age,
		Class: Class,
	}
}

//展示所有学生信息
func ShowAll(stus []*student) {
	for i := 0; i < len(stus); i++ {
		fmt.Println(*stus[i])
	}
}

//添加学生
func AddStudent(stus *[]*student) {
	var stu student
	fmt.Println("请输入学生信息：")
	fmt.Print("Id:")
	fmt.Scan(&stu.Id)
	fmt.Print("Name:")
	fmt.Scan(&stu.Name)
	fmt.Print("Age:")
	fmt.Scan(&stu.Age)
	fmt.Print("Class:")
	fmt.Scan(&stu.Class)
	var flag int
	fmt.Println("请确认输入信息无误：", stu, "正确请输入1，错误请输入0")
	fmt.Scan(&flag)
	if flag == 1 {
		*stus = append(*stus, &stu)
		fmt.Println("添加成功！")
		return
	} else {
		fmt.Println("请重新选择并输入")
		return
	}

}

//编辑学生信息
func modify(stus *[]*student) {
	var Id int
	var stu *student
	var temp student
	ShowAll(*stus)
	fmt.Println("请选择要修改的学生Id")
	fmt.Scan(&Id)
	for i := 0; i < len(*stus); i++ {
		if (*stus)[i].Id == Id {
			stu = (*stus)[i]
		}
	}
	if stu == nil {
		fmt.Println("Id输入错误")
		return
	}
	fmt.Println("请输入学生信息：")
	fmt.Print("Id:")
	fmt.Scanln(&temp.Id)
	fmt.Print("Name:")
	fmt.Scanln(&temp.Name)
	fmt.Print("Age:")
	fmt.Scanln(&temp.Age)
	fmt.Print("Class:")
	fmt.Scanln(&temp.Class)
	var flag int
	fmt.Println("请确认输入信息无误：", temp, "正确请输入1，错误请输入0")
	fmt.Scan(&flag)
	if flag == 1 {
		stu.Age = temp.Age
		stu.Class = temp.Class
		stu.Id = temp.Id
		stu.Name = temp.Name
		return
	} else {
		fmt.Println("请重新选择并输入")
		return
	}

}

//删除学生
func Delete(stus *[]*student) {
	var Id int
	ShowAll(*stus)
	fmt.Println("请选择要删除的学生Id")
	fmt.Scan(&Id)
	for i := 0; i < len(*stus); i++ {
		if (*stus)[i].Id == Id {
			*stus = append((*stus)[0:i], (*stus)[i+1:]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Println("Id输入错误")
	return

}

func main() {
	var choice int
	students := make([]*student, 0, 200)
	for {
		fmt.Println("欢迎来到学生信息管理系统，请选择您要选择的功能：")
		fmt.Println("1.查看全部学生")
		fmt.Println("2.添加学生信息")
		fmt.Println("3.编辑学生信息")
		fmt.Println("4.删除学生信息")
		fmt.Println("5.退出")
		fmt.Println("请输入您要选择的功能:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			ShowAll(students)
		case 2:
			AddStudent(&students)
			ShowAll(students)
		case 3:
			modify(&students)
			ShowAll(students)
		case 4:
			Delete(&students)
			ShowAll(students)
		case 5:
			return
		default:
			fmt.Println("输入错误，请重新输入!")
		}
	}

}
