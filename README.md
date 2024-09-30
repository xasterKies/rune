# ![](resources/rune-lang.png)

## A mini proof of concept bytecode interpreted Programing Language

---

## Why?

I find compilers and interpreters fascinating and I was curious about what happens when you click the compile button in an IDE. In an attempt to understand, I started working on this.

## About 📖
Rune is an intuitive, high-level dynamically typed Programming Language.

---

## Rune's Interpreter & Compiler ⚙️

 Rune lang has an inbuilt interpreter, compiler and interactive shell for invoking and interacting with rune lang's API and instances of rune's runtime environements.

 - ### Interpreter design


 - ### Compiler design
 > Note: Rune's interpreter & compiler are still in development.

 ---

## Features 🧮
Rune supports the following:
- integers
- booleans
- strings
- arrays
- hashes
- prefix-, infix- and index operators
- conditionals
- global and local bindings
- first-class functions
- return statements
- closures

---

 ## Sample language code syntax  🔨

 ```javascript
 let name = "rune";
let version = 1.0;
let inspirations = ["Scheme", "Lisp", "JavaScript"];
let lang = {
"title": "rune",
"version": "1.0",
"author": "xasterKies"
};

let printLangName = fn(title) {
let title = book["title"];
let author = book["author"];
puts(author + " - " + title);
};

printLangName(title);
// => prints: "xasterKies - rune"
 ```

---

## Installation 💾:

### Source 📜:

> You need **Go** 1.21.6+ installed.

```sh
# clone the repo
$ git clone https://github.com/xasterKies/rune/

# Navigate into the repo
$ cd rune

# Run the REPL(Read Eval Print Loop)
$ go run main.go


```

---

## Contributing 👥 🔧:

> Note: You need **Go** 1.21.6+ and **git** installed.<br>

```sh
# clone the repo and Navigate into repo
$ git clone https://github.com/xasterKies/rune/ && cd rune

# create a branch to host your feature

```

> Just hack on it as you wish!

The rune folder has the following structure:-

```sh
$ tree
rune
.
| -ast #abstract syntax tree
| -lexer
| -parser
| -repl #read eval print loop
| -token 
| -resources #random associated files to rune
.
```

>Note: This structure can change since its still in development.

Feel free to open a PR or issue to discuss any bugs 🪲, improvements 📈, Ideas 💡, etc.

If you find a typo, help me fix it... Thanks.



