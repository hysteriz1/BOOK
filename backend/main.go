package main
 
import (
   "context"
   "fmt"
   "log"
 
   "github.com/m16_z/app/controllers"
   _ "github.com/m16_z/app/docs"
   "github.com/m16_z/app/ent"
   "github.com/m16_z/app/ent/roomtype"
   "github.com/m16_z/app/ent/roominfo"
   "github.com/m16_z/app/ent/roomstatus"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	USERNAME  string
	USEREMAIL string
}


type Bookstatuses struct {
	Bookstatus []Bookstatus
}

type Bookstatus struct {
	BOOKSTATUSDATA     string
}

type Rooms struct {
	Room []Room
}

type Room struct {
	ROOMNAME string
	ROOMTYPEID int
	ROOMINFOID int
	ROOMSTATUSID int
}

type RoomInfos struct {
	RoomInfo []RoomInfo
}

type RoomInfo struct {
    INFOBED int
    INFOREFRIGERAT int
    INFOWARDROB int
    INFOOFFICEDE int
}
type RoomStatuses struct {
	RoomStatus []RoomStatus
}

type RoomStatus struct {
    STATUSDATA string
}

type RoomTypes struct {
	RoomType []RoomType
}

type RoomType struct {
    ROOMTYPEDATA string
    COSTDATA int
}

type UserStatuses struct {
	UserStatus []UserStatus
}

type UserStatus struct {
    STATUS string
}


// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
   controllers.NewUserController(v1, client)
   controllers.NewUserStatusController(v1, client)
   controllers.NewRoomController(v1, client)
   controllers.NewRoomInfoController(v1, client)
   controllers.NewRoomStatusController(v1, client)
   controllers.NewRoomTypeController(v1, client)
   controllers.NewBookController(v1, client)
   controllers.NewBookstatusController(v1, client)
   
	// Set Users Data
	users := Users{
		User: []User{
			User{"นายวัชระพงศ์ ทาระมล", "B6107505@gmail.com"},
			User{"นายฟูฟู หมาบิน", "fufu@example.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetUSERNAME(u.USERNAME).
			SetUSEREMAIL(u.USEREMAIL).
			Save(context.Background())
	}

	// Set RoomStatuses Data
	roomstatuses := RoomStatuses{
		RoomStatus: []RoomStatus{
			RoomStatus{"ว่าง"},
			RoomStatus{"ไม่ว่าง"},
		},
	}

	for _, rst := range roomstatuses.RoomStatus {
		client.RoomStatus.
			Create().
			SetSTATUSDATA(rst.STATUSDATA).
			Save(context.Background())
	}

	// Set Bookstatuses Data
	bookstatuses := Bookstatuses{
		Bookstatus: []Bookstatus{
			Bookstatus{"กำลังดำเนินการ"},
			Bookstatus{"ดำเนินการเสร็จสิ้น"},
		},
	}

	for _, bst := range bookstatuses.Bookstatus {
		client.Bookstatus.
			Create().
			SetBOOKSTATUSDATA(bst.BOOKSTATUSDATA).
			Save(context.Background())
	}
	
	// Set RoomTypes Data
	roomtypes := RoomTypes{
		RoomType: []RoomType{
			RoomType{"ห้องพัดลม",2500},
			RoomType{"ห้องแอร์",3500},	
		},
	}
	for _, rt := range roomtypes.RoomType {
		client.RoomType.
			Create().
			SetROOMTYPEDATA(rt.ROOMTYPEDATA).
			SetCOSTDATA(rt.COSTDATA).
			Save(context.Background())
	}
		
	// Set RoomInfos Data
	roominfos := RoomInfos{
		RoomInfo: []RoomInfo{
			RoomInfo{1, 1, 1, 1},
			RoomInfo{2, 1, 1, 1},	
		},
	}
	for _, rif := range roominfos.RoomInfo {
		client.RoomInfo.
			Create().
			SetINFOBED(rif.INFOBED).
			SetINFOREFRIGERAT(rif.INFOREFRIGERAT).			
			SetINFOWARDROB(rif.INFOWARDROB).
			SetINFOOFFICEDE(rif.INFOOFFICEDE).
			Save(context.Background())
	}

	// Set Rooms Data
	rooms := Rooms{
		Room: []Room{
			Room{"001", 1, 1, 1},
			Room{"002", 2, 2, 1},
			Room{"003", 1, 1, 1},
		},
	}
	for _, r := range rooms.Room {

		rt, err := client.RoomType.
			Query().
			Where(roomtype.IDEQ(int(r.ROOMTYPEID))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		
		rif, err := client.RoomInfo.
			Query().
			Where(roominfo.IDEQ(int(r.ROOMINFOID))).
			Only(context.Background())

			if err != nil {
				fmt.Println(err.Error())
				return
		}

		rst, err := client.RoomStatus.
				Query().
				Where(roomstatus.IDEQ(int(r.ROOMSTATUSID))).
				Only(context.Background())
	
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		client.Room.
			Create().
			SetROOMNAME(r.ROOMNAME).
			SetROOMINFO(rif).
			SetROOMROOMTYPE(rt).
			SetROOMSTATUS(rst).
			Save(context.Background())
	}

   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}
