package handler

import (
	"github.com/Raman5837/kafka.go/app/service"
	"github.com/Raman5837/kafka.go/app/types"
	"github.com/Raman5837/kafka.go/app/utils"
	"github.com/Raman5837/kafka.go/base/config"
	base "github.com/Raman5837/kafka.go/base/utils"
	"github.com/gofiber/fiber/v2"
)

// Produce New Message
func ProduceMessageHandler(context *fiber.Ctx) (exception error) {

	payload := types.ProduceMessageRequestEntity{}

	if validationError := context.QueryParser(&payload); validationError != nil {
		exception := validationError.Error()
		return context.Status(fiber.StatusBadRequest).JSON(base.HttpResponseFail(nil, "Invalid Payload!", exception))
	}

	config := config.GetKafkaConfig()
	assigner := utils.NewPartitionAssigner()
	service := service.NewProducerService(assigner, config.Producer)

	response, additionErr := service.AddNewMessage(&payload)

	if additionErr != nil {
		exception := additionErr.Error()
		return context.Status(fiber.StatusInternalServerError).JSON(base.HttpResponseFail(nil, "Unable To Produce Message", exception))
	}

	return context.Status(fiber.StatusOK).JSON(base.HttpResponseOK(response, "Successfully Produced The New Message"))
}
