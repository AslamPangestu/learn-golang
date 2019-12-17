package models

import (
	"fmt"

	"github.com/aslampangestu/learn-golang/api-example/views"
)

// InsertNote -> Insert Note
func InsertNote(title, desc string) error {
	insert, err := conn.Query("INSERT INTO notes (title, description) VALUES(?,?)", title, desc)
	defer insert.Close()
	if err != nil {
		return err
	}
	return nil
}

// SelectAllNote -> Get List All Notes
func SelectAllNote() ([]views.GetRequestNote, error) {
	rows, err := conn.Query("SELECT * FROM notes")
	if err != nil {
		return nil, err
	}
	todos := []views.GetRequestNote{}
	for rows.Next() {
		data := views.GetRequestNote{}
		rows.Scan(&data.ID, &data.Title, &data.Desc)
		fmt.Println(data)
		todos = append(todos, data)
	}
	return todos, nil
}

// SelectIDNote -> Get List All Notes
func SelectIDNote(id int) (views.GetRequestNote, error) {
	row, err := conn.Query("SELECT * FROM notes WHERE id = ?", id)
	if err != nil {
		return views.GetRequestNote{}, err
	}
	data := views.GetRequestNote{}
	row.Scan(&data.ID, &data.Title, &data.Desc)
	return data, nil
}

// SelectValueNote -> Get List All Notes
func SelectValueNote(title, desc string) ([]views.GetRequestNote, error) {
	rows, err := conn.Query("SELECT * FROM notes WHERE title = ?", title)
	if err != nil {
		return nil, err
	}
	todos := []views.GetRequestNote{}
	for rows.Next() {
		data := views.GetRequestNote{}
		rows.Scan(&data.ID, &data.Title, &data.Desc)
		fmt.Println(data)
		todos = append(todos, data)
	}
	return todos, nil
}

// DeleteNote -> Insert Note
func DeleteNote(id int) error {
	delete, err := conn.Query("DELETE FROM notes WHERE id = ?", id)
	defer delete.Close()
	if err != nil {
		return err
	}
	return nil
}
