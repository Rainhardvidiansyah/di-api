package handler

//
//func SetupRoutes(db *gorm.DB) *gin.Engine {
//	r := gin.Default()
//	r.Use(func(c *gin.Context) {
//		c.Set("db", db)
//	})
//	r.GET("/tasks")
//	r.POST("/tasks", controllers.CreateTask)
//	r.GET("/tasks/:id", controllers.FindTask)
//	r.PATCH("/tasks/:id", controllers.UpdateTask)
//	r.DELETE("tasks/:id", controllers.DeleteTask)
//	return r
//}
