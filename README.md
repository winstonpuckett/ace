# The Ace Programming Language

A tiny stack-based, forth-like language.

# Examples

Demonstrating everything in the language:
```ace
:helloMessage "Hello, "
:user $USER

:printTwo `echo +1 +2`

helloMessage user print
```

More simply:
```ace
`echo Hello, $USER`
```


# Learn

In ace, we can:
- Declare words by starting a line with a colon and ending with two new line characters: `:this is a word`
- Declare strings between two double quotes: `"This is a string"`
- Declare calls between two tick marks: `\`this is an external call\``
- Add from the stack to a call with a plus and the number on the stack: `echo +1 +2 +1`
- Call a declared words by its name: `call`
- Reference environment variables with dollar sign: `$VARIABLE`

# Install

*Note: ace only works on LF, not CRLF*

Make sure you have golang installed, then

```bash
git clone https://github.com/winstonpuckett/ace
cd ace
go build
```
# Why/where is this going?

I'm not satisfied with today's build tools. They are either tightly bound to an ecosystem or not well documented.

The idea for this language was to create an easy to learn, run-anywhere, fast build system. 

# TODO

- [ ] Automated testing for all current features
- [ ] Handle parsing errors gracefully
- [ ] Handle CRLF
- [ ] Package Management
- [ ] Idempotency?
- [ ] Pipeline
- [ ] Easier install

