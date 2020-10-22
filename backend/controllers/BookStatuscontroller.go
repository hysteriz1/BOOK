package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/m16_z/app/ent"
   "github.com/m16_z/app/ent/bookstatus"
   "github.com/gin-gonic/gin"
)
// BookstatusController  defines the struct for the bookstatus controller
type BookstatusController struct {
    client *ent.Client
    router gin.IRouter
 }
 
 // CreateBookstatus handles POST requests for adding bookstatus entities
 // @Summary Create bookstatus
 // @Description Create bookstatus
 // @ID create-bookstatus
 // @Accept   json
 // @Produce  json
 // @Param bookstatus body ent.Bookstatus true "Bookstatus entity"
 // @Success 200 {object} ent.Bookstatus
 // @Failure 400 {object} gin.H
 // @Failure 500 {object} gin.H
 // @Router /bookstatus [post]
 func (ctl *BookstatusController) CreateBookstatus(c *gin.Context) {
     obj := ent.Bookstatus{}
     if err := c.ShouldBind(&obj); err != nil {
         c.JSON(400, gin.H{
             "error": "Bookstatus binding failed",
         })
         return
     }
   
     bst, err := ctl.client.Bookstatus.
     Create().
     SetBOOKSTATUSDATA(obj.BOOKSTATUSDATA).
     Save(context.Background())
 
 
     if err != nil {
         c.JSON(400, gin.H{
             "error": "saving failed",
         })
         return
     }
   
     c.JSON(200, bst)
  }
 
 
 // GetBookstatus handles GET requests to retrieve a bookstatus entity
 // @Summary Get a bookstatus entity by ID
 // @Description get bookstatus by ID
 // @ID get-bookstatus
 // @Produce  json
 // @Param id path int true "Bookstatus ID"
 // @Success 200 {object} ent.Bookstatus
 // @Failure 400 {object} gin.H
 // @Failure 404 {object} gin.H
 // @Failure 500 {object} gin.H
 // @Router /bookstatus/{id} [get]
 func (ctl *BookstatusController) GetBookstatus(c *gin.Context) {
     id, err := strconv.ParseInt(c.Param("id"), 10, 64)
     if err != nil {
         c.JSON(400, gin.H{
             "error": err.Error(),
         })
         return
     }
   
    bst, err := ctl.client.Bookstatus.
         Query().
         Where(bookstatus.IDEQ(int(id))).
         Only(context.Background())
     if err != nil {
         c.JSON(404, gin.H{
             "error": err.Error(),
         })
         return
     }
   
     c.JSON(200, bst)
  }
 
 // ListBookstatus handles request to get a list of bookstatus entities
 // @Summary List bookstatus entities
 // @Description list bookstatus entities
 // @ID list-bookstatus
 // @Produce json
 // @Param limit  query int false "Limit"
 // @Param offset query int false "Offset"
 // @Success 200 {array} ent.Bookstatus
 // @Failure 400 {object} gin.H
 // @Failure 500 {object} gin.H
 // @Router /bookstatus [get]
 func (ctl *BookstatusController) ListBookstatus(c *gin.Context) {
     limitQuery := c.Query("limit")
     limit := 10
     if limitQuery != "" {
         limit64, err := strconv.ParseInt(limitQuery, 10, 64)
         if err == nil {limit = int(limit64)}
     }
   
     offsetQuery := c.Query("offset")
     offset := 0
     if offsetQuery != "" {
         offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
         if err == nil {offset = int(offset64)}
     }
   
     Bookstatus, err := ctl.client.Bookstatus.
         Query().
         Limit(limit).
         Offset(offset).
         All(context.Background())
         if err != nil {
         c.JSON(400, gin.H{"error": err.Error(),})
         return
     }
   
     c.JSON(200, Bookstatus)
  }
 
 
 // DeleteBookstatus handles DELETE requests to delete a bookstatus entity
 // @Summary Delete a bookstatus entity by ID
 // @Description get bookstatus by ID
 // @ID delete-bookstatus
 // @Produce  json
 // @Param id path int true "Bookstatus ID"
 // @Success 200 {object} gin.H
 // @Failure 400 {object} gin.H
 // @Failure 404 {object} gin.H
 // @Failure 500 {object} gin.H
 // @Router /bookstatus/{id} [delete]
 func (ctl *BookstatusController) DeleteBookstatus(c *gin.Context) {
     id, err := strconv.ParseInt(c.Param("id"), 10, 64)
     if err != nil {
         c.JSON(400, gin.H{
             "error": err.Error(),
         })
         return
     }
   
     err = ctl.client.Bookstatus.
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
 
 
   
 // UpdateBookstatus handles PUT requests to update a bookstatus entity
 // @Summary Update a bookstatus entity by ID
 // @Description update bookstatus by ID
 // @ID update-bookstatus
 // @Accept   json
 // @Produce  json
 // @Param id path int true "Bookstatus ID"
 // @Param Bookstatus body ent.Bookstatus true "Bookstatus entity"
 // @Success 200 {object} ent.Bookstatus
 // @Failure 400 {object} gin.H
 // @Failure 500 {object} gin.H
 // @Router /bookstatus/{id} [put]
 func (ctl *BookstatusController) UpdateBookstatus(c *gin.Context) {
     id, err := strconv.ParseInt(c.Param("id"), 10, 64)
     if err != nil {
         c.JSON(400, gin.H{
             "error": err.Error(),
         })
         return
     }
   
     obj := ent.Bookstatus{}
     if err := c.ShouldBind(&obj); err != nil {
         c.JSON(400, gin.H{
             "error": "bookstatus binding failed",
         })
         return
     }
     obj.ID = int(id)
     bst, err := ctl.client.Bookstatus.
         UpdateOne(&obj).
         Save(context.Background())
     if err != nil {
         c.JSON(400, gin.H{"error": "update failed",})
         return
     }
   
     c.JSON(200,bst)
  }
  
 // NewBookstatusController creates and registers handles for the bookstatus controller
 func NewBookstatusController(router gin.IRouter, client *ent.Client) *BookstatusController {
     uc := &BookstatusController{
         client: client,
         router: router,
     }
     uc.register()
     return uc
  }
   
  // InitBookstatusController registers routes to the main engine
  func (ctl *BookstatusController) register() {
     users := ctl.router.Group("/bookstatus")
   
     users.GET("", ctl.ListBookstatus)
   
     // CRUD
     users.POST("", ctl.CreateBookstatus)
     users.GET(":id", ctl.GetBookstatus)
     users.PUT(":id", ctl.UpdateBookstatus)
     users.DELETE(":id", ctl.DeleteBookstatus)
  }