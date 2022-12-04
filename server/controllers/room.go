package controllers

import (
	"net/http"

	"planning-poker/server/models"

	"github.com/gin-gonic/gin"
)

type RoomController struct{}

var roomModel = new(models.RoomModel)

// Create ...
func (ctrl RoomController) Create(c *gin.Context) {
	var newRoom models.Room

	// Call BindJSON to bind the received JSON to newRoom.
	if err := c.BindJSON(&newRoom); err != nil {
		return
	}
	roomModel.Create(newRoom)

	c.IndentedJSON(http.StatusCreated, newRoom.ID)
}

// All ...
func (ctrl RoomController) All(c *gin.Context) {
	var rooms []models.Room

	roomModel.All(rooms)

	c.IndentedJSON(http.StatusOK, rooms)
}

// //One ...
// func (ctrl RoomController) One(c *gin.Context) {
// 	userID := getUserID(c)

// 	id := c.Param("id")

// 	getID, err := strconv.ParseInt(id, 10, 64)
// 	if getID == 0 || err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
// 		return
// 	}

// 	data, err := roomModel.One(userID, getID)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Article not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": data})
// }

// //Update ...
// func (ctrl RoomController) Update(c *gin.Context) {
// 	userID := getUserID(c)

// 	id := c.Param("id")

// 	getID, err := strconv.ParseInt(id, 10, 64)
// 	if getID == 0 || err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
// 		return
// 	}

// 	var form forms.CreateArticleForm

// 	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
// 		message := articleForm.Create(validationErr)
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
// 		return
// 	}

// 	err = roomModel.Update(userID, getID, form)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be updated"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Article updated"})
// }

// //Delete ...
// func (ctrl RoomController) Delete(c *gin.Context) {
// 	userID := getUserID(c)

// 	id := c.Param("id")

// 	getID, err := strconv.ParseInt(id, 10, 64)
// 	if getID == 0 || err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
// 		return
// 	}

// 	err = roomModel.Delete(userID, getID)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be deleted"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})

// }
