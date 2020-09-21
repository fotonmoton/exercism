module BankAccount

type Account = decimal option ref

let mkBankAccount(): Account = ref None

let openAccount (account: Account): Account = 
    account := Some 0m
    account

let closeAccount account = 
    account := None
    account

let getBalance account = !account 

let updateBalance (change: decimal) (account: Account) = 
    lock account (fun () -> 
        match !account with
        | None -> account := Some change
        | Some balance -> account := Some (balance + change))
    account