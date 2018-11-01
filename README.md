## Go sekwence library

Sekwence is a library for parsing bash-like patterns `{00..12}`, `{1,3,5,7}` and generate corresponding string sequences.

##### Example

```
sekwence.ExpandPattern("host{00..03}{a,c,f}.example.com")
> host00a.example.com
> host00c.example.com
> host00f.example.com
> host01a.example.com
> host01c.example.com
> host01f.example.com
> host02a.example.com
> host02c.example.com
> host02f.example.com
> host03a.example.com
> host03c.example.com
> host03f.example.com
```
