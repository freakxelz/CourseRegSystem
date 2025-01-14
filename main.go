package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	Course    string
	Age       int
}

var students []Student

func studentRegister(firstName string, lastName string, course string, age int) {
	if age < 18 {
		fmt.Println("siswa harus berusia 18 tahun ke atas")
		return
	}
	student := Student{
		FirstName: firstName,
		LastName:  lastName,
		Course:    course,
		Age:       age,
	}

	students = append(students, student)
	fmt.Println(firstName, "berhsail didaftarkan")

}

func listStudent() {
	if len(students) == 0 {
		fmt.Println("Belum ada siswa yang terdaftar!")
		return
	}
	fmt.Println("Berikut ini adalah siswa yang sudah terdaftar:")
	for i, s := range students {
		fmt.Printf("%d. Nama: %s %s, Umur: %d, Kursus: %s\n", i+1, s.FirstName, s.LastName, s.Age, s.Course)
	}
}

func main() {
	for {
		fmt.Println("1. Mendaftarkan Siswa")
		fmt.Println("2. List Daftar Siswa")
		fmt.Println("3. Keluar")
		fmt.Println("Masukkan Menu Disini: ")

		var menu int
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			var firstName, lastName, course string
			var age int
			fmt.Print("Masukkan nama Depan peserta: ")
			fmt.Scanln(&firstName)
			fmt.Print("Masukkan nama Belakang peserta: ")
			fmt.Scanln(&lastName)
			fmt.Print("Masukkan umur peserta: ")
			fmt.Scanln(&age)
			fmt.Print("Masukkan jenis kursus: ")
			fmt.Scanln(&course)
			studentRegister(firstName, lastName, course, age)
		case 2:
			listStudent()
		case 3:
			fmt.Println("Terima kasih sudah berkunjung!")
			return
		default:
			fmt.Println("Menu yang kamu masukkan tidak valid!")
		}
	}
}
