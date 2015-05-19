package "main"
import (
    "models/Address"
    )



type Branch struct
{
    Id string 
    Name string 
    Address Address 
    PhoneNumber string 
    Hours []string
}