# Tables

Data can be organized into tables. Good for test and seed data, but also for any data coming from a database, or for data that can benefit from more organization: A column is array-like, but includes the column name. A row is struct-like, but ???

Example:
```
struct Poster(
    franchise, variant string
    price decimal
)

var posters table<Poster>(
    "Planet of the Apes"    "Statue of Liberty"         11.99
    "Planet of the Apes"    "Marcus, Head of Security"  8.99
    "One Piece"             "Luffy's Bounty"            12.99
    "One Piece"             "Cross Guild"               10.99
    "Studio Ghibli"         "Characters Collage"        15.99
    "Studio Ghibli"         "Kiki's Delivery Service"   12.99
)

func findPostersWithPriceInRange(min, max decimal) (found rows<Poster>) {
    found = posters.Where | $.price >= min and $.price <= min
}

func allPosterPrices() (prices column<decimal>) {
    prices = posters.prices
}

```

## Nil Values

Fundamental and compound data can be assigned values or left nil, and is equivalent to the zero value. A nil value can be used anywhere an assigned zero value can. For a compound type, a nil value means all constituent data is also nil (again, not invalid just equivalent to zero values).
