package database

import (
	"sync"

	"github.com/Raman5837/kafka.go/app/model"
)

// type Model Represents All Models Present In This Project
type Model interface {
	String() string
	TableName() string
}

// Reset Tables
func (manager *DatabaseManager) ResetTable(tables []Model) chan error {

	errors := make(chan error)

	for _, table := range tables {
		errors <- manager.SqliteDB.Migrator().DropTable(table)
	}

	return errors

}

// Migrate Given Model
func (manager *DatabaseManager) MigrateTable(model Model) error {
	return manager.SqliteDB.AutoMigrate(&model)
}

// GetModels Returns A List Of All Models In The Project
func GetAllModels() []Model {

	modelTypes := []Model{
		&model.Topic{}, &model.Partition{}, &model.Message{}, &model.LastAssignedPartition{},
		&model.Consumer{}, &model.ConsumerGroup{}, &model.Offset{}, &model.ConsumerAssignment{},
	}

	return modelTypes
}

// Auto Migrate All Models Present
func (manager *DatabaseManager) MigrateAllModels() chan map[string]error {

	allModels := GetAllModels()
	var waitGroup sync.WaitGroup
	migrationErrors := make(chan map[string]error, len(allModels))

	for _, model := range allModels {

		waitGroup.Add(1)

		go func(model Model) {

			defer waitGroup.Done()

			if migrationFailed := manager.MigrateTable(model); migrationFailed != nil {
				migrationErrors <- map[string]error{model.TableName(): migrationFailed}
			}

		}(model)

	}

	// Wait For All Migrations To Complete
	go func() {
		waitGroup.Wait()
		close(migrationErrors)
	}()

	return migrationErrors

}
