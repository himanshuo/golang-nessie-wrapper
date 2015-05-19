package "main"


//todo: change date's from strings to golang date types

type Bill struct
{
    Id string 
    Status string 
    Payee string 
    Nickname string  
    CreationDate string
    PaymentDate string
    RecurringDate string
    UpcomingPaymentDate string
    PaymentAmount int
}