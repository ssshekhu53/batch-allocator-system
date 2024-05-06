package main

import (
	adminPkg "batch-allocation-system/service/admin"
	studentPkg "batch-allocation-system/service/student"
	"fmt"
)

func main() {
	student := studentPkg.New()
	admin := adminPkg.New(student)

	id1 := student.Register("Akhilesh", "MALE")
	id2 := student.Register("Komal", "FEMALE")
	id3 := student.Register("Rajnish", "MALE")
	id4 := student.Register("Mayuri", "FEMALE")

	student.Enroll(id1, "IIT")
	student.Enroll(id2, "IIT")
	student.Enroll(id3, "NEET")
	student.Enroll(id4, "IIT")

	aID1 := admin.Register("Kamesh", "MALE")
	aID2 := admin.Register("M", "MALE")

	admin.CreateBatch(aID1, 3, "IIT", "Morning")
	admin.CreateBatch(aID2, 3, "IIT", "Morning")

	fmt.Println(admin.AllocateBatch(aID2, id2, "Gender Based"))
	fmt.Println(admin.AllocateBatch(aID2, id4, "Higher Capacity"))
	fmt.Println(admin.AllocateBatch(aID1, id1, "Higher Capacity"))
	fmt.Println(admin.AllocateBatch(aID1, id3, "Higher Capacity"))
}
