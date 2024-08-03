package errorsutilsv1

import "fmt"

func HandleError(activity, process string, err error) error {

	return fmt.Errorf("[ERROR]-[%s]- %s : %v", activity, process, err)

}
