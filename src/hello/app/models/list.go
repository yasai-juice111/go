package models

import (
        "time"
)

type List struct {
        id        		uint64     `gorm:"primary_key" json:"id"`
        name  			string     `sql:"size:64" json:"name" validate:"max=64"`
        createDatatime 	time.Time  `json:"create_datetime"`
        updateDatetime 	time.Time  `json:"update_datetime"`
        deleteFlag 		*time.Time `json:"delete_flag"`
}