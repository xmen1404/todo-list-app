package main

import (
	"fmt"
	"net/http"

	"database/sql"

	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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
var CLIENT_DOMAIN string = "http://localhost"
var CLIENT_PORT string = "3000"
var CLIENT_URL string = CLIENT_DOMAIN + ":" + CLIENT_PORT

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

	row, err := db.Query("SELECT * FROM taskinfo WHERE userId = '" + userId + "';")
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

func corsAcessMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", CLIENT_URL)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Next()
}

func authMiddleware(c *gin.Context) {
	authToken, errAuth := c.Cookie("authToken")
	fmt.Println(authToken)
	if errAuth != nil {
		fmt.Println(errAuth.Error(), "cookie error")
		c.AbortWithStatus(http.StatusForbidden)
	}

	rows, err := db.Query("SELECT * FROM userinfo WHERE id = '" + authToken + "';")
	if err != nil {
		fmt.Println(err.Error(), "error in query userInfo")
		c.AbortWithStatus(http.StatusForbidden)
	}
	defer rows.Close()

	c.Next()
}

func getTodoList(c *gin.Context) {
	authToken, _ := c.Cookie("authToken") // error checked in authMiddleware
	ndata, err := loadTaskList(authToken)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, todoListStruct{})
		return
	}
	c.IndentedJSON(http.StatusOK, todoListStruct{ndata})
}

func processDbStmt(statement string) error {
	fmt.Println(statement)
	stmt, err := db.Prepare(statement)
	if err == nil {
		stmt.Exec()
	}
	defer stmt.Close()
	return err
}

func addTaskHandler(c *gin.Context) {
	fmt.Println("passed authProcess")
	authToken, _ := c.Cookie("authToken") // error checked in authMiddleware
	nTaskName := c.PostForm("taskname")
	nTaskId := uuid.New().String()
	nTaskStatus := "0"

	err := processDbStmt("INSERT INTO taskinfo VALUES ('" + nTaskId + "','" + authToken + "','" + nTaskName + "','" + nTaskStatus + "');")
	if err != nil {
		fmt.Println(err.Error(), "error with inserting to db")
		c.String(http.StatusBadRequest, "Cannot add task")
		return
	}

	c.String(http.StatusCreated, "Added task successfully")
}

func removeTaskHandler(c *gin.Context) {
	authToken, _ := c.Cookie("authToken")
	nTaskID := c.PostForm("taskid")
	err := processDbStmt("DELETE FROM taskinfo WHERE id = '" + nTaskID + "' AND userid = '" + authToken + "'")
	if err != nil {
		fmt.Println(err.Error(), "error with deleting from db")
		c.String(http.StatusBadRequest, "Cannot delete task")
		return
	}

	c.String(http.StatusCreated, "Delete task successfully")
}

func changeStatusHandler(c *gin.Context) {
	authToken, _ := c.Cookie("authToken")
	nTaskID := c.PostForm("taskid")

	err := processDbStmt("UPDATE taskinfo SET taskstatus = 1 - taskstatus " + "WHERE id = '" + nTaskID + "' AND userid = '" + authToken + "';")
	if err != nil {
		fmt.Println(err.Error(), "error with updating db")
		c.String(http.StatusBadRequest, "Cannot change taskStatus")
		return
	}
	c.String(http.StatusCreated, "Change task status successfully")
}

func loginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	rows, err := db.Query("SELECT * FROM userinfo WHERE username = " + "'" + username + "' AND password = '" + password + "';")
	if err != nil || rows == nil {
		fmt.Println(err.Error(), "error with Select command")
		c.String(http.StatusBadRequest, "Login failed")
	}
	if !rows.Next() {
		c.String(http.StatusBadRequest, "Login failed")
	} else {
		var userId, nUsername, nPassword string
		rows.Scan(&userId, &nUsername, &nPassword)
		fmt.Println(userId, nUsername, nPassword)
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("authToken", userId, 1800, "/", CLIENT_DOMAIN, true, true)
		c.String(http.StatusAccepted, "Login successfully")
	}
	defer rows.Close()
}

func registerHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	rows, _ := db.Query("SELECT * FROM userinfo WHERE username = " + "'" + username + "';")
	if rows != nil && rows.Next() {
		c.String(http.StatusBadRequest, "Username duplicated")
		fmt.Print("username duplicated", rows.Next())
	} else {
		userId := uuid.New().String()
		err := processDbStmt("INSERT INTO userinfo VALUES ('" + userId + "','" + username + "','" + password + "');")
		if err != nil {
			fmt.Println(err.Error(), "error with insert registration info")
			c.String(http.StatusBadRequest, "Cannot create account")
			return
		}
		c.String(http.StatusCreated, "Account created successfully")
	}
	defer rows.Close()
}

func main() {
	var err error = nil
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err = sql.Open("mysql", "root:root@/todo_list_app")
	if err != nil {
		fmt.Println("Error with login to db:", err.Error())
		return
	}
	fmt.Println("Connected to mysql database")

	defer db.Close()

	err = processDbStmt("CREATE TABLE IF NOT EXISTS userinfo(id VARCHAR(100) PRIMARY KEY NOT NULL, username VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL);")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("userinfo Table created")
	err = processDbStmt("CREATE TABLE IF NOT EXISTS taskinfo(id VARCHAR(100) PRIMARY KEY NOT NULL, userid VARCHAR(100) NOT NULL, taskname VARCHAR(100) NOT NULL, taskstatus INT NOT NULL);")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("taskinfo table created")

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{CLIENT_URL}
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Credentials", "Access-Control-Allow-Origin"}

	router.Use(cors.New(corsConfig))
	// router.Use(authMiddleware)
	router.Use(corsAcessMiddleware)
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)

	todoListRoutes := router.Group("/todo-list")
	todoListRoutes.Use(authMiddleware)
	{
		todoListRoutes.GET("/get-task-list", getTodoList)
		todoListRoutes.POST("/remove-task", removeTaskHandler)
		todoListRoutes.POST("/add-task", addTaskHandler)
		todoListRoutes.POST("/change-task-status", changeStatusHandler)
	}

	router.Run(":8000")
}
