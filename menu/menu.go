package menu

import "fmt"

type Menu struct {
	listMenu []string
}

func (m *Menu) AddMenu(menu string) {
	m.listMenu = append(m.listMenu, menu)
}
func (m *Menu) InitMenu() {
	m.AddMenu("1. Create employees")
	m.AddMenu("2. Update employees")
	m.AddMenu("3. Delete employees")
	m.AddMenu("4. GetByID employees")
	m.AddMenu("5. Get all employees")
	m.AddMenu("6. Exit")
}

func (m *Menu) DisplayMenu() {
	for _, menu := range m.listMenu {
		fmt.Println(menu)
	}
}
