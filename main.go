package main

//exercise 1
/* 

import "fmt"

func main() {
	var count int
	var totalScore float64

	fmt.Print("请输入学生人数: ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		var name string
		var score float64

		fmt.Printf("请输入第%d个学生的姓名: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("请输入%s的成绩: ", name)
		fmt.Scan(&score)

		var level string
		if score >= 90 {
				level = "优秀"
		} else if score >= 70 {
				level = "良好"
		} else if score >= 60 {
				level = "及格"
		} else {
				level = "不合格"
		}

		fmt.Printf("学生%s的成绩%.1f, 等级：%s\n\n", name, score, level)
		totalScore += score
	}

	average := totalScore / float64(count)
	fmt.Printf("班级平均分：%.2f\n", average)
}
*/

//exercise 2
/*

import "fmt"

type Student struct {
	ID int
	Name string
	Score float64
}

var students []Student
var studentMap map[int]Student

func main() {
	studentMap = make(map[int]Student)

	addStudent(1, "XiaoMing", 85.5)
	addStudent(2, "XiaoHong", 92.0)
	addStudent(3, "XiaoGang", 78.5)

	displayAllStudents()

	if student, found := findStudent(2); found {
		fmt.Printf("Student Name:%s, Score: %.1f\n", student.Name, student.Score)
	} else {
		fmt.Printf("Student Not Found")
	}

	fmt.Printf("Class Average: %.2f\n", calculateAverage())
}

func addStudent(id int, name string, score float64) {
	student := Student {
		ID:    id,
		Name:  name,
		Score: score,
	}
	students = append(students, student)
	studentMap[id] = student
}

func displayAllStudents() {
	fmt.Printf("All Student Information")
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Score: %.1f\n", student.ID, student.Name, student.Score)
	}
	fmt.Println()
}

func findStudent(id int) (Student, bool) {
	student, exists := studentMap[id]
	return student,exists
}

func calculateAverage() float64 {
	if len(students) == 0 {
		return 0
	}

	var total float64
	for _, student := range students {
		total += student.Score
	}
	return total / float64(len(students))
}
*/

//exercise 3
/*
import (
	"fmt"
	"os"
)

type FileOperator interface {
	Save(filename string) error
	Load(filename string) error
}

type StudentManager struct {
	Students []Student
}

func (sm *StudentManager) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Create file fail: %v", err)
	}
	defer file.Close()

	for _, student := range sm.Students{
		_, err := fmt.Fprintf(file, "%d,%s,%.1f\n", student.ID, student.Name, student.Score)
		if err! = nil {
			return fmt.Errorf("Write file fail: %v", err)
		}
	}
	return nil
}

func (sm *StudentManager) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Load file fail: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sm.Students = nil

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) != 3 {
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		name := parts[1]

		score, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			continue
		}

		student := Student {
			ID:    id,
			Name:  name,
			Score: score,
		}
		sm.Students = append(sm.Students, student)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("load file fail: %v", err)
	}

	sm.IDMAP = make(map[int]*Student)
	for i := range sm.Students {
		sm.IDMAP[sm.Students[i].ID] = &sm.Students[i]
	}

	return nil
}
	*/

import (
	"log"
	"student-score/app"
)

func main () {
	application, err := app.NewApplication("config.json")
	if err != nil {
		log.Fatalf("启动失败: %v", err)
	}

	if err := application.start(); err != nil {
		log.Fatalf("运行失败: %v", err)
	}
}