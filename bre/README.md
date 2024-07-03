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