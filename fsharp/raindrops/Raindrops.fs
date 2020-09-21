module Raindrops

// type Raindrop = (int, String)

let convert number =
    
    let getSound (divider, sound) = 
        if (number % divider = 0) 
        then sound 
        else "" 

    let soundString drops =
        drops 
        |> List.map getSound
        |> List.reduce (+)

    let sound = soundString [(3, "Pling"); (5, "Plang"); (7, "Plong")]
    
    if sound = ""
    then string number
    else sound