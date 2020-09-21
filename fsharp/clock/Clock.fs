module Clock

type Clock = int

let private minutesPerHour = 60

let private minutesPerDay = minutesPerHour * 24

let private modulo x y = ((x % y) + y) % y

let create hours minutes: Clock = 
    modulo (hours * minutesPerHour + minutes) minutesPerDay

let add minutes clock: Clock = create 0 (clock + minutes)

let subtract minutes clock: Clock = add -minutes clock

let display clock =
    let hours = clock / minutesPerHour
    let minutes = clock % minutesPerHour
    sprintf "%02d:%02d" hours minutes