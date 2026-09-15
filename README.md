# Go Virus Prank

This is a small virus prank to scare your friends, built in Go.
It is based off of my earlier project of building a [virus prank in python](https://github.com/MrCheesie/python-virus-prank)

> [!NOTE]
> This is a completely harmless program, and IT IS NOT DANGEROUS AT ALL.

## What it does
This program will scare your friends by speaking random phrases at set intervals. It uses the default text-to-speech command `say` for MacOS, and
`voice.Speak` for Windows, hence it may not run as intended for older versions of the OS.

## How to use
Download the appropriate file for your operating system:
- [MacOS](https://github.com/MrCheesie/go-virus-prank/releases/download/release/virus_MacOS)
- [Windows](https://github.com/MrCheesie/go-virus-prank/releases/download/release/virus_WIndows_ARM64.exe)

and double-click on them to start the program.

## Building from source

First, ensure Go is installed on your device, if not, grab it from [Go](https://go.dev). If not installed already, also get [git](https://git-scm.com).

After installing Go and Git, run
```bash
git clone https://github.com/MrCheesie/go-virus-prank.git
cd python-virus-prank
```

and

### MacOS
```bash
go run virus.go
```
to run or
```bash
go build virus.go 
```
to build the executable.


### Windows
```bash
go run winVirus.go
```
to run or
```bash
go build winVirus.go 
```
to build the executable.

## License

Licensed under the MIT license. Do whatever you want with it.

## FAQ's 

<details>
<summary>How do I stop the program?</summary>
Press <code>^C</code> or <code>c-C</code> or <kbd>Ctrl</kbd>+<kbd>C</kbd> to stop the program.
</details>

<details>
<summary>How do I get my friend to install it?</summary>
That's an issue (iss-you), not an iss-me. Get creative.
</details>

<details>
<summary>Why did you choose Go over Python?</summary>
Because Go compiles easily to an executable that also works cross-platform, so your unsuspecting friend doesn't need Go or Python or any external tools, which makes it easier for them to get pranked.
</details>

<details>
<summary>Will my friend realise how to stop it?</summary>
The terminal will open when program runs, you can minimize or hide it, but the icon will still be in the dock. If your friend is remotely smart, they will close the terminal, which stops the program.
</details>


<details>
<summary>Why is it not on Linux?</summary>
As far am I'm aware, there is no native text-to-speech command (like <code>say</code> on MacOS) we can use on Linux. It is possible to achieve this via a Go module or package, but I don't know much about Go, and wanted to keep the binary as small as possible.
</details>

<details>
<summary>Why don't you rebuild it in Rust?</summary>
Get out.
</details>
