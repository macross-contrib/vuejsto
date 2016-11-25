package handlers

import (
	"github.com/insionng/macross"
	"github.com/macross-contrib/vuejsto/models"
)

func GetMain(self *macross.Context) error {
	return self.Render(macross.StatusOK, "index")
}

func GetTasks(c *macross.Context) error {
	tks, _ := models.GetTasks(0, 0, "id")
	var data = map[string]interface{}{}
	data["items"] = tks
	return c.JSON(macross.StatusOK, data)
}

func PostTask(c *macross.Context) error {
	var task models.Task
	c.Bind(&task)

	id, err := models.PostTask(task.Name)
	if err != nil {
		return err
	}
	var data = map[string]interface{}{}
	data["updated"] = id
	return c.JSON(macross.StatusCreated, data)

}

func PutTask(c *macross.Context) error {
	var task models.Task
	c.Bind(&task)

	if id, err := models.PutTask(task); err == nil {
		var data = map[string]interface{}{}
		data["updated"] = id
		return c.JSON(macross.StatusOK, data)
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
	return self.JSON(macross.StatusOK, data)
}
