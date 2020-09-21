module GradeSchool

type School = Map<int, string list>

let empty: School = Map.empty

let grade number school: string list =
    school
    |> Map.tryFind number
    |> Option.defaultValue []

let add student number school: School =

    let addTo school grade students = Map.add grade students school

    school
    |> grade number
    |> List.append [ student ]
    |> List.sort
    |> addTo school number

let roster school: string list =
    school
    |> Map.toList
    |> List.collect snd
