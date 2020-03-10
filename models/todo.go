package models

import "gin_gorm_demo/dao"

// 表名默认为结构体名称的复数："todos"
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// CreateATodo 创建一个待办事项
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

// GetAllTodo 查询所有待办事项的列表
func GetAllTodo(todoList *[]Todo) (err error) {
	err = dao.DB.Find(todoList).Error
	return
}

// GetATodo 根据id查询一个待办事项
func GetATodo(todo *Todo, id string) (err error) {
	err = dao.DB.Where("id = ?", id).First(todo).Error
	return
}

// UpdateATodo 更新一个待办事项
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

// DeleteATodo 删除一个待办事项
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return

}
