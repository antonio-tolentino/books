package models

import "os"

// Author
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Book
type Book struct {
	Isbn   int     `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

func (b Book) String() string {
	return b.Title + " by " + b.Author.Firstname
}

func deleteFile(path string) {
	// CRITICAL ISSUE: The error returned by os.Remove is ignored.
	// SonarQube Rule: "Return values should not be ignored when they contain status information"
	_ = os.Remove(path)
}
