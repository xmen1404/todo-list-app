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

// var userInfo map[string]string // temporary userId tracker

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

	row, err := db.Query("SELECT * FROM taskInfo WHERE userId == '" + userId + "';")
	if err != nil {
		fmt.Println(err.Error(), "load task list error")
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
		fmt.Println("checking loadTaskList: ", taskId, taskName, taskStatusB)
		taskData = append(taskData, todoItem{taskId, taskName, taskStatusB})
	}

	return taskData, err
}

func getAuth(c *gin.Context) (string, error) { // AuthToken = userId
	authToken, errAuth := c.Cookie("authToken")
	fmt.Println(authToken)
	if errAuth != nil {
		fmt.Println(errAuth.Error(), "cookie error")
		return "", errAuth
	}

	_, err := db.Query("SELECT * FROM userInfo WHERE id == '" + authToken + "';")
	if err != nil {
		fmt.Println(err.Error(), "error in query userInfo")
		return "", errAuth
	}
	// if rows == nil {
	// 	fmt.Println("access denied")

	// }
	return authToken, nil
}

func getTodoList(c *gin.Context) {
	authToken, errAuth := getAuth(c)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
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
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if errAuth != nil {
		c.JSON(http.StatusForbidden, todoListStruct{})
		return
	}

	nTaskName := c.PostForm("taskname")
	nTaskId := uuid.New().String()
	nTaskStatus := "0"

	err := processDbStmt("INSERT INTO taskInfo VALUES ('" + nTaskId + "','" + authToken + "','" + nTaskName + "','" + nTaskStatus + "');")
	if err != nil {
		fmt.Println(err.Error(), "error with inserting to db")
		c.String(http.StatusBadRequest, "Cannot add task")
		return
	}

	c.String(http.StatusCreated, "Added task successfully")
}

func removeTaskHandler(c *gin.Context) {
	authToken, errAuth := getAuth(c)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if errAuth != nil {
		c.JSON(http.StatusForbidden, todoListStruct{})
		return
	}

	nTaskID := c.PostForm("taskid")
	err := processDbStmt("DELETE FROM taskInfo WHERE id == '" + nTaskID + "' AND userId == '" + authToken + "'")
	if err != nil {
		fmt.Println(err.Error(), "error with deleting from db")
		c.String(http.StatusBadRequest, "Cannot delete task")
		return
	}

	c.String(http.StatusCreated, "Delete task successfully")
}

func changeStatusHandler(c *gin.Context) {
	authToken, errAuth := getAuth(c)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if errAuth != nil {
		c.JSON(http.StatusForbidden, todoListStruct{})
		return
	}

	nTaskID := c.PostForm("taskid")

	err := processDbStmt("UPDATE taskInfo SET taskStatus = 1 - (SELECT taskStatus FROM taskInfo WHERE id == '" + nTaskID + "' AND userId == '" + authToken + "') " + "WHERE id == '" + nTaskID + "' AND userId == '" + authToken + "'")
	if err != nil {
		fmt.Println(err.Error(), "error with updating db")
		c.String(http.StatusBadRequest, "Cannot delete task")
		return
	}
	c.String(http.StatusCreated, "Change task status successfully")
}

func loginHandler(c *gin.Context) {
	fmt.Println("passed here")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	username := c.PostForm("username")
	password := c.PostForm("password")

	rows, err := db.Query("SELECT * FROM userInfo WHERE username == " + "'" + username + "' AND password == '" + password + "';")
	if err != nil {
		fmt.Println(err.Error(), "error with Select command")
		c.String(http.StatusBadRequest, "Login failed")
	}
	if rows == nil {
		c.String(http.StatusBadRequest, "Login failed")
	} else {
		rows.Next()
		var userId, nUsername, nPassword string
		rows.Scan(&userId, &nUsername, &nPassword)
		fmt.Println(userId, nUsername, nPassword, "sdfsdfsdfsdfd")
		c.SetCookie("authToken", userId, 1800, "/", "localhost", true, true)
		c.String(http.StatusAccepted, "Login successfully")
	}
}

func registerHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	// fmt.Println(username, password)
	// c.String(http.StatusBadRequest, "wrong")
	rows, _ := db.Query("SELECT * FROM userInfo WHERE username == " + "'" + username + "';")
	if rows != nil && rows.Next() {
		c.String(http.StatusBadRequest, "Username duplicated")
		fmt.Print("username duplicated", rows.Next())
	} else {
		userId := uuid.New().String()
		err := processDbStmt("INSERT INTO userInfo VALUES ('" + userId + "','" + username + "','" + password + "');")
		if err != nil {
			fmt.Println(err.Error(), "error with insert registration info")
			c.String(http.StatusBadRequest, "Cannot create account")
			return
		}
		c.String(http.StatusCreated, "Account created successfully")
	}
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
	fmt.Println("Connected to sqlite database")
	defer db.Close()

	err = processDbStmt("CREATE TABLE IF NOT EXISTS userInfo(id TEXT PRIMARY KEY NOT NULL, username TEXT NOT NULL, password TEXT NOT NULL);")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("UserInfo Table created")
	err = processDbStmt("CREATE TABLE IF NOT EXISTS taskInfo(id TEXT PRIMARY KEY NOT NULL, userId TEXT NOT NULL, taskName TEXT NOT NULL, taskStatus INT NOT NULL);")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("TaskInfo table created")

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Credentials", "Access-Control-Allow-Origin", "Cookie"}

	router.Use(cors.New(corsConfig))
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)
	router.GET("/todo-list/get-task-list", getTodoList)
	router.POST("/todo-list/remove-task", removeTaskHandler)
	router.POST("/todo-list/add-task", addTaskHandler)
	router.POST("todo-list/change-task-status", changeStatusHandler)

	router.Run("localhost:8000")
}
