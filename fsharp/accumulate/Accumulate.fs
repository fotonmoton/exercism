module Accumulate

let accumulate func input =

    // to utilize tail call recursion
    let rec mapish f state list =
        match list with
        | [] -> state
        | head :: tail -> mapish f (f head :: state) tail

    let rev = mapish id []

    input
    |> mapish func []
    |> rev