package tango_errors

import "fmt"

func MsgIDNotFound(id int) string {
	return fmt.Sprintf("No record found with the id %d", id)
}

func MsgZeroRecordsFound() string {
	return "Zero records found"
}
