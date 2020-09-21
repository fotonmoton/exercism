module Leap

let leapYear (year: int): bool =
    let by4 = year % 4 = 0
    let by100 = year % 100 = 0
    let by400 = year % 400 = 0
    by4 && not by100 || by400 