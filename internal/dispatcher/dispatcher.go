package dispatcher

import (
	"fmt"
	"time"
	"url-healthcheck/internal/utils"
	"url-healthcheck/pkg/db"
	pb "url-healthcheck/pkg/grpc"
	"url-healthcheck/pkg/method"
)

type urlDispatcher struct {
	db        db.DB
	urlConfig utils.URLConfig
}

type Task struct {
	utils.URL
	done chan bool
}

// NewURLDispatcher create new dispatcher
func NewURLDispatcher(conn db.DB) (d *urlDispatcher) {
	d = &urlDispatcher{
		db: conn,
	}
	return
}

// StartURLDispatcher launch dispatcher witch check urls
func (d *urlDispatcher) StartURLDispatcher(frequency time.Duration, configFile string) (err error) {
	var urls utils.URLConfig
	for {
		select {
		case <-time.After(frequency * time.Minute):

			urls, err = utils.ReadJSON(configFile)
			if err == nil {
				d.urlConfig = urls
			} else {
				if len(d.urlConfig.Urls) == 0 {
					err = fmt.Errorf("Cannot got url configs")
					return
				}
			}

			fmt.Printf("\n\n### start check urls\n")
			for _, config := range d.urlConfig.Urls {
				task := new(Task)
				task = &Task{config, make(chan bool)}
				var check *db.Check
				check = d.CheckTask(task)
				if err != nil {
					return
				}
				err = d.db.CreateCheck(check)
				if err != nil {
					return
				}
				//<-task.done
			}
			continue
		}
	}
}

// CheckURL launch checking process
func (d *urlDispatcher) CheckTask(task *Task) (check *db.Check) {
	var err error
	var checkMethod string

	count := task.Count
	for _, checkMethod = range task.Checks {
		//fmt.Printf("count %d\n", count)
		if count == 0 {
			break

		} else if checkMethod == pb.Check_CHECK_TYPE_STATUS_CODE.String() && count != 0 {
			err = method.CheckStatusCode(task.Url)
			if err != nil {
				fmt.Printf("%s: %s (%v) %s\n", task.Url, pb.Check_CHECK_STATUS_FAIL.String(), task.Checks, err)
				check = &db.Check{
					Url:     task.Url,
					Type:    checkMethod,
					Status:  pb.Check_CHECK_STATUS_FAIL.String(),
					Comment: err.Error(),
				}
				return
			}

		} else if checkMethod == pb.Check_CHECK_TYPE_TEXT.String() && count != 0 {
			err = method.CheckText(task.Url)
			if err != nil {
				fmt.Printf("%s: %s (%v) %s\n", task.Url, pb.Check_CHECK_STATUS_FAIL.String(), task.Checks[1], err)
				check = &db.Check{
					Url:     task.Url,
					Type:    checkMethod,
					Status:  pb.Check_CHECK_STATUS_FAIL.String(),
					Comment: err.Error(),
				}
				return
			}
		}

		count--
	}

	fmt.Printf("%s: %s\n", task.Url, pb.Check_CHECK_STATUS_OK.String())
	check = &db.Check{
		Url:    task.Url,
		Type:   checkMethod, // fixed last successfull method
		Status: pb.Check_CHECK_STATUS_OK.String(),
	}
	return
}
