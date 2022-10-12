package calendar

import "fmt"

//creating a date struct that contains 3 fields
type Date struct{
	year int
	month int
	day int /*the rreason behind un-exporting these
	fields directly is to avoid users from inserting
	invalid values in the struct call at the main block*/
}
 
//creating a setter method to set fields values and store
//in the getter method
func (d *Date) SetYear(year int)error {
	if year < 1{
		return fmt.Errorf("please enter a valid year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int)error{
	if month < 1 || month > 12{
        return fmt.Errorf("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int)error{
	if day < 1 || day > 31{
		return fmt.Errorf("invalid day")
	}
	d.day = day
	return nil
}



//creating a getter method to get values from get method
func (d *Date) Year()int{
	return d.year
}
func (d *Date)Month()int{
	return d.month
}
func (d *Date)Day()int{
	return d.day
}



