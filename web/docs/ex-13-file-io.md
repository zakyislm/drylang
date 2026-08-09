# File Io

```javascript
// 13_file_io.y
// Demonstrates file system builtins

pt "--- Write File ---"
FILE "example.txt"
w(FILE, "Hello, file system!\nThis is line 2.")
pt "File created: " + FILE

pt "--- Read File ---"
content r(FILE)
pt "Content of file:"
pt content

pt "--- List Directory ---"
files dir(".")
pt "Files in current directory: " + str(len(files))
lp len(files) {
    // Only print first 5 files to avoid long output
    if i = 5 {
        pt "..."
        done
    }
    pt "- " + files[i]
}

pt "--- Delete File ---"
del(FILE)
pt "File deleted."

```
