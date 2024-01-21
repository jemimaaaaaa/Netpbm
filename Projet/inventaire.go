package main

import (
	"fmt"
)

func main () {
	inventory := (ShowInventory,"Inventory"). (string)
	if item <= 0 || item2 <= 0 || item3 <= 0 {
		return ""
		else "inventory"

		fmt.Println(inventory)

		"potion vie" := 1
		"potion mort" := 2
		"chaudron" := 3
	}
}

func ShowInventory(Inventory string) {
	for key, values := range(inventoryFile, key:"inventory").(map[string]interface{}) {
		if values.(float64) > 0 {
			fmt.Println(key + " => " + strconv.Itoa(int(values.(float64))))
		}
	}
}

func Add0nceItem(inventoryFile string, nameItem, amount int) {
	inventory := (inventoryFile, key:"inventory").(map[string]interface{})
	valeurNouvelle := int(inventory[nameItem].(float64)) + amount
	inventory[nameItem] = valeurNouvelle	
	(inventoryFile, key:"inventory", inventory)
}

func Remove0nceItem(inventoryFile string, nameItem string, amountToRemove int) {
	inventory := (inventoryFile, key:"inventory").(map[string]interface{})
	valeurNouvelle := int(inventory[nameItem].(float64)) - amountToRemove
	inventory[nameItem] = valeurNouvelle
	(inventoryFile, key:"inventory", inventory)
}

func GetAmout0fItem(inventoryFile string, nameItem string) int {
	inventory := (inventoryFile, key:"inventory").(map[string]interface{})
	return int(inventory{nameItem}.(float64))
}
 