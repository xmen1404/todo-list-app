package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoItem struct {
	TaskName   string `json:"taskname"`
	TaskStatus bool   `json:"taskstatus"`
}

type todoListStruct struct {
	TodoList []todoItem `json:"todolist"`
}

var currentTodoList []todoItem = []todoItem{}
var const_db_file string = "db.json"

func loadDB(filename string) ([]todoItem, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return []todoItem{}, err
	}

	ndata := todoListStruct{}
	err = json.Unmarshal([]byte(file), &ndata)
	return ndata.TodoList, err
}

func saveDB(filename string, data []todoItem) error {
	ndata, err := json.MarshalIndent(todoListStruct{data}, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, ndata, 0644)
}

func getTodoList(c *gin.Context) {
	ndata, err := loadDB(const_db_file)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, todoListStruct{})
		return
	}
	c.IndentedJSON(http.StatusOK, todoListStruct{ndata})
}

func addTaskHandler(c *gin.Context) {
	nTaskName := c.PostForm("taskname")

	ntodoItem := todoItem{nTaskName, false}
	currentTodoList = append(currentTodoList, ntodoItem)
	if err := saveDB(const_db_file, currentTodoList); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ntodoItem)
	}

	c.IndentedJSON(http.StatusCreated, ntodoItem)
}

func removeTaskHandler(c *gin.Context) {
	nTaskName := c.PostForm("taskname")
	nTodoList := []todoItem{}
	for _, item := range currentTodoList {
		if item.TaskName != nTaskName {
			nTodoList = append(nTodoList, item)
		}
	}
	currentTodoList = nTodoList
	saveDB(const_db_file, currentTodoList)

	c.IndentedJSON(http.StatusCreated, currentTodoList)
}

func changeStatusHandler(c *gin.Context) {
	nTaskName := c.PostForm("taskname")
	for idx, item := range currentTodoList {
		if item.TaskName == nTaskName {
			currentTodoList[idx].TaskStatus = !currentTodoList[idx].TaskStatus
		}
	}
	saveDB(const_db_file, currentTodoList)
	c.IndentedJSON(http.StatusCreated, currentTodoList)
}

func main() {
	var err error = nil
	currentTodoList, err = loadDB(const_db_file)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()
	router.GET("/todo-list", getTodoList)
	router.POST("/todo-list/remove-task", removeTaskHandler)
	router.POST("/todo-list/add-task", addTaskHandler)
	router.POST("todo-list/change-task-status", changeStatusHandler)

	router.Run("localhost:8080")
}
