package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// create handler func
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		// return status error if input onther type
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//loop info to show  , range is show info all , "_," is index , book is database
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Not found eiei") // it how to return id show

}

func createBook(c *fiber.Ctx) error {
	book := new(Book)

	// send err first
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// handler err if can not fine book pointer
	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Auther = bookUpdate.Auther
			return c.JSON(books[i])
			// force thow err
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("not found")
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books = append(books[:i], books[1+i:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Not found")
}
