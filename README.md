# io.Writer to *os.File

## Usage

```
f := wtof.New(ansicolor.NewAnsiColorWriter(os.Stdout), 1024)
defer f.Close()
os.Stdout = f.File
```
- Must be close.

## License
MIT