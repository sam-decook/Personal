A basic regular expression implementation following this blog:

https://rhaeguard.github.io/posts/regex/

It supports:
- concatenation
- choice (`|`)
- range (`[a-zBDXY]`)
- groups (`( re )`)
- repetition
    - Kleene (`*`)
    - one or more (`+`)
    - optional (`?`)
    - specified (`{min, max}`) (min or max are optional)

I did little more than copy and paste, so I was inspired to make my own. This one didn't separate lexing from parsing and evaluated straight from the NFA. I wanted to split those two stages up and reduce my NFA to a minimal DFA.