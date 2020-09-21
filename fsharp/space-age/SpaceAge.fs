module SpaceAge

type Planet =
    | Earth
    | Mercury
    | Venus
    | Mars
    | Jupiter
    | Saturn
    | Uranus
    | Neptune

let secondsPerEarthYear = 365.25 * 24. * 60. * 60.

let planetCoefficients = 
    [
        Earth,      1.0
        Mercury,    0.2408467
        Venus,      0.61519726
        Mars,       1.8808158
        Jupiter,    11.862615
        Saturn,     29.447498
        Uranus,     84.016846
        Neptune,    164.79132
    ] |> Map.ofList

let age (planet: Planet) (seconds: int64) =
    float seconds / (secondsPerEarthYear * planetCoefficients.[planet])