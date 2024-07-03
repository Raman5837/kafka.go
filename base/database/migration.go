package database

// type Model Represents All Models Present In This Project
type Model interface{}

// Reset Tables
func (manager *DatabaseManager) ResetTable(tables []string) chan error {

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
