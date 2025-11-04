package main

import "fmt"

func main() {
	things := make(map[string]*Product)
	for {
		var contral string
		var all float64
		var thing Product
		fmt.Print("请输入你要干嘛？(售卖/进货/exit): ")
		fmt.Scanln(&contral)
		if contral == "售卖" {
			fmt.Print("售卖商品:")
			fmt.Scanln(&thing.Name)
			fmt.Print("售卖数量:")
			fmt.Scanln(&thing.Stock)
			if Thing, ok := things[thing.Name]; ok {
				fmt.Println(Thing.Sell(thing.Stock))
				if (*Thing).Stock == 0 {
					delete(things, thing.Name)
				}
			} else {
				fmt.Println("没有可售卖的商品")
			}
			for _, Thing := range things {
				fmt.Println((*Thing).Info())
			}
		} else if contral == "进货" {
			fmt.Print("进货商品:")
			fmt.Scanln(&thing.Name)
			fmt.Print("进货数量:")
			fmt.Scanln(&thing.Stock)
			if Thing, ok := things[thing.Name]; ok {
				Thing.Restock(thing.Stock) //调用方法赋值
			} else {
				fmt.Print("售卖单价:")
				fmt.Scanln(&thing.Price)
				things[thing.Name] = &thing
			}
			for _, Thing := range things {
				fmt.Println((*Thing).Info())
			}
		} else if contral == "exit" {
			return
		} else {
			fmt.Println("请输入正确的指令")
		}
		for _, Things := range things {
			all += Things.TotalValue()
		}
		fmt.Printf("商品总价值为%.2f\n", all)
	}
}

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (all Product) TotalValue() float64 { //判断总价值
	total := float64(all.Stock) * all.Price
	return total

}
func (Stock Product) IsInStock() bool { //判断库存
	return Stock.Stock >= 0
}
func (infor Product) Info() string { //显示商品信息
	return fmt.Sprintf("商品:%s, 单价:%.2f,库存:%d\n", infor.Name, infor.Price, infor.Stock)
}
func (Input *Product) Restock(amount int) { //增加库存
	Input.Stock += amount
}
func (Message *Product) Sell(amount int) (success bool, message string) { //用来售卖
	Message.Stock = Message.Stock - amount
	if !Message.IsInStock() { //没有库存
		success = false
		message = "库存不足"
		Message.Stock = Message.Stock + amount
		return
	}
	success = true
	message = "售买成功"
	return
}
