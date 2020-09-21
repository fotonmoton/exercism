type allergen = Eggs
              | Peanuts
              | Shellfish
              | Strawberries
              | Tomatoes
              | Chocolate
              | Pollen
              | Cats

let allergen_score = [
    (Cats, 128);
    (Pollen, 64);
    (Chocolate, 32);
    (Tomatoes, 16);
    (Strawberries, 8);
    (Shellfish, 4);
    (Peanuts, 2);
    (Eggs, 1);
]

let allergic_to score allergen =
    List.fold_left (fun acc (a, s) -> 
        match score mod s <> 0 with
        | true -> true
        | false -> acc )

let allergies _ =
    failwith "'allergies' is missing"
