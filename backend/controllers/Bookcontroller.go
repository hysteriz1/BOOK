package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m16_z/app/ent"
	"github.com/m16_z/app/ent/book"
	"github.com/m16_z/app/ent/bookstatus"
	"github.com/m16_z/app/ent/room"
	"github.com/m16_z/app/ent/user"
)

// BookController  defines the struct for the book controller
type BookController struct {
	client *ent.Client
	router gin.IRouter
}
type Book struct {
	RESERVATIONS string
	User       	 int
	Room       	 int
	BookStatus 	 int
}

// CreateBook handles POST requests for adding book entities
// @Summary Create book
// @Description Create book
// @ID create-book
// @Accept   json
// @Produce  json
// @Param book body Book true "Book entity"
// @Success 200 {object} ent.Book
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /book [post]
func (ctl *BookController) CreateBook(c *gin.Context) {
	obj := Book{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "book binding failed",
		})
		return
	}

	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	bst, err := ctl.client.Bookstatus.
		Query().
		Where(bookstatus.IDEQ(int(obj.BookStatus))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bookstatus not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.RESERVATIONS)

	b, err := ctl.client.Book.
		Create().
		SetRESERVATIONS(times).
		SetBOOKUSER(u).		
		SetBOOKROOM(r).
		SetBOOKBOOKSTATUS(bst).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// GetBook handles GET requests to retrieve a book entity
// @Summary Get a book entity by ID
// @Description get book by ID
// @ID get-book
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} ent.Book
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /book/{id} [get]
func (ctl *BookController) GetBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := ctl.client.Book.
		Query().
		Where(book.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// ListBook handles request to get a list of book entities
// @Summary List book entities
// @Description list book entities
// @ID list-book
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Book
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /book [get]
func (ctl *BookController) ListBook(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	Book, err := ctl.client.Book.
		Query().
		WithBOOKUSER().
		WithBOOKBOOKSTATUS().
		WithBOOKROOM().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Book)
}

// DeleteBook handles DELETE requests to delete a book entity
// @Summary Delete a book entity by ID
// @Description get book by ID
// @ID delete-book
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /book/{id} [delete]
func (ctl *BookController) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Book.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateBook handles PUT requests to update a book entity
// @Summary Update a book entity by ID
// @Description update book by ID
// @ID update-book
// @Accept   json
// @Produce  json
// @Param id path int true "Book ID"
// @Param Book body ent.Book true "Book entity"
// @Success 200 {object} ent.Book
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /book/{id} [put]
func (ctl *BookController) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Book{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "book binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.Book.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, b)
}

// NewBookController creates and registers handles for the book controller
func NewBookController(router gin.IRouter, client *ent.Client) *BookController {
	uc := &BookController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitBookController registers routes to the main engine
func (ctl *BookController) register() {
	users := ctl.router.Group("/book")

	users.GET("", ctl.ListBook)

	// CRUD
	users.POST("", ctl.CreateBook)
	users.GET(":id", ctl.GetBook)
	users.PUT(":id", ctl.UpdateBook)
	users.DELETE(":id", ctl.DeleteBook)
}
