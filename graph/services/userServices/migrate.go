package userservices

import (
	custommodels "graph/customModels"
)

func (s *Conn) AutoMigrate() error {
	//if s.db.Migrator().HasTable(&User{}) {
	//	return nil
	//}

	// err := s.db.Migrator().DropTable(&custommodels.Company{}, &custommodels.Job{})
	// if err != nil {
	// 	return err
	// }
	err := s.db.Migrator().AutoMigrate(&custommodels.Users{}, &custommodels.Company{}, &custommodels.Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
