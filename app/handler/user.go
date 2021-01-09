package handler

import (
	"golang-clean-architecture-v1/app/api/user"
	"golang-clean-architecture-v1/app/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAll(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		response := helper.APIResponse("Failed get all user", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if len(users) > 0 {
		response := helper.APIResponse("List of users", http.StatusOK, "success", user.FormatUsers(users))
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := helper.APIResponse("No data users", http.StatusOK, "success", user.FormatUsers(users))
		c.JSON(http.StatusOK, response)
		return
	}
}

func (h *userHandler) GetByID(c *gin.Context) {
	var input user.GetUserDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get detail of user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userDetail, err := h.userService.GetUserByID(input.ID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User detail", http.StatusOK, "success", user.FormatUser(userDetail))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var input user.CreateUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.CreateUser(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Failed to create user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create user", http.StatusOK, "success", user.FormatUser(newUser))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var inputID user.GetUserDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData user.UpdateUserInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.UpdateUser(inputID, inputData)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Update user successfully", http.StatusOK, "success", user.FormatUser(newUser))
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	var input user.GetUserDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.DeleteUser(input.ID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("User has been deleted", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}
