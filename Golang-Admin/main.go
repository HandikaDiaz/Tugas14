package main

import (
	"context"
	"fmt"
	connection "gola1/conection"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Blog struct {
	ID          int
	Title       string
	Description string
	StartDate   string
	EndDate     string
	Author      string
	Duration    string
	Image       string
	Animal      bool
	Human       bool
	Demon       bool
	Robot       bool
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

// var dataBlog = []Blog{
// 	{
// 		Title:       "Hallo Title 1",
// 		Description: "Halo Content 1",
// 		Author:      "Alex",
// 		Image:       "franky.jpg",
// 	},
// 	{
// 		Title:       "Hallo Title 2",
// 		Description: "Halo Content 2",
// 		Author:      "Alexis",
// 		Image:       "nami.jpg",
// 	},
// }

func main() {
	connection.DatabaseConnection()

	e := echo.New()

	e.Static("/public", "public")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/form-blog", formAddBlog)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/blog-edit/:id", formEditBlog)

	e.GET("/register-form", registerForm)
	e.POST("/register", register)

	e.GET("/login-form", loginForm)
	e.POST("/login", login)

	e.POST("/add-blog", addBlog)
	e.POST("/blog-delete/:id", deleteBlog)
	e.POST("/blog-edit/:id", editBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, description, image, start_date, end_date, animal, human, demon, robot, duration FROM tb_blog")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Description, &each.Image, &each.StartDate, &each.EndDate, &each.Animal, &each.Human, &each.Demon, &each.Robot, &each.Duration)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Author = "Alex"

		result = append(result, each)
	}

	sess, _ := session.Get("session", c)

	blogs := map[string]interface{}{
		"Blogs":        result,
		"Flashstatus":  sess.Values["status"],
		"Flashmessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), blogs)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, description, image, start_date, end_date, animal, human, demon, robot, duration FROM tb_blog")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Description, &each.Image, &each.StartDate, &each.EndDate, &each.Animal, &each.Human, &each.Demon, &each.Robot, &each.Duration)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		result = append(result, each)
	}

	var tmpl, errtemplate = template.ParseFiles("views/blog.html")

	if errtemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errtemplate.Error()})
	}

	blogs := map[string]interface{}{
		"Blogs": result,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func formAddBlog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, description, image, start_date, end_date, animal, human, demon, robot, duration FROM tb_blog WHERE id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Title, &BlogDetail.Description, &BlogDetail.Image, &BlogDetail.StartDate, &BlogDetail.EndDate, &BlogDetail.Animal, &BlogDetail.Human, &BlogDetail.Demon, &BlogDetail.Robot, &BlogDetail.Duration)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, errtemplate = template.ParseFiles("views/blog-detail.html")

	if errtemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("input-tittle")
	description := c.FormValue("input-description")
	image := c.FormValue("input-image")
	startdate := c.FormValue("input-start-date")
	enddate := c.FormValue("input-end-date")
	duration := countDuration(startdate, enddate)

	var animal bool
	if c.FormValue("check-animal") == "yes" {
		animal = true
	}

	var human bool
	if c.FormValue("check-human") == "yes" {
		human = true
	}

	var demon bool
	if c.FormValue("check-demon") == "yes" {
		demon = true
	}

	var robot bool
	if c.FormValue("check-robot") == "yes" {
		robot = true
	}

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (title, description, start_date, end_date, image, animal, human, demon, robot, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", title, description, startdate, enddate, image, animal, human, demon, robot, duration)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("ID : ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func formEditBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var EditBlog = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, description, image, start_date, end_date, animal, human, demon, robot, duration FROM tb_blog WHERE id=$1", id).Scan(
		&EditBlog.ID, &EditBlog.Title, &EditBlog.Description, &EditBlog.Image, &EditBlog.StartDate, &EditBlog.EndDate, &EditBlog.Animal, &EditBlog.Human, &EditBlog.Demon, &EditBlog.Robot, &EditBlog.Duration)

	Blogs := map[string]interface{}{
		"Blogs": EditBlog,
	}

	var tmpl, errtemplate = template.ParseFiles("views/form-edit-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errtemplate.Error()})
	}

	return tmpl.Execute(c.Response(), Blogs)
}

func editBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("input-tittle")
	description := c.FormValue("input-description")
	image := c.FormValue("input-image")
	startdate := c.FormValue("input-start-date")
	enddate := c.FormValue("input-end-date")
	duration := countDuration(startdate, enddate)

	var animal bool
	if c.FormValue("check-animal") == "yes" {
		animal = true
	}

	var human bool
	if c.FormValue("check-human") == "yes" {
		human = true
	}

	var demon bool
	if c.FormValue("check-demon") == "yes" {
		demon = true
	}

	var robot bool
	if c.FormValue("check-robot") == "yes" {
		robot = true
	}

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_blog SET title=$1, description=$2, start_date=$3, end_date=$4, image=$5, animal=$6, human=$7, demon=$8, robot=$9, duration=$10 WHERE id=$11", title, description, startdate, enddate, image, animal, human, demon, robot, duration, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func countDuration(inputStartDate string, inputEndDate string) string {
	startDate, _ := time.Parse("2006-01-02", inputStartDate)
	endDate, _ := time.Parse("2006-01-02", inputEndDate)

	durationDate := int(endDate.Sub(startDate).Hours())
	durationDays := durationDate / 24
	durationMonths := durationDays / 30
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + "Years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + "year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + "Months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + "Month"
		} else {
			if durationDays > 1 {
				duration = strconv.Itoa(durationDays) + "Days"
			} else if durationDays > 0 {
				duration = strconv.Itoa(durationDays) + "Day"
			}
		}
	}

	return duration
}

func registerForm(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/register-form.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func register(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("input-name")
	email := c.FormValue("input-email")
	password := c.FormValue("input-pw")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register Failed, Please Try Again", false, "/register-form")
	}

	return redirectWithMessage(c, "Register Success !", true, "/login-form")
}

func loginForm(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "status")
	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/login-form.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("input-email")
	password := c.FormValue("input-pw")

	user := User{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Incorrect", false, "/login-form")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect", false, "/login-form")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 108000
	sess.Values["message"] = "Login is Success !"
	sess.Values["status"] = true
	sess.Values["id"] = user.ID
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["IsLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, path)
}
