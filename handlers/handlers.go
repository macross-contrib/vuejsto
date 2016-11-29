package handlers

import (
	"github.com/insionng/macross"
	"github.com/macross-contrib/vuejsto/models"
)

func GetMain(self *macross.Context) error {
	return self.Render("index")
}

func GetTasks(self *macross.Context) error {
	tks, _ := models.GetTasks(0, 0, "id")
	var data = map[string]interface{}{}
	data["items"] = tks
	return self.JSON(data)
}

func PostTask(self *macross.Context) error {
	var task models.Task
	self.Bind(&task)

	id, err := models.PostTask(task.Name)
	if err != nil {
		return err
	}
	var data = map[string]interface{}{}
	data["updated"] = id
	return self.JSON(data, macross.StatusCreated)

}

func PutTask(self *macross.Context) error {
	var task models.Task
	self.Bind(&task)

	if id, err := models.PutTask(task); err == nil {
		var data = map[string]interface{}{}
		data["updated"] = id
		return self.JSON(data)
	} else {
		return err
	}
}

func DeleteTask(self *macross.Context) error {
	id := self.Param("id").MustInt64()
	if _, err := models.DeleteTask(id); err != nil {
		return err
	}
	var data = map[string]interface{}{}
	data["deleted"] = id
	return self.JSON(data)
}
