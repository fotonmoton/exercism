module TwoFer

let twoFer name =
    name 
    |> Option.defaultValue "you" 
    |> sprintf "One for %s, one for me." 
