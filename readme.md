
# GenericList

[![tag](https://img.shields.io/github/tag/squarehole/genericlist.svg)](https://github.com/squarehole/genericlist/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/squarehole/genericlist?status.svg)](https://pkg.go.dev/github.com/squarehole/genericlist)
![Build Status](https://github.com/squarehole/genericlist/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/squarehole/genericlist)](https://goreportcard.com/report/github.com/squarehole/genericlist)
[![Coverage](https://img.shields.io/codecov/c/github/squarehole/genericlist)](https://codecov.io/gh/squarehole/genericlist)
[![Contributors](https://img.shields.io/github/contributors/squarehole/genericlist)](https://github.com/squarehole/genericlist/graphs/contributors)
[![License](https://img.shields.io/github/license/squarehole/genericlist)](./LICENSE)




## 🚀 Install

```sh
go get github.com/squarehole/genericlist
```

**Compatibility**: go >= 1.21


## 💡 Usage

### Description

a versatile and type-safe generic list for any type that implements the 'comparable' contract, which allows for comparisons using

```go
func CreateList() {
    list := &GenericList[int]{}
    list.New()
    list.Add(1)
    list.Add(2)
	list.Add(3)
	
	// Increment the value of every item in the list with 1
    list.ForEach(func(i *int) { *i++ })
}
```

## 🤝 Contributing

- Ping me on mastodon [@jvanrhyn](https://mastodon.world/@jvanrhyn) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/squarehole/genericlist)
- Fix [open issues](https://github.com/squarehole/genericlist/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=squarehole/genericlist)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/jvanrhyn?style=for-the-badge)](https://github.com/sponsors/jvanrhyn)

## 📝 License

Copyright © 2023 [Johan van Rhyn](https://github.com/jvanrhyn).

This project is [MIT](./LICENSE) licensed.
