<img src="https://images.pexels.com/photos/8325702/pexels-photo-8325702.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1" width="700" height="450">

# Kemialaskuri
Calculator which excels at calculating chemical formulas.

Work in progress.

## How to compile and use
After running `git clone <URL of repo>` and having this repo cloned on your computer, you have to compile this into a native executable.
This is done by this simple `go` command: `go build .`. However, this assumes *two things*:
* You have **golang** installed on your computer, with a package manager such as **scoop** or from the website.
    * If you are on GNU/Linux, it's as simple as `sudo apt install go` or similar.
* You are currently in the directory, with the source codes. This is done by using `cd` command with your terminal.

After all this, you should have a native executable for your operating system, be it *.exe* or just a regular linux binary.
You can start with it a simple double click, since it uses an infinite loop. You can also just type `./kemialaskuri.exe` on your terminal, because you most likely are already there.