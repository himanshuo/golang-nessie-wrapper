package "main"


//todo: instead of account id's make it have straight up accounts instead.

type Account struct
{
    Id string 
    Type string 
    Nickname string 
    Rewards int  
    Balance int
    BillIds []string
    CustomerIds []string
}