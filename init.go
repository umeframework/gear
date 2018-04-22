package gear

import (
	"fmt"
	"github.com/umeframework/gear/orm"
)

// Framework initialize
func init() {
	fmt.Println("Gear init...")
}

// Create Orm
func GetDao() *orm.Orm {
	return &orm.Orm{}
}

