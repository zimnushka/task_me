package repositories

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
)

type ProjectUserRepository struct {
	taskMeDB TaskMeDB
}

func (repository ProjectUserRepository) GetProjectsByUser(id int) ([]models.Project, error) {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT projects.id, projects.title,projects.color FROM UsersProjects INNER JOIN projects ON UsersProjects.project_id=projects.id AND UsersProjects.user_id=%d", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Project, itemsLng)

	for results.Next() {
		var item models.Project
		err := results.Scan(&item.Id, &item.Title, &item.Color)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (repository ProjectUserRepository) GetUsersByProject(id int) ([]models.User, error) {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT users.id, users.name, users.email FROM UsersProjects INNER JOIN users ON UsersProjects.user_id=users.id AND UsersProjects.project_id=%d", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.User, itemsLng)

	for results.Next() {
		var item models.User
		err := results.Scan(&item.Id, &item.Name, &item.Email)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (repository ProjectUserRepository) AddLink(projectId, userId int) error {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO UsersProjects (user_id, project_id) VALUES ('%d','%d')", userId, projectId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}

func (repository ProjectUserRepository) DeleteLink(projectId, userId int) error {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM UsersProjects WHERE user_id = %d AND project_id = %d", userId, projectId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}

func (repository ProjectUserRepository) DeleteAllLinkByProject(projectId int) error {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM UsersProjects WHERE project_id = %d", projectId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err

}

func (repository ProjectUserRepository) DeleteAllLinkByUser(userId int) error {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM UsersProjects WHERE user_id = %d", userId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err

}
