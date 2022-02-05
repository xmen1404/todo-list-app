package main

import (
	"fmt"
	"net/http"

	"database/sql"

	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type todoItem struct {
	TaskID     string `json:"taskid"`
	TaskName   string `json:"taskname"`
	TaskStatus bool   `json:"taskstatus"`
}

type todoListStruct struct {
	TodoList []todoItem `json:"todolist"`
}

var currentTodoList []todoItem = []todoItem{}
var const_db_file string = "db.json"
var db *sql.DB
var userInfo map[string]string

func stringToInt(stVal string) int {
	intVal, err := strconv.Atoi(stVal)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return intVal
}

func intToString(intVal int) string {
	return strconv.Itoa(intVal)
}

func loadTaskList(userId string) ([]todoItem, error) {
	var taskData []todoItem
	// id
	// userId
	// taskName
	// taskStatus

	row, err := db.Query("SELECT * FROM taskInfo WHERE userId == " + "'" + userId + "'")
	if err != nil {
		fmt.Println(err.Error())
		return []todoItem{}, err
	}

	var taskId, taskName string
	var taskStatus int

	for row.Next() {
		row.Scan(&taskId, &userId, &taskName, &taskStatus)
		taskStatusB := false
		if taskStatus == 1 {
			taskStatusB = true
		}
		taskData = append(taskData, todoItem{taskId, taskName, taskStatusB})
	}

	return taskData, err
}

// func saveDB(filename string, data []todoItem) error {
// 	ndata, err := json.MarshalIndent(todoListStruct{data}, "", " ")
// 	if err != nil {
// 		return err
// 	}

// 	return ioutil.WriteFile(filename, ndata, 0644)
// }

func getAuth(c *gin.Context) (string, error) { // AuthToken = userId
	authToken, errAuth := c.Cookie("authToken")
	if errAuth != nil {
		fmt.Println(errAuth.Error())
		return "", errAuth
	}

	_, exists := userInfo[authToken]
	if !exists {
		fmt.Println("access denied")
		return "", errAuth
	}
	return authToken, nil
}

func getTodoList(c *gin.Context) {
	authToken, errAuth := getAuth(c)
	if errAuth != nil {
		c.JSON(http.StatusForbidden, todoListStruct{})
		return
	}

	ndata, err := loadTaskList(authToken)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, todoListStruct{})
		return
	}
	c.IndentedJSON(http.StatusOK, todoListStruct{ndata})
}

func processDbStmt(statement string) error {
	stmt, err := db.Prepare(statement)
	if err == nil {
		stmt.Exec()
	}
	return err
}

func addTaskHandler(c *gin.Context) {
	authToken, errAuth := getAuth(c)
	if errAuth != nil {
		c.JSON(http.StatusForbidden, todoListStruct{})
		return
	}

	nTaskName := c.PostForm("taskname")
	nTaskId := uuid.New().String()
	nTaskStatus := "0"

	err := processDbStmt("INSERT INTO taskInfo VALUES (" + nTaskId + "," + authToken + "," + nTaskName + "," + nTaskStatus + ")")
	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusBadRequest, "Cannot add task")
		return
	}

	c.String(http.StatusCreated, "Added task successfully")
}

func removeTaskHandler(c *gin.Context) {
	authToken, errAuth := getAuth(c)
	if errAuth != nil {
		c.JSON(http.StatusForbidden, todoListStruct{})
		return
	}

	nTaskID := c.PostForm("taskid")
	err := processDbStmt("DELETE FROM taskInfo WHERE id == '" + nTaskID + "' AND userId == '" + authToken + "'")
	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusBadRequest, "Cannot delete task")
		return
	}

	c.String(http.StatusCreated, "Delete task successfully")
}

func changeStatusHandler(c *gin.Context) {
	authToken, errAuth := getAuth(c)
	if errAuth != nil {
		c.JSON(http.StatusForbidden, todoListStruct{})
		return
	}

	nTaskID := c.PostForm("taskid")

	err := processDbStmt("UPDATE taskInfo SET taskStatus = 1 - (SELECT taskStatus FROM taskInfo WHERE id == '" + nTaskID + "' AND userId == '" + authToken + "') " + "WHERE id == '" + nTaskID + "' AND userId == '" + authToken + "'")
	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusBadRequest, "Cannot delete task")
		return
	}
	c.String(http.StatusCreated, "Change task status successfully")
}

func main() {
	var err error = nil
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err = sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = processDbStmt("CREATE TABLE IF NOT EXISTS userInfo (id INT PRIMARY KEY, userName TEXT, password TEXT);")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = processDbStmt("CREATE TABLE IF NOT EXISTS taskInfo (id INT PRIMARY KEY, userId INT, taskName TEXT, taskStatus INT);")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/todo-list/get-task-list", getTodoList)
	router.POST("/todo-list/remove-task", removeTaskHandler)
	router.POST("/todo-list/add-task", addTaskHandler)
	router.POST("todo-list/change-task-status", changeStatusHandler)

	router.Run("localhost:8000")
}
