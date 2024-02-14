package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/russross/blackfriday"
	"golang.org/x/time/rate"

	"bufio"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var db *sql.DB
var limiter = rate.NewLimiter(100, 1)
var (
	adminLogin     string
	adminPassword  string
	dbName         string
	dbUser         string
	dbPassword     string
	createTableSQL string
	insertDataSQL  string
)

type Article struct {
	ID      int64
	Title   string
	Content string
}

type Pagination struct {
	Page        int  `json:"page"`
	PageSize    int  `json:"page_size"`
	TotalCount  int  `json:"total_count"`
	TotalPages  int  `json:"total_pages"`
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
}

func main() {
	err := readAdminCredentials()
	if err != nil {
		log.Fatal(err)
	}
	connectToDatabase()
	routerProcess()
}

func routerProcess() {
	router := gin.Default()

	admin := router.Group("/admin")
	admin.Use(AuthRequired())

	store := cookie.NewStore([]byte("secret3"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(rateLimitMiddleware())

	router.SetFuncMap(template.FuncMap{
		"sub":      sub,
		"add":      add,
		"markDown": markDowner,
	})

	router.StaticFS("/images", http.Dir("./images"))
	router.StaticFS("/css", http.Dir("./css"))
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", showIndexPage)
	router.GET("/article/:id", getArticle)

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", loginRoute)

	router.GET("/admin", getAdminNew)
	router.POST("/admin/new", postAdminNew)

	router.Run(":8888")
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.Next()
	}
}

func showIndexPage(c *gin.Context) {
	page := 1
	pageSize := 3

	if pageParam, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		page = pageParam
	}

	articles, err := getArticles(page, pageSize)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	totalCount, err := getArticleCount()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	pagination := NewPagination(page, pageSize, totalCount)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "My amazing blog",
		"articles":   articles,
		"pagination": pagination,
	})
}

func getArticle(c *gin.Context) {
	id := c.Param("id")

	article, err := getArticleByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title":   article.Title,
			"article": Article{ID: article.ID, Title: article.Title, Content: article.Content},
		},
	)
}

func postAdminNew(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	postArticle(c)

	c.Redirect(http.StatusFound, "/admin")
}

func postArticle(c *gin.Context) {
	type newArticle struct {
		Title   string `form:"title" binding:"required"`
		Content string `form:"content" binding:"required"`
	}

	var input newArticle
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := Article{
		Title:   input.Title,
		Content: input.Content,
	}
	_, err := db.Exec("INSERT INTO articles (title, content) VALUES ($1, $2)", article.Title, article.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}
}

func NewPagination(page int, pageSize int, totalCount int) *Pagination {
	pages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	hasNextPage := page < pages
	hasPrevPage := page > 1

	return &Pagination{
		Page:        page,
		PageSize:    pageSize,
		TotalCount:  totalCount,
		TotalPages:  pages,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
	}
}

func sub(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func loginRoute(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == adminLogin && password == adminPassword {
		session := sessions.Default(c)
		session.Set("user", username)
		session.Save()

		c.Redirect(http.StatusFound, "/admin")
		return
	}

	c.HTML(http.StatusBadRequest, "login.html", gin.H{
		"error": "Invalid username or password",
	})
}

func getAdminNew(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "admin.html", nil)
}

func connectToDatabase() {
	connStr := "user=" + dbUser +
		" password=" + dbPassword +
		" host=127.0.0.1" +
		" port=5432" +
		" dbname=" + dbName +
		" sslmode=disable" +
		" TimeZone=UTC" +
		" client_encoding=UTF8"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	}

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM articles").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		_, err = db.Exec(insertDataSQL)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getArticleCount() (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM articles").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func getArticleByID(id string) (Article, error) {
	var art Article
	row := db.QueryRow("SELECT * FROM articles WHERE id = $1", id)
	if err := row.Scan(&art.ID, &art.Title, &art.Content); err != nil {
		if err == sql.ErrNoRows {
			return art, fmt.Errorf("getArticleByID %s: no such article", id)
		}
		return art, fmt.Errorf("getArticleByID: %v", err)
	}
	return art, nil
}

func getArticles(page int, pageSize int) ([]Article, error) {
	var articles []Article

	offset := (page - 1) * pageSize

	rows, err := db.Query("SELECT * FROM articles ORDER BY id DESC LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("getArticles: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var art Article
		if err := rows.Scan(&art.ID, &art.Title, &art.Content); err != nil {
			return nil, fmt.Errorf("getArticles: %v", err)
		}
		articles = append(articles, art)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getArticles: %v", err)
	}

	return articles, nil
}

func markDowner(args ...interface{}) template.HTML {
	s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	return template.HTML(s)
}

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}

func readAdminCredentials() error {
	file, err := os.Open("admin_credentials.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		switch key {
		case "DBUSER":
			dbUser = value
		case "DBNAME":
			dbName = value
		case "DBPASSWORD":
			dbPassword = value
		case "LOGIN":
			adminLogin = value
		case "PASSWORD":
			adminPassword = value
		case "CREATE_TABLE_SQL":
			createTableSQL = value
		case "INSERT_DATA_SQL":
			insertDataSQL = value
		}
	}
	return scanner.Err()
}
