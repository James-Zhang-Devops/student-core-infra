package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"student-score/model"
	"student-score/service"
)

type StudentController struct {
	service service.StudentService
}

func NewStudentController(service service.StudentService) *StudentController {
	return &StudentController{
		service: service,
	}
}

func (c *StudentController) Run() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		c.showMenu()

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1": 
			c.addStudent(scanner)
		case "2":
			c.listStudents()
		case "3":
			c.findStudent(scanner)
		case "4":
			fmt.Println("Exit System...")
			return nil
		default:
			fmt.Println("无效选择，请重新输入!")
		}
	}
	return nil
}

func (c *StudentController) showMenu() {
    fmt.Println("\n=== 学生管理系统 ===")
    fmt.Println("1. 添加学生")
    fmt.Println("2. 查看所有学生") 
    fmt.Println("3. 查找学生")
    fmt.Println("4. 退出")
    fmt.Print("请选择操作: ")	
}

func (c *StudentController) addStudent(scanner *bufio.Scanner) {
	fmt.Print("请输入学生姓名：")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	fmt.Print("请输入学生成绩：")
	scanner.Scan()
	scoreStr := strings.TrimSpace(scanner.Text())

	score, err := strconv.ParseFloat(scoreStr, 64)
	if err != nil {
		fmt.Println("成绩格式错误")
		return
	}

	student := model.Student{
		Name:  name,
		Score: score,
	}

	if err := c.service.AddStudent(student); err != nil {
		fmt.Printf("添加学生失败：%v\n", err)
	} else {
		fmt.Println("添加学生成功")
	}
}

func (c *StudentController) listStudents() {
	students, err := c.service.ListStudent()
	if err != nil {
		fmt.Printf("获取学生列表失败：%v\n", err)
		return
	}

	if len(students) == 0 {
		fmt.Println("没有学生记录")
		return
	}

	fmt.Println("\n学生列表:")
	fmt.Println("ID\t姓名\t成绩\t等级")
	for _, student := range students {
		level := c.service.GetGradeLevel(student.Score)
		fmt.Printf("%d\t%s\t%.1f\t%s\n", student.ID, student.Name, student.Score, level)
	}
}

func (c *StudentController) findStudent(scanner *bufio.Scanner) {
	fmt.Print("请输入学生ID：")
	scanner.Scan()
	idStr := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID格式错误")
		return
	}

	student, err := c.service.GetStudent(id)
	if err != nil {
		fmt.Printf("查找学生失败：%v\n", err)
		return
	}

	if student == nil {
		fmt.Println("未找到该学生!")
		return
	}

	level := c.service.GetGradeLevel(student.Score)
	fmt.Printf("找到学生: ID=%d, 姓名=%s, 成绩=%.1f, 等级=%s\n", student.ID, student.Name, student.Score, level)
}