package controllers
 
import (
	"github.com/m16_z/app/ent/roominfo"
	"github.com/m16_z/app/ent/roomstatus"
	"github.com/m16_z/app/ent/roomtype"
   "context"
   "fmt"
   "strconv"
   "github.com/m16_z/app/ent"
   "github.com/m16_z/app/ent/room"
   "github.com/gin-gonic/gin"
)
 
// RoomController  defines the struct for the room controller
type RoomController  struct {
   client *ent.Client
   router gin.IRouter
}
type Room struct{
    ROOMNAME string
    RoomType int
    RoomStatus int
    RoomInfo int
}
// CreateRoom handles POST requests for adding room entities
// @Summary Create room
// @Description Create room
// @ID create-room
// @Accept   json
// @Produce  json
// @Param room body ent.Room true "Room entity"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /room [post]
func (ctl *RoomController) CreateRoom(c *gin.Context) {
    obj := Room{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "Room binding failed",
        })
        return
    }
    rt, err := ctl.client.RoomType.
    Query().
    Where(roomtype.IDEQ(int(obj.RoomType))).
    Only(context.Background())

    if err != nil {
        c.JSON(400, gin.H{
            "error": "roomtype not found",
        })
        return
    }
    rst, err := ctl.client.RoomStatus.
    Query().
    Where(roomstatus.IDEQ(int(obj.RoomStatus))).
    Only(context.Background())

    if err != nil {
        c.JSON(400, gin.H{
            "error": "roomstatus not found",
        })
        return
    }
    rif, err := ctl.client.RoomInfo.
    Query().
    Where(roominfo.IDEQ(int(obj.RoomInfo))).
    Only(context.Background())

    if err != nil {
        c.JSON(400, gin.H{
            "error": "roominfo not found",
        })
        return
    }
    r, err := ctl.client.Room.
    Create().
    SetROOMNAME(obj.ROOMNAME).
    SetROOMROOMTYPE(rt).
    SetROOMSTATUS(rst).
    SetROOMINFO(rif).
    Save(context.Background())

    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }
  
    c.JSON(200, r)
 }


// GetRoom handles GET requests to retrieve a room entity
// @Summary Get a room entity by ID
// @Description get room by ID
// @ID get-room
// @Produce  json
// @Param id path int true "Room ID"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /room/{id} [get]
func (ctl *RoomController) GetRoom(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    r, err := ctl.client.Room.
        Query().
        Where(room.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    c.JSON(200, r)
 }

// ListRoom handles request to get a list of room entities
// @Summary List room entities
// @Description list room entities
// @ID list-room
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /room [get]
func (ctl *RoomController) ListRoom(c *gin.Context) {
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
  
    Room, err := ctl.client.Room.
        Query().
        WithROOMINFO().
        WithROOMROOMTYPE().
        WithROOMSTATUS().
        Limit(limit).
        Offset(offset).
        All(context.Background())
        if err != nil {
        c.JSON(400, gin.H{"error": err.Error(),})
        return
    }
  
    c.JSON(200, Room)
 }


// DeleteRoom handles DELETE requests to delete a room entity
// @Summary Delete a room entity by ID
// @Description get room by ID
// @ID delete-room
// @Produce  json
// @Param id path int true "Room ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /room/{id} [delete]
func (ctl *RoomController) DeleteRoom(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    err = ctl.client.Room.
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


  
// UpdateRoom handles PUT requests to update a room entity
// @Summary Update a room entity by ID
// @Description update room by ID
// @ID update-room
// @Accept   json
// @Produce  json
// @Param id path int true "Room ID"
// @Param Room body ent.Room true "Room entity"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /room/{id} [put]
func (ctl *RoomController) UpdateRoom(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    obj := ent.Room{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "room binding failed",
        })
        return
    }
    obj.ID = int(id)
    r, err := ctl.client.Room.
        UpdateOne(&obj).
        SetROOMSTATUSID(2).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed",})
        return
    }
  
    c.JSON(200, r)
 }
 
// NewRoomController creates and registers handles for the room controller
func NewRoomController(router gin.IRouter, client *ent.Client) *RoomController {
    uc := &RoomController{
        client: client,
        router: router,
    }
    uc.register()
    return uc
 }
  
 // InitRoomController registers routes to the main engine
 func (ctl *RoomController) register() {
    users := ctl.router.Group("/room")
  
    users.GET("", ctl.ListRoom)
  
    // CRUD
    users.POST("", ctl.CreateRoom)
    users.GET(":id", ctl.GetRoom)
    users.PUT(":id", ctl.UpdateRoom)
    users.DELETE(":id", ctl.DeleteRoom)
 }
      