package main

import "fmt"

// Book comment
type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title: \t\t%q\n"+
			"Author: \t\t%q\n"+
			"Published: \t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "Title 1",
		Author:        "Author 1",
		YearPublished: 1979,
	},
	{
		ID:            2,
		Title:         "Title 2",
		Author:        "Author 2",
		YearPublished: 1980,
	},
	{
		ID:            3,
		Title:         "Title 3",
		Author:        "Author 3",
		YearPublished: 1983,
	},
	{
		ID:            4,
		Title:         "Title 4",
		Author:        "Author 4",
		YearPublished: 1979,
	},
	{
		ID:            5,
		Title:         "Title 5",
		Author:        "Author 5",
		YearPublished: 1985,
	},
	{
		ID:            6,
		Title:         "Title 6",
		Author:        "Author 6",
		YearPublished: 1986,
	},
	{
		ID:            7,
		Title:         "Title 7",
		Author:        "Author 7",
		YearPublished: 1987,
	},
	{
		ID:            8,
		Title:         "Title 8",
		Author:        "Author 8",
		YearPublished: 1988,
	},
	{
		ID:            9,
		Title:         "Title 9",
		Author:        "Author 9",
		YearPublished: 1989,
	},
	{
		ID:            10,
		Title:         "Title 10",
		Author:        "Author 10",
		YearPublished: 2020,
	},
}
