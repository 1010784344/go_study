package models

import "gin_demo/dao"


// Todo model
type Todo struct {

	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`

}


func CreateATodo(todo *Todo) (err error){
	err = dao.DB.Create(todo).Error
	return

}


func GetAllTodo() (todolist []*Todo,err error){
	err = dao.DB.Find(&todolist).Error
	if err != nil{

		return nil,err
	}else {
		return
	}

}


func GetATodo(id string) (todo *Todo,err error){
	todo = new(Todo)

	err = dao.DB.Where("id=?",id).First(todo).Error

	if err != nil{

		return nil,err
	}else {
		return
	}

}


func UpdateATodo(todo *Todo)(err error){

	err = dao.DB.Save(&todo).Error
	return

}


func DeleteATodo(id string)(err error){

	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}


