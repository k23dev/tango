package tango_errors

import "fmt"

func MsgIDNotFound(id int) string {
	return fmt.Sprintf("Record not found with the id %d", id)
}

func MsgNotFound() string {
	return fmt.Sprintf("Record not found")
}

func MsgZeroRecordsFound() string {
	return "Zero records found"
}
