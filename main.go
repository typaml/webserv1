package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"net/http"
)

type Book struct {
	ID                int    `json:"ID"`
	Author            string `json:"Author"`
	Title             string `json:"Title"`
	Genres            string `json:"Genres"`
	Description       string `json:"Description"`
	YearOfPublication int    `json:"Year_of_Publication"`
	Language          string `json:"Language"`
	IsPopular         string `json:"Is_Popular"`
	IsBestseller      string `json:"Is_Bestseller"`
	Pricebook         string `json:"Pricebook"`
	// Добавьте другие поля, если необходимо
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Favorite struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}

var (
	db, err = sql.Open("mysql", "root:1234s@/library")
	store   = sessions.NewCookieStore([]byte("your-secret-key"))
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/proverka.html")
	})

	http.HandleFunc("/static/audiopage.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/audiopage.html")
	})

	http.HandleFunc("/static/proverka.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/proverka.html")
	})
	http.HandleFunc("/static/username.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/username.html")
	})
	http.HandleFunc("/static/profile.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/profile.html")
	})
	http.HandleFunc("/removeFromFavorites", removeFromFavorites)
	http.HandleFunc("/addToFavorites", addToFavoritesHandler)
	http.HandleFunc("/getFavorites", getFavorites)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/login", handleLogin)
	// Добавьте обработчик для эндпоинта /logout
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/searchBooks", searchBooks)
	http.HandleFunc("/widget1func", widget1)
	http.HandleFunc("/widget2func", widget2)
	http.HandleFunc("/widget3func", widget3)
	http.HandleFunc("/widget4func", widget4)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}

