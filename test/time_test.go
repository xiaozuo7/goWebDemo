package test

import (
	"fmt"
	"github.com/jinzhu/now"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	createAt, _ := now.Parse(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(createAt)

}
