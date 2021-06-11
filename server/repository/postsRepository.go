package repository

import (
	"example/driver"
	"example/models"
	"fmt"
	"time"
	"unicode/utf8"
)

type ResponseModel struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}
//Menampilkan semua artikel
func ReadAllPosts() []models.PostsModel {
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.PostsModel

	items, err := db.Query("select id, title, content, category, status from posts")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.PostsModel{}
		var err = items.Scan(&each.Id, &each.Title, &each.Content, &each.Category, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}
//Membuat article baru
func CreatePosts(U *models.PostsModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	date := time.Now()

	lenTitle := utf8.RuneCountInString(U.Title)
	lenContent := utf8.RuneCountInString(U.Content)
	lenCategory := utf8.RuneCountInString(U.Category)

	if lenTitle >= 20 && lenContent >= 200 && lenCategory >= 3 && (U.Status == "Draft" || U.Status == "Publish" || U.Status == "Trash") {

	_, err = db.Exec(`INSERT INTO posts (title, content, category, status, created_date) VALUES (?, ?, ?, ?, ?)`, U.Title, U.Content, U.Category, U.Status, date)

	fmt.Println("insert success!")
	Res = &ResponseModel{200, "Success save Data"}

	} else {
		fmt.Println("does not meet requirements!")
		Res = &ResponseModel{400, "does not meet requirements"}
	}

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	
	return Res
}
//Menampilkan article dengan id yang di-request
func ReadPostsById(Id int) []models.PostsModel {
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.PostsModel

	items, err := db.Query("select title, content, category, status from posts where id=?", Id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.PostsModel{}
		var err = items.Scan(&each.Title, &each.Content, &each.Category, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}
//Merubah data article dengan id yang di-request
func UpdatePosts(U *models.PostsModel, Id int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	date := time.Now()

	lenTitle := utf8.RuneCountInString(U.Title)
	lenContent := utf8.RuneCountInString(U.Content)
	lenCategory := utf8.RuneCountInString(U.Category)

	if lenTitle >= 20 && lenContent >= 200 && lenCategory >= 3 && (U.Status == "Draft" || U.Status == "Publish" || U.Status == "Trash") {

	_, err = db.Exec("update posts set title = ?, content = ?, category = ?, status = ?, updated_date = ? where id = ?", U.Title, U.Content, U.Category, U.Status, date, Id)
	
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}

	} else {
		fmt.Println("does not meet requirements!")
		Res = &ResponseModel{400, "does not meet requirements"}
	}
	
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	
	return Res
}
//Menghapus data article dengan id yang di request
func DeletePosts(Id int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("delete from posts where id = ?", Id)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Delete success!")
	Res = &ResponseModel{200, "Success delete Data"}
	return Res
}
//Menampilkan seluruh article di database dengan paging pada parameter limit & offset.
func ReadPostsLimit(Offset int, Limit int) []models.PostsModel {
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.PostsModel

	items, err := db.Query("select title, content, category, status from posts Limit ?, ?", Offset, Limit)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.PostsModel{}
		var err = items.Scan(&each.Title, &each.Content, &each.Category, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

//API untuk kebutuhan Frontend
// Menampilkan artikel dengan status publish (untuk data di halaman all post (tab publish))
func ReadPostsPublish() []models.PostsModel {
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.PostsModel

	items, err := db.Query("select id, title, content, category, status from posts where status='publish'")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.PostsModel{}
		var err = items.Scan(&each.Id, &each.Title, &each.Content, &each.Category, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

// Menampilkan artikel dengan status draft (untuk data di halaman all post (tab draft))
func ReadPostsDraft() []models.PostsModel {
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.PostsModel

	items, err := db.Query("select id, title, content, category, status from posts where status='Draft'")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.PostsModel{}
		var err = items.Scan(&each.Id, &each.Title, &each.Content, &each.Category, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

// Menampilkan artikel dengan status trash (untuk data di halaman all post (tab trash))
func ReadPostsTrash() []models.PostsModel {
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.PostsModel

	items, err := db.Query("select id, title, content, category, status from posts where status='Trash'")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.PostsModel{}
		var err = items.Scan(&each.Id, &each.Title, &each.Content, &each.Category, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

// Update status menjadi Trash
func UpdateStatusToTrash(Id int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	date := time.Now()

	_, err = db.Exec("update posts set status = 'Trash', updated_date = ? where id = ?", date, Id)
	
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}

	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}

	return Res
}
