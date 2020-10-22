package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/m16_z/app/ent"
   "github.com/m16_z/app/ent/roomstatus"
   "github.com/gin-gonic/gin"
)
 
// RoomStatusController  defines the struct for the roomstatus controller
type RoomStatusController  struct {
   client *ent.Client
   router gin.IRouter
}
// CreateRoomStatus handles POST requests for adding roomstatus entities
// @Summary Create roomstatus
// @Description Create roomstatus
// @ID create-roomstatus
// @Accept   json
// @Produce  json
// @Param roomstatus body ent.RoomStatus true "RoomStatus entity"
// @Success 200 {object} ent.RoomStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstatus [post]
func (ctl *RoomStatusController) CreateRoomStatus(c *gin.Context) {
    obj := ent.RoomStatus{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "RoomStatus binding failed",
        })
        return
    }
  
    rst, err := ctl.client.RoomStatus.
    Create().
    SetSTATUSDATA(obj.STATUSDATA).
    Save(context.Background())


    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }
  
    c.JSON(200, rst)
 }


// GetRoomStatus handles GET requests to retrieve a roomstatus entity
// @Summary Get a roomstatus entity by ID
// @Description get roomstatus by ID
// @ID get-roomstatus
// @Produce  json
// @Param id path int true "RoomStatus ID"
// @Success 200 {object} ent.RoomStatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstatus/{id} [get]
func (ctl *RoomStatusController) GetRoomStatus(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    rst, err := ctl.client.RoomStatus.
        Query().
        Where(roomstatus.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    c.JSON(200, rst)
 }

// ListRoomStatus handles request to get a list of roomstatus entities
// @Summary List roomstatus entities
// @Description list roomstatus entities
// @ID list-roomstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RoomStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstatus [get]
func (ctl *RoomStatusController) ListRoomStatus(c *gin.Context) {
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
  
    RoomStatus, err := ctl.client.RoomStatus.
        Query().
        Limit(limit).
        Offset(offset).
        All(context.Background())
        if err != nil {
        c.JSON(400, gin.H{"error": err.Error(),})
        return
    }
  
    c.JSON(200, RoomStatus)
 }


// DeleteRoomStatus handles DELETE requests to delete a roomstatus entity
// @Summary Delete a roomstatus entity by ID
// @Description get roomstatus by ID
// @ID delete-roomstatus
// @Produce  json
// @Param id path int true "RoomStatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstatus/{id} [delete]
func (ctl *RoomStatusController) DeleteRoomStatus(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    err = ctl.client.RoomStatus.
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


  
// UpdateRoomStatus handles PUT requests to update a roomstatus entity
// @Summary Update a roomstatus entity by ID
// @Description update roomstatus by ID
// @ID update-roomstatus
// @Accept   json
// @Produce  json
// @Param id path int true "RoomStatus ID"
// @Param RoomStatus body ent.RoomStatus true "RoomStatus entity"
// @Success 200 {object} ent.RoomStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstatus/{id} [put]
func (ctl *RoomStatusController) UpdateRoomStatus(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    obj := ent.RoomStatus{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "roomstatus binding failed",
        })
        return
    }
    obj.ID = int(id)
    rst, err := ctl.client.RoomStatus.
        UpdateOne(&obj).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed",})
        return
    }
  
    c.JSON(200, rst)
 }
 
// NewRoomStatusController creates and registers handles for the roomstatus controller
func NewRoomStatusController(router gin.IRouter, client *ent.Client) *RoomStatusController {
    uc := &RoomStatusController{
        client: client,
        router: router,
    }
    uc.register()
    return uc
 }
  
 // InitRoomStatusController registers routes to the main engine
 func (ctl *RoomStatusController) register() {
    users := ctl.router.Group("/roomstatus")
  
    users.GET("", ctl.ListRoomStatus)
  
    // CRUD
    users.POST("", ctl.CreateRoomStatus)
    users.GET(":id", ctl.GetRoomStatus)
    users.PUT(":id", ctl.UpdateRoomStatus)
    users.DELETE(":id", ctl.DeleteRoomStatus)
 }
      