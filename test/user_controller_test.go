package test

import (
	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	mockDB = map[string]*models.User{
// 		"jon@labstack.com": &models.User{Name: "Jon Snow", Age: 30},
// 	}
// 	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
// )

// func TestGetusers(t *testing.T) {
// 	e := router.NewRouter()
// 	req := httptest.NewRequest(http.MethodGet, "/users", nil)
// 	rec := httptest.NewRecorder()

// 	e.ServeHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.JSONEq(t, `{"name": "Jiro", "email": "jiro@example.com"}`, rec.Body.String())
// }

// var db *sql.DB

// func TestMain(m *testing.M) {
// 	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}

// 	// pulls an image, creates a container based on it and runs it
// 	resource, err := pool.Run("mysql", "8.0", []string{"MYSQL_ROOT_PASSWORD=root"})
// 	if err != nil {
// 		log.Fatalf("Could not start resource: %s", err)
// 	}

// 	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
// 	if err := pool.Retry(func() error {
// 		var err error
// 		db, err = sql.Open("mysql", fmt.Sprintf("root:root@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
// 		if err != nil {
// 			return err
// 		}
// 		return db.Ping()
// 	}); err != nil {
// 		log.Fatalf("Could not connect to database: %s", err)
// 	}

// 	code := m.Run()

// 	// You can't defer this because os.Exit doesn't care for defer
// 	if err := pool.Purge(resource); err != nil {
// 		log.Fatalf("Could not purge resource: %s", err)
// 	}

// 	os.Exit(code)

// 	// e := router.NewRouter()
// 	// req := httptest.NewRequest(http.MethodGet, "/users", nil)
// 	// rec := httptest.NewRecorder()

// 	// e.ServeHTTP(rec, req)

// 	// assert.Equal(t, http.StatusOK, rec.Code)
// 	// assert.JSONEq(t, `{"name": "Jiro", "email": "jiro@example.com"}`, rec.Body.String())
// }
