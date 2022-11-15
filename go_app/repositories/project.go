package repositories

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
)

type ProjectRepository struct {
	taskMeDB TaskMeDB
}

func (projectRepository ProjectRepository) GetProjectFromTitle(title string) (*models.Project, error) {
	db, err := projectRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM projects WHERE title = '%s' LIMIT 1", title)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var item models.Project
		err := results.Scan(&item.Id, &item.Title, &item.Color)
		if err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, errors.New("Unexpected error user repository")
}

func (projectRepository ProjectRepository) GetProjectFromId(id int) (*models.Project, error) {
	db, err := projectRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM projects WHERE id = '%d' LIMIT 1", id)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var project models.Project
		err := results.Scan(&project.Id, &project.Title, &project.Color)
		if err != nil {
			return nil, err
		}
		return &project, nil
	}

	return nil, errors.New("Unexpected error user repository")
}

func (projectRepository ProjectRepository) GetProjects() ([]models.Project, error) {
	db, err := projectRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	results, err := db.Query("SELECT * FROM projects")
	defer results.Close()
	if err != nil {
		return nil, err
	}

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

func (projectRepository ProjectRepository) AddProject(project models.Project) error {
	db, err := projectRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO projects (title, color) VALUES ('%s','%d')", project.Title, project.Color)
	results, err := db.Query(query)
	defer results.Close()

	return err
}

func (projectRepository ProjectRepository) UpdateProject(project models.Project) error {
	db, err := projectRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE projects SET title = '%s', color = '%d' WHERE id = %d", project.Title, project.Color, project.Id)
	results, err := db.Query(query)
	defer results.Close()

	return err
}

func (projectRepository ProjectRepository) DeleteProject(id int) error {
	db, err := projectRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM projects WHERE id = %d", id)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return err
	}
	return nil

}
