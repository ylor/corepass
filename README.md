# corepass - generate chunkable passwords

corepass is a password generator that does [one thing well](https://en.wikipedia.org/wiki/Unix_philosophy): generate secure and chunkable passwords.

## Usage

Download the release for your operating system and architecture and invoke:

```
➜ corepass
Netwuc-tolnyx3
```

```
➜ corepass --strong
riplog6-Dyskiv-jogmeb
```

## FAQ

### What do you mean secure?

- By default generated passwords have 49 Shannon entropy bits, and when using the `--strong` flag the entropy increases to 86. Generated passwords include lowercase characters, an uppercase character, a number, and a symbol.

### What do you mean chunkable?

- Pronouncable syllables construct memorable phrases by using a consonant, vowel, consonant pattern. Phrases are separated by hyphens to allow for intuitive chunking. Read more about this pattern [here](https://rmondello.com/2024/10/07/apple-passwords-generated-strong-password-format/).

### What's with the name?

- This approach to the password UX first came about on Apple platforms.

### Why did you make this?

- I wanted to learn Go.

### Why did you _really_ make this?

- Due to [sickness](https://playvalorant.com), I frequently hop between systems and found myself wanting to generate this kind of password on any OS.

### Is it good?

- No.

## Acknowledgements

- [Ricky Mondello](https://rmondello.com) - they work on authentication technologies at Apple and generously shared details on the design of this password pattern.
