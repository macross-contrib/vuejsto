package handlers

import (
	"github.com/insionng/macross"
	"github.com/macross-contrib/vuejsto/models"
)

type Data map[string]interface{}

func GetMain(self *macross.Context) error {
	return self.Render(macross.StatusOK, "index")
}

func GetTasks(c *macross.Context) error {
	return c.JSON(macross.StatusOK, models.GetTasks())
}

func PostTask(c *macross.Context) error {
	var task models.Task
	c.Bind(&task)

	id, err := models.PostTask(task.Name)
	if err == nil {
		return c.JSON(macross.StatusCreated, Data{
			"created": id,
		})
	} else {
		return err
	}

}

func PutTask(c *macross.Context) error {
	var task models.Task
	c.Bind(&task)

	if id, err := models.PutTask(task); err == nil {
		return c.JSON(macross.StatusOK, Data{
			"updated": id,
		})
	} else {
		return err
	}
}

func DeleteTask(self *macross.Context) error {
	id := self.Param("id").MustInt()
	_, err := models.DeleteTask(id)

	if err == nil {
		return self.JSON(macross.StatusOK, Data{
			"deleted": id,
		})
	} else {
		return err
	}
}
