# Interface

## hash.Hash
```golang

```
- `hash.Hash32`
- `hash.Hash64`

## io.Writer
```golang
Write(p []byte) (n int, err error)
```

## io.Reader
```golang
Read(p []byte) (n int, err error)
```

### io.Closer
```golang
Close() error
```

### io.Seeker
```golang
Seek(offset int64, whence int) (int64, error)
```


### io.ReadWriter
```golang
Read(p []byte) (n int, err error)
Write(p []byte) (n int, err error)
```

- `io.WriterCloser`
- `io.ReadWriteCloser`
- `io.ReadSeeker`
- `io.WriterSeeker`
- `io.ReadWriteSeeker`

### io.ReaderFrom
```golang
ReadFrom(r Reader) (n int64, err error)
```

### io.WriterTo
```golang
WriteTo(w Writer) (n int64, err error)
```

### io.ReaderAt
```golang
ReadAt(p []byte, off int64) (n int, err error)
```

### io.WriterAt
```golang
WriteAt(p []byte, off int64) (n int, err error)
```

### io.ByteReader
```golang
ReadByte() (byte, error)
```

### io.ByteScanner
```golang
ByteReader
UnreadByte() error
```