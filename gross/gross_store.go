package main

func main()  {
	
}

func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
	}
}

func NewBill() map[string]int {
	return map[string] int{}
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
    if !exists{
        return false
    }
    bill[item] += unitValue
	return true
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
    billValue, itemExists := bill[item]
    	if !itemExists{
            return false
        }
    unitValue, unitExists := units[unit]
    if !unitExists{
        return false
    }
    newQuantity := billValue - unitValue
    if newQuantity < 0{
        return false
    }
    if newQuantity == 0 {
		delete(bill, item)
		return true
	}
    bill[item] = newQuantity
    return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	billValue, itemExists := bill[item]
    if !itemExists{
        return 0, false
    }
    return billValue, true
}