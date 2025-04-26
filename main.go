package main

import (
	"os"
	"projectgin/controllers"
	"projectgin/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	Init()
	r.Use(initializers.SetupCors())
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/getuser", controllers.RequireAuth, controllers.GetUser)

	r.POST("/workouts", controllers.RequireAuth, controllers.CreateWorkout)
	r.GET("/workouts", controllers.RequireAuth, controllers.GetWorkout)
	r.POST("/updateworkout", controllers.RequireAuth, controllers.UpdateWorkout)
	r.POST("/deletexercise", controllers.RequireAuth, controllers.DeleteExercise)
	r.POST("/deleteworkout", controllers.RequireAuth, controllers.DeleteWorkout)
	r.GET("/exercisesnames", controllers.RequireAuth, controllers.GetExerciseName)

	r.GET("/workoutposts", controllers.RequireAuth, controllers.GetWorkoutsPost)
	r.GET("/workoutposts/:id", controllers.RequireAuth, controllers.GetWorkoutPostsByID)
	r.POST("/workoutposts", controllers.RequireAuth, controllers.CreateWorkoutPost)
	r.POST("/updateworkoutposts/:id", controllers.RequireAuth, controllers.UpdateWorkoutPost)
	r.POST("/deleteworkoutposts/:id", controllers.RequireAuth, controllers.DeleteWorkoutPost)

	if err := os.MkdirAll("./images", os.ModePerm); err != nil {
		panic("Failed to create images directory")
	}

	r.Static("/images", "./images")
	r.POST("/upload", controllers.RequireAuth, controllers.UploadImage)

	r.GET("/logout", controllers.RequireAuth, controllers.Logout)
	r.GET("/getuserdata", controllers.RequireAuth, controllers.GetUserData)

	r.POST("/updateuserdata", controllers.RequireAuth, controllers.UpdateUserData)
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	r.Run(host + ":" + port)
}

func Init() {
	// init func that combines other functions
	initializers.LoadEnvVariables()
	initializers.InitDB()
}
