package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/m16_z/app/ent"
   "github.com/m16_z/app/ent/roominfo"
   "github.com/gin-gonic/gin"
)
 
// RoomInfoController  defines the struct for the roominfo controller
type RoomInfoController  struct {
   client *ent.Client
   router gin.IRouter
}
// CreateRoomInfo handles POST requests for adding roominfo entities
// @Summary Create roominfo
// @Description Create roominfo
// @ID create-roominfo
// @Accept   json
// @Produce  json
// @Param roominfo body ent.RoomInfo true "RoomInfo entity"
// @Success 200 {object} ent.RoomInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfo [post]
func (ctl *RoomInfoController) CreateRoomInfo(c *gin.Context) {
    obj := ent.RoomInfo{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "RoomInfo binding failed",
        })
        return
    }
  
   rif, err := ctl.client.RoomInfo.
    Create().
    SetINFOBED(obj.INFOBED).
    SetINFOREFRIGERAT(obj.INFOREFRIGERAT).
    SetINFOWARDROB(obj.INFOWARDROB).
    SetINFOOFFICEDE(obj.INFOOFFICEDE).
    Save(context.Background())


    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }
  
    c.JSON(200,rif)
 }


// GetRoomInfo handles GET requests to retrieve a roominfo entity
// @Summary Get a roominfo entity by ID
// @Description get roominfo by ID
// @ID get-roominfo
// @Produce  json
// @Param id path int true "RoomInfo ID"
// @Success 200 {object} ent.RoomInfo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfo/{id} [get]
func (ctl *RoomInfoController) GetRoomInfo(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
   rif, err := ctl.client.RoomInfo.
        Query().
        Where(roominfo.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    c.JSON(200,rif)
 }

// ListRoomInfo handles request to get a list of roominfo entities
// @Summary List roominfo entities
// @Description list roominfo entities
// @ID list-roominfo
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RoomInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfo [get]
func (ctl *RoomInfoController) ListRoomInfo(c *gin.Context) {
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
  
    RoomInfo, err := ctl.client.RoomInfo.
        Query().
        Limit(limit).
        Offset(offset).
        All(context.Background())
        if err != nil {
        c.JSON(400, gin.H{"error": err.Error(),})
        return
    }
  
    c.JSON(200, RoomInfo)
 }


// DeleteRoomInfo handles DELETE requests to delete a roominfo entity
// @Summary Delete a roominfo entity by ID
// @Description get roominfo by ID
// @ID delete-roominfo
// @Produce  json
// @Param id path int true "RoomInfo ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfo/{id} [delete]
func (ctl *RoomInfoController) DeleteRoomInfo(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    err = ctl.client.RoomInfo.
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


  
// UpdateRoomInfo handles PUT requests to update a roominfo entity
// @Summary Update a roominfo entity by ID
// @Description update roominfo by ID
// @ID update-roominfo
// @Accept   json
// @Produce  json
// @Param id path int true "RoomInfo ID"
// @Param RoomInfo body ent.RoomInfo true "RoomInfo entity"
// @Success 200 {object} ent.RoomInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfo/{id} [put]
func (ctl *RoomInfoController) UpdateRoomInfo(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    obj := ent.RoomInfo{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "roominfo binding failed",
        })
        return
    }
    obj.ID = int(id)
   rif, err := ctl.client.RoomInfo.
        UpdateOne(&obj).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed",})
        return
    }
  
    c.JSON(200,rif)
 }
 
// NewRoomInfoController creates and registers handles for the roominfo controller
func NewRoomInfoController(router gin.IRouter, client *ent.Client) *RoomInfoController {
    uc := &RoomInfoController{
        client: client,
        router: router,
    }
    uc.register()
    return uc
 }
  
 // InitRoomInfoController registers routes to the main engine
 func (ctl *RoomInfoController) register() {
    users := ctl.router.Group("/roominfo")
  
    users.GET("", ctl.ListRoomInfo)
  
    // CRUD
    users.POST("", ctl.CreateRoomInfo)
    users.GET(":id", ctl.GetRoomInfo)
    users.PUT(":id", ctl.UpdateRoomInfo)
    users.DELETE(":id", ctl.DeleteRoomInfo)
 }
      