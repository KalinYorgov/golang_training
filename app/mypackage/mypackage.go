package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Kalin"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("Invalid day value")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(m int) error {
	if (m < 1) || (m > 12) {
		return errors.New("Invalid month value")
	}
	d.month = m
	return nil
}

func (d *Date) SetYear(y int) error {
	if (y < 1875) || (y > time.Now().Year()) {
		return errors.New("Invalid year value")
	}
	d.year = y
	return nil
}

func (d *Date) GetDay() int {
	return d.day
}

func (d *Date) GetMonth() int {
	return d.month
}

func (d *Date) GetYear() int {
	return d.year
}
