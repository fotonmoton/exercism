module Bob

type Response =
    | Agreement
    | Indifference
    | Shout
    | Confidence
    | Obedience

let private responseToString =
    function
    | Obedience -> "Fine. Be that way!"
    | Confidence -> "Calm down, I know what I'm doing!"
    | Agreement -> "Sure."
    | Shout -> "Whoa, chill out!"
    | Indifference -> "Whatever."

let private flip d f = f d

let private stringToUpper = String.map System.Char.ToUpper

let private trim = String.filter (System.Char.IsWhiteSpace >> not)

let private isEmpty clause = clause = ""

let private hasChars = String.exists System.Char.IsLetter

let private isShout clause = clause = stringToUpper clause && hasChars clause

let private isQuestion clause = Seq.last clause = '?'

let private isForcefulQuestion clause = [ isShout; isQuestion; hasChars ] |> List.forall (flip clause)

let response (input: string): string =
    match trim input with
    | i when isEmpty i -> Obedience
    | i when isForcefulQuestion i -> Confidence
    | i when isQuestion i -> Agreement
    | i when isShout i -> Shout
    | _ -> Indifference
    |> responseToString
