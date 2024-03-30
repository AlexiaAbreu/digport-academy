package main

import "fmt"

func main() {
	contacts := make(map[string]int)

	contacts["Alexia"] = 51993673098
	contacts["Carol"] = 51993673097
	contacts["Viviane"] = 51993673099

	fmt.Println(contacts)

	contactsTwo := map[string]int{
		"Julia":  51993673098,
		"Camila": 51993673097,
		"Vivi":   51993673099,
	}

	fmt.Println(contactsTwo)

	for _, value := range contacts {
		contacts.
	}

}
