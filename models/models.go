package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	Engine   *xorm.Engine
	DataType string
	Path     string
)

type Task struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Done int64  `json:"done"`
}

func init() {

	DataType = "mysql"

	var _error error
	if Engine, _error = SetEngine(); _error != nil {
		log.Fatal("Vuejsto.models.init() errors:", _error.Error())
	} else {
		log.Println("Vuejsto already standby!")
	}

	if _error = createTables(Engine); _error != nil {
		log.Fatal("Fail to creatTables errors:", _error.Error())
	}

}

func ConDb() (*xorm.Engine, error) {
	switch {
	case DataType == "sqlite":
		Path = "./data/storage.db"

	case DataType == "mysql":
		Path = "root:@tcp(127.0.0.1:3306)/db?charset=utf8"

	case DataType == "postgres":
		Path = "user=postgres password=yourpass dbname=pgsql sslmode=disable"
	}

	return xorm.NewEngine(DataType, Path)
}

func SetEngine() (*xorm.Engine, error) {
	var _error error
	if Engine, _error = ConDb(); _error != nil {
		return nil, fmt.Errorf("Fail to connect to database: %s", _error.Error())
	} else {
		Engine.SetMapper(core.GonicMapper{})
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 10240)
		Engine.SetDefaultCacher(cacher)

		logPath := path.Join("./logs", "xorm.log")
		os.MkdirAll(path.Dir(logPath), os.ModePerm)
		f, err := os.Create(logPath)
		if err != nil {
			return Engine, fmt.Errorf("Fail to create xorm.log: %s", err.Error())
		}

		Engine.SetLogger(xorm.NewSimpleLogger(f))
		Engine.ShowSQL(false)

		if location, err := time.LoadLocation("Asia/Shanghai"); err == nil {
			Engine.TZLocation = location
		}

		return Engine, err
	}
}

func createTables(Engine *xorm.Engine) error {
	return Engine.Sync2(&Task{})
}

func GetTasks(offset int, limit int, field string) (*[]*Task, error) {
	tks := new([]*Task)
	err := Engine.Limit(limit, offset).Desc(field).Find(tks)
	return tks, err
}

func PostTask(name string) (int64, error) {
	var tk = &Task{Name: name, Done: 0}
	_, err := Engine.Insert(tk)
	return tk.Id, errors.New(fmt.Sprintf("PostTask Error:%v", err))
}

func PutTask(task Task) (int64, error) {
	_, err := Engine.Id(task.Id).Update(task)
	return task.Id, err
}

func DeleteTask(id int64) (int64, error) {
	row, err := Engine.Id(id).Delete(new(Task))
	return row, err
}