func searchBooks(w http.ResponseWriter, r *http.Request) {
	// Получение параметра поиска из запроса
	query := r.URL.Query().Get("query")

	// Выполнение SQL-запроса с использованием LIKE для поиска по автору или названию
	rows, err := db.Query("SELECT ID, Author, Title, Genres, Description, Year_of_Publication, Language, CASE WHEN Is_Bestseller = '1' THEN 'Да' ELSE 'Нет' END AS Is_Bestseller, pricebook FROM books WHERE Author LIKE ? OR Title LIKE ?", query, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Итерация по результатам запроса и вывод результатов
	// (вы можете использовать json.Marshal для более структурированного вывода)
	var booksData []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Genres, &book.Description, &book.YearOfPublication, &book.Language, &book.IsBestseller, &book.Pricebook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksData = append(booksData, book)
	}

	// Преобразовываем данные в формат JSON и отправляем клиенту
	jsonData, err := json.Marshal(booksData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func widget1(w http.ResponseWriter, r *http.Request) {
	// Выполнение SQL-запроса с использованием LIKE для поиска по автору или названию
	rows, err := db.Query("SELECT ID, Author, Title, Genres, Description, Year_of_Publication, Language, CASE WHEN Is_Bestseller = '1' THEN 'Да' ELSE 'Нет' END AS Is_Bestseller, pricebook FROM books WHERE Is_Bestseller LIKE ?", true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Итерация по результатам запроса и вывод результатов
	// (вы можете использовать json.Marshal для более структурированного вывода)
	var booksData []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Genres, &book.Description, &book.YearOfPublication, &book.Language, &book.IsBestseller, &book.Pricebook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksData = append(booksData, book)
	}

	// Преобразовываем данные в формат JSON и отправляем клиенту
	jsonData, err := json.Marshal(booksData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func widget2(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT ID, Author, Title, Genres, Description, Year_of_Publication, Language, CASE WHEN Is_Bestseller = '1' THEN 'Да' ELSE 'Нет' END AS Is_Bestseller, pricebook FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var booksData []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Genres, &book.Description, &book.YearOfPublication, &book.Language, &book.IsBestseller, &book.Pricebook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksData = append(booksData, book)
	}

	// Преобразовываем данные в формат JSON и отправляем клиенту
	jsonData, err := json.Marshal(booksData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func widget3(w http.ResponseWriter, r *http.Request) {
	// Выполнение SQL-запроса с использованием LIKE для поиска по автору или названию
	rows, err := db.Query("SELECT ID, Author, Title, Genres, Description, Year_of_Publication, Language, CASE WHEN Is_Bestseller = '1' THEN 'Да' ELSE 'Нет' END AS Is_Bestseller, pricebook FROM books WHERE Is_Popular LIKE ?", true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Итерация по результатам запроса и вывод результатов
	// (вы можете использовать json.Marshal для более структурированного вывода)
	var booksData []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Genres, &book.Description, &book.YearOfPublication, &book.Language, &book.IsBestseller, &book.Pricebook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksData = append(booksData, book)
	}

	// Преобразовываем данные в формат JSON и отправляем клиенту
	jsonData, err := json.Marshal(booksData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func widget4(w http.ResponseWriter, r *http.Request) {
	// Выполнение SQL-запроса с использованием LIKE для поиска по автору или названию
	rows, err := db.Query("SELECT ID, Author, Title, Genres, Description, Year_of_Publication, Language, CASE WHEN Is_Bestseller = '1' THEN 'Да' ELSE 'Нет' END AS Is_Bestseller, pricebook FROM books WHERE (Is_Popular = '1' or Is_Bestseller = '1') and (pricebook = 'Бесплатно')")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Итерация по результатам запроса и вывод результатов
	// (вы можете использовать json.Marshal для более структурированного вывода)
	var booksData []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Genres, &book.Description, &book.YearOfPublication, &book.Language, &book.IsBestseller, &book.Pricebook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksData = append(booksData, book)
	}

	// Преобразовываем данные в формат JSON и отправляем клиенту
	jsonData, err := json.Marshal(booksData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Обработчик регистрации пользователя
func handleRegister(w http.ResponseWriter, r *http.Request) {
	// Получение данных из формы регистрации
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form data:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	username := r.Form.Get("registerUsername")
	password := r.Form.Get("registerPassword")

	// Проверка, что пользователь с таким именем не существует
	exists, err := userExists(username)
	if err != nil {
		fmt.Println("Error checking if user exists:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if exists {
		fmt.Println("Username already exists:", username)
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	// Добавление пользователя в базу данных
	err = addUser(username, password)
	if err != nil {
		fmt.Println("Error adding user to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User{Username: username, Password: password})
}

// Обработчик входа пользователя
func handleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form data:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	username := r.Form.Get("loginUsername")
	password := r.Form.Get("loginPassword")

	// Проверка, что пользователь существует и пароль совпадает
	valid, err := validateUser(username, password)
	if err != nil {
		fmt.Println("Error validating user:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !valid {
		fmt.Println("Invalid username or password:", username)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Создание сессии
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Установка состояния аутентификации в сессии
	session.Values["authenticated"] = true
	session.Values["username"] = username
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User{Username: username, Password: password})
}

// Проверка существования пользователя в базе данных
func userExists(username string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Добавление пользователя в базу данных
func addUser(username, password string) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		return err
	}
	return nil
}

// Проверка правильности логина и пароля пользователя
func validateUser(username, password string) (bool, error) {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		return false, err
	}
	return storedPassword == password, nil
}

// Обработчик выхода пользователя
func handleLogout(w http.ResponseWriter, r *http.Request) {
	// Получение сессии
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Очистка состояния аутентификации в сессии
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Перенаправление на главную страницу
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getFavorites(w http.ResponseWriter, r *http.Request) {
	// Получение текущего пользователя из сессии
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username, ok := session.Values["username"].(string)
	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	// Выполнение SQL-запроса для получения избранных книг пользователя
	rows, err := db.Query("SELECT b.ID, b.Author, b.Title, b.Genres, b.Description, b.Year_of_Publication, b.Language, "+
		"CASE WHEN b.Is_Bestseller = '1' THEN 'Да' ELSE 'Нет' END AS Is_Bestseller, b.pricebook "+
		"FROM books b "+
		"JOIN favorites f ON b.ID = f.book_id "+
		"JOIN users u ON f.user_id = u.ID "+
		"WHERE u.username = ?", username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Итерация по результатам запроса и вывод результатов
	var booksData []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Genres, &book.Description, &book.YearOfPublication, &book.Language, &book.IsBestseller, &book.Pricebook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksData = append(booksData, book)
	}

	// Преобразовываем данные в формат JSON и отправляем клиенту
	jsonData, err := json.Marshal(booksData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func addToFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	// Распаковка данных из запроса
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Получение информации о книге и пользователях из данных запроса
	bookID := int(data["bookID"].(float64))

	// Получение текущего пользователя из сессии
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username, ok := session.Values["username"].(string)
	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	// Получение ID пользователя из базы данных
	var userID int
	err = db.QueryRow("SELECT ID FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Добавление записи в избранное
	_, err = db.Exec("INSERT INTO favorites (user_id, book_id) VALUES (?, ?)", userID, bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}
func removeFromFavorites(w http.ResponseWriter, r *http.Request) {
	// Получение текущего пользователя из сессии
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username, ok := session.Values["username"].(string)
	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	// Получение ID книги из параметра запроса
	bookID := r.URL.Query().Get("bookId")
	if bookID == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}

	// Выполнение SQL-запроса для удаления книги из избранного
	_, err = db.Exec("DELETE FROM favorites WHERE user_id = (SELECT ID FROM users WHERE username = ?) AND book_id = ?", username, bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный статус
	w.WriteHeader(http.StatusOK)
}
