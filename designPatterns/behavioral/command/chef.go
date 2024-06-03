package command

import "fmt"

// Chef 厨师 核心计算
type Chef struct {
}

func (c *Chef) StirFry() {
	fmt.Println("Chef StirFry")
}

func (c *Chef) SoupSimmering() {
	fmt.Println("Chef SoupSimmering")
}
