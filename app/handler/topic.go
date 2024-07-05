package handler

import (
	"github.com/Raman5837/kafka.go/app/service"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/utils"
	"github.com/gofiber/fiber/v2"
)

// Get Topic Object
func GetTopicHandler(context *fiber.Ctx) (exception error) {

	payload := types.GetTopicRequestEntity{}

	if validationError := context.QueryParser(&payload); validationError != nil {
		exception := validationError.Error()
		return context.Status(fiber.StatusBadRequest).JSON(utils.HttpResponseFail(nil, "Invalid Payload!", exception))
	}

	service := service.NewTopicService()
	response, queryErr := service.GetTopic(&payload)

	if queryErr != nil {
		exception := queryErr.Error()
		return context.Status(fiber.StatusBadRequest).JSON(utils.HttpResponseFail(nil, "Topic Not Found", exception))
	}

	return context.Status(fiber.StatusOK).JSON(utils.HttpResponseOK(response, "Successfully Fetched Requested Topic"))
}

// Create Topic Object
func CreateTopicHandler(context *fiber.Ctx) (exception error) {

	payload := types.CreateTopicRequestEntity{}

	if validationError := context.QueryParser(&payload); validationError != nil {
		exception := validationError.Error()
		return context.Status(fiber.StatusBadRequest).JSON(utils.HttpResponseFail(nil, "Invalid Payload!", exception))
	}

	service := service.NewTopicService()
	response, queryErr := service.AddNewTopic(&payload)

	if queryErr != nil {
		exception := queryErr.Error()
		return context.Status(fiber.StatusBadRequest).JSON(utils.HttpResponseFail(nil, "Unable To Create New Topic", exception))
	}

	return context.Status(fiber.StatusOK).JSON(utils.HttpResponseOK(response, "Successfully Created New Topic"))
}
