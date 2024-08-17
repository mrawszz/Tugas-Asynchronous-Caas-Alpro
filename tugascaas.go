package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Member struct {
	Name  string
	Tasks []Task
}

type Task struct {
	Name        string
	Description string
	DueDate     time.Time
}

var members []Member

func main() {
	for {
		fmt.Println("\n========== Menu Utama ==========")
		fmt.Println("1. Tambahkan member")
		fmt.Println("2. Lihat daftar member")
		fmt.Println("3. Hapus member")
		fmt.Print("Masukkan pilihan menu (angka): ")

		choice := input()
		switch choice {
		case "1":
			addMember()
		case "2":
			viewMembers()
		case "3":
			deleteMember()
		default:
			fmt.Println("Pilihan tidak valid. Mohon masukan pilihan yang tersedia.")
		}
	}
}

func addMember() {
	fmt.Print("Masukkan nama member: ")
	name := input()
	members = append(members, Member{Name: name})
	fmt.Println("Member berhasil ditambahkan.")
}

func viewMembers() {
	if len(members) == 0 {
		fmt.Println("Belum ada member yang terdaftar.")
		return
	}

	fmt.Println("Daftar Member:")
	for i, member := range members {
		fmt.Printf("%d. %s\n", i+1, member.Name)
	}

	fmt.Print("Masukkan pilihan member untuk add task (masukkan 00 jika ingin kembali): ")
	choice := input()

	if choice == "00" {
		return
	}

	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(members) {
		fmt.Println("Pilihan tidak valid. Mohon masukan pilihan yang tersedia.")
		return
	}

	memberMenu(&members[index-1])
}

func memberMenu(member *Member) {
	for {
		fmt.Printf("\n======= Menu ======= \nmember: %s\n", member.Name)
		fmt.Println("1. Tambahkan task")
		fmt.Println("2. Lihat daftar task")
		fmt.Println("3. Edit task")
		fmt.Println("4. Hapus task")
		fmt.Println("5. Kembali ke menu sebelumnya")
		fmt.Print("Masukkan pilihan: ")

		choice := input()
		switch choice {
		case "1":
			addTask(member)
		case "2":
			viewTasks(member)
		case "3":
			editTask(member)
		case "4":
			deleteTask(member)
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid. Mohon masukan pilihan yang tersedia.")
		}
	}
}

func addTask(member *Member) {
	fmt.Print("Nama Task: ")
	taskName := input()

	fmt.Print("Deskripsi Task: ")
	description := input()

	fmt.Print("Set Due Date (format dd-mm-yyyy): ")
	dueDateStr := input()

	dueDate, err := time.Parse("02-01-2006", dueDateStr)
	if err != nil {
		fmt.Println("Format tanggal tidak valid. Mohon masukkan format dd-mm-yyy.")
		return
	}

	member.Tasks = append(member.Tasks, Task{Name: taskName, Description: description, DueDate: dueDate})
	fmt.Println("Task berhasil disimpan.")
}

func viewTasks(member *Member) {
	if len(member.Tasks) == 0 {
		fmt.Println("Tidak ada task.")
		return
	}

	fmt.Println("Daftar Task:")
	for i, task := range member.Tasks {
		fmt.Printf("%d. %s - %s (Due: %s)\n", i+1, task.Name, task.Description, task.DueDate.Format("02-01-2006"))
	}
}

func editTask(member *Member) {
	viewTasks(member)
	if len(member.Tasks) == 0 {
		return
	}

	fmt.Print("Masukkan task yang ingin diedit: ")
	choice := input()

	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(member.Tasks) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	task := &member.Tasks[index-1]

	fmt.Printf("Edit nama task (%s): ", task.Name)
	task.Name = input()

	fmt.Printf("Edit deskripsi task (%s): ", task.Description)
	task.Description = input()

	fmt.Printf("Edit due date task (%s) (format dd-mm-yyyy): ", task.DueDate.Format("02-01-2006"))
	dueDateStr := input()

	dueDate, err := time.Parse("02-01-2006", dueDateStr)
	if err != nil {
		fmt.Println("Format tanggal tidak valid.")
		return
	}
	task.DueDate = dueDate

	fmt.Println("Task berhasil diedit.")
}

func deleteTask(member *Member) {
	viewTasks(member)
	if len(member.Tasks) == 0 {
		return
	}

	fmt.Print("Masukkan task yang ingin dihapus: ")
	choice := input()

	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(member.Tasks) {
		fmt.Println("Pilihan tidak valid. Masukan pilihan yang tersedia.")
		return
	}

	member.Tasks = append(member.Tasks[:index-1], member.Tasks[index:]...)
	fmt.Println("Task berhasil dihapus.")
}

func deleteMember() {
	if len(members) == 0 {
		fmt.Println("Belum ada member yang terdaftar.")
		return
	}

	fmt.Println("Daftar Member:")
	for i, member := range members {
		fmt.Printf("%d. %s\n", i+1, member.Name)
	}

	fmt.Print("Masukkan nomor member yang ingin dihapus (masukkan 00 jika ingin kembali): ")
	choice := input()

	if choice == "00" {
		return
	}

	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(members) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	members = append(members[:index-1], members[index:]...)
	fmt.Println("Member berhasil dihapus.")
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
