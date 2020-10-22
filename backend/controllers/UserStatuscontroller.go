package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/m16_z/app/ent"
   "github.com/m16_z/app/ent/userstatus"
   "github.com/gin-gonic/gin"
)
 
// UserStatusController  defines the struct for the userstatus controller
type UserStatusController  struct {
   client *ent.Client
   router gin.IRouter
}
// CreateUserStatus handles POST requests for adding userstatus entities
// @Summary Create userstatus
// @Description Create userstatus
// @ID create-userstatus
// @Accept   json
// @Produce  json
// @Param userstatus body ent.UserStatus true "UserStatus entity"
// @Success 200 {object} ent.UserStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus [post]
func (ctl *UserStatusController) CreateUserStatus(c *gin.Context) {
    obj := ent.UserStatus{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "UserStatus binding failed",
        })
        return
    }
  
   ust, err := ctl.client.UserStatus.
    Create().
    SetSTATUS(obj.STATUS).
    Save(context.Background())


    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }
  
    c.JSON(200,ust)
 }


// GetUserStatus handles GET requests to retrieve a userstatus entity
// @Summary Get a userstatus entity by ID
// @Description get userstatus by ID
// @ID get-userstatus
// @Produce  json
// @Param id path int true "UserStatus ID"
// @Success 200 {object} ent.UserStatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus/{id} [get]
func (ctl *UserStatusController) GetUserStatus(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
   ust, err := ctl.client.UserStatus.
        Query().
        Where(userstatus.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    c.JSON(200,ust)
 }

// ListUserStatus handles request to get a list of userstatus entities
// @Summary List userstatus entities
// @Description list userstatus entities
// @ID list-userstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.UserStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus [get]
func (ctl *UserStatusController) ListUserStatus(c *gin.Context) {
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
  
    UserStatus, err := ctl.client.UserStatus.
        Query().
        Limit(limit).
        Offset(offset).
        All(context.Background())
        if err != nil {
        c.JSON(400, gin.H{"error": err.Error(),})
        return
    }
  
    c.JSON(200, UserStatus)
 }


// DeleteUserStatus handles DELETE requests to delete a userstatus entity
// @Summary Delete a userstatus entity by ID
// @Description get userstatus by ID
// @ID delete-userstatus
// @Produce  json
// @Param id path int true "UserStatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus/{id} [delete]
func (ctl *UserStatusController) DeleteUserStatus(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    err = ctl.client.UserStatus.
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


  
// UpdateUserStatus handles PUT requests to update a userstatus entity
// @Summary Update a userstatus entity by ID
// @Description update userstatus by ID
// @ID update-userstatus
// @Accept   json
// @Produce  json
// @Param id path int true "UserStatus ID"
// @Param UserStatus body ent.UserStatus true "UserStatus entity"
// @Success 200 {object} ent.UserStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus/{id} [put]
func (ctl *UserStatusController) UpdateUserStatus(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    obj := ent.UserStatus{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "userstatus binding failed",
        })
        return
    }
    obj.ID = int(id)
   ust, err := ctl.client.UserStatus.
        UpdateOne(&obj).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed",})
        return
    }
  
    c.JSON(200,ust)
 }
 
// NewUserStatusController creates and registers handles for the userstatus controller
func NewUserStatusController(router gin.IRouter, client *ent.Client) *UserStatusController {
    uc := &UserStatusController{
        client: client,
        router: router,
    }
    uc.register()
    return uc
 }
  
 // InitUserStatusController registers routes to the main engine
 func (ctl *UserStatusController) register() {
    users := ctl.router.Group("/userstatus")
  
    users.GET("", ctl.ListUserStatus)
  
    // CRUD
    users.POST("", ctl.CreateUserStatus)
    users.GET(":id", ctl.GetUserStatus)
    users.PUT(":id", ctl.UpdateUserStatus)
    users.DELETE(":id", ctl.DeleteUserStatus)
 }
      