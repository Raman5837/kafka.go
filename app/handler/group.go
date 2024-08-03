package handler

import (
	"github.com/Raman5837/kafka.go/app/service"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/base/config"
	"github.com/Raman5837/kafka.go/base/utils"
	"github.com/gofiber/fiber/v2"
)

// Add New Consumer Group API Handler
func AddNewConsumerGroupHandler(context *fiber.Ctx) (exception error) {

	payload := types.AddNewConsumerGroupRequestEntity{}

	if validationError := context.QueryParser(&payload); validationError != nil {
		exception := validationError.Error()
		return context.Status(fiber.StatusBadRequest).JSON(utils.HttpResponseFail(nil, "Invalid Payload!", exception))
	}

	config := config.GetKafkaConfig()
	service := service.NewConsumerGroupService(config.ConsumerGroup)
	response, queryErr := service.AddNewConsumerGroup(&payload)

	if queryErr != nil {
		exception := queryErr.Error()
		return context.Status(fiber.StatusBadRequest).JSON(utils.HttpResponseFail(nil, "Something Went Wrong", exception))
	}

	return context.Status(fiber.StatusOK).JSON(utils.HttpResponseOK(response, "Successfully Added New Consumer Group"))

}
