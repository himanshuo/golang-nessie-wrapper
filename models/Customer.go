package "main"
import (
    "models/Address"
    )

//todo: instead of account id's make it have straight up accounts instead.

type Customer struct
{
    Id string 
    FirstName string 
    LastName string 
    Address Address  
    AccountIds []string
}