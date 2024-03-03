package main

import (
	//"admin/database"
	"admin/database"
	"admin/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
 

func main() {
    database.Connect()

   // fmt.Println(db)// n1, n2 := two()//fmt.Println(n1,n2)

    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowCredentials:  true,
        AllowOriginsFunc: func(origin string) bool {
            return true
          },
        AllowOrigins: "http://localhost:3000",
       // AllowOrigins: []string{"https://example.com", "https://another-example.com"},
    }))

    routes.Setup(app)

    

    app.Listen(":8000")
}


// func two() (int,int){
//     return 3,5
// }
// func main() {
//     app := *fiber.New()

//     app.Get("/example", func(c fiber.Ctx) error {
//         // Your logic here
//         data := fiber.Map{"message": "Hello, Fiber!"}

//         // Use c.JSON on the non-pointer *fiber.Ctx
//         return c.JSON(data)
//     })

//     app.Listen(":3000")
// }
