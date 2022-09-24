<img src="https://images.pexels.com/photos/8325702/pexels-photo-8325702.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1" width="700" height="450">

# Kemialaskuri
Calculator which excels at calculating chemical formulas.

Work in progress, some documentation down below.
There is a lot dedication to try to write the app as well as possible for my skill level, and also as 
non-repetitively as possible.

## FUTURE FEATURES I AM PLANNING ON (ALSO PHILOSOPHIES I WILL FOLLOW)
* Obviously, it will do much more calculation than right now out-of-the-box
* Straight-forward extensibility
* Simple template for adding new modes
* Better docs for new programmers to be able to extend it
* Cool things to help reduce the repetitiviness. Not sure how to implement these yet, but we'll see. 

## How to compile and use
After running `git clone <URL of repo>` and having this repo cloned on your computer, you have to compile this into a native executable.
This is done by this simple `go` command: `go build .`. However, this assumes *two things*:
* You have **golang** installed on your computer, with a package manager such as **scoop** or from the website.
    * If you are on GNU/Linux, it's as simple as `sudo apt install go` or similar.
* You are currently in the directory, with the source codes. This is done by using `cd` command with your terminal.

After all this, you should have a native executable for your operating system, be it *.exe* or just a regular linux binary.
You can start with it a simple double click, since it uses an infinite loop. You can also just type `./kemialaskuri.exe` on your terminal, because you most likely are already there.

## Documentation
### Directory hirarchy
While the file directory structure is bound to change, this is my vision right now:

Source file | Purpose
---|---
Kemialaskuri.go | Entry point of the app; allows you to select a mode
utils/utils.go | Re-usable functions which do a variety of stuff and help the developer
modes/*.go | All of the modes live here. These are specific calculators you use to calculate formulas.

### Types
This app uses almost exclusively **float64** types for numbers. There are good reasons for this; floats can represent either integers, if rounded down, or well, floats. Since go is a statically typed language, casting and converting values is a lot more harder than on languages like Python. This allows for easy life on the programmer's end and also simplifies the program quite a lot.

**Why am I telling you this?**
Well, I want to make this as customizable/extendable as I can. Everyone needs to be on the same page on this. So yeah, just use float64's.