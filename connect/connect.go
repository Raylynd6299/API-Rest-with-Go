package connect

import (
	"log"

	"../structures"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var connection *gorm.DB

const engineSQL string = "mysql"

const username string = "root"
const password string = "pass"//ingresar password
const database string = "REST"

//"username:password@/database"

func InitializaDatabase() {
	connection = ConnectORM(CreateSSt())
	log.Println("La coneccion con la baase de datos fue exitosa")
}
func CloseConnectionBD() {
	connection.Close()
	log.Println("La coneccion con la baase de datos fue cerrada")

}

//ConnectORM crea la coneccion con la base de datos
func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open(engineSQL, stringConnection)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection
}

//CreateSSt construye la cadena especial para abrir canal con BD
func CreateSSt() string {
	return username + ":" + password + "@/" + database
}

func DeleteUser(id string) {
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	connection.Delete(&user)
}
func UpdateUser(id string, user structures.User) structures.User {
	currentuser := structures.User{}
	connection.Where("id = ?", id).First(&currentuser)
	currentuser.Username = user.Username
	currentuser.FirstName = user.FirstName
	currentuser.LastName = user.LastName

	connection.Save(&currentuser)

	return currentuser
}

func GetUser(id string) structures.User {
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)

	return user
}
func CreateUser(us structures.User) structures.User {
	connection.Create(&us) // Se asigna un id
	return us
}
