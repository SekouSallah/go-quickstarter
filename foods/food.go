package foods

import (
	"errors"
	"fmt"
)

func FoodOrder(quantity int, title string) (string, error) {

	if quantity <= 0 {
		return "", errors.New("il faut commander au moins un article")
	}

	return fmt.Sprintf("Vous avez commandé %v %v", quantity, title), nil
}
