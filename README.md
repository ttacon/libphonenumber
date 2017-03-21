libphonenumber
==============

golang port of Google's libphonenumber

[![forthebadge](http://forthebadge.com/images/badges/no-ragrets.svg)](http://forthebadge.com)

[![Build Status](https://travis-ci.org/ttacon/libphonenumber.svg?branch=master)](https://travis-ci.org/ttacon/libphonenumber)
[![GoDoc](https://godoc.org/github.com/ttacon/libphonenumber?status.png)](https://godoc.org/github.com/ttacon/libphonenumber)

WARNING
=======

There is currently a lot going on, I started this a while ago and
recently picked it back up and got it functional. It was initially
translated from the Java version of libphonenumber, but I wasn't a
fan of always relying on loading proto files that were encoded in
ObjectStreams (so all metadata is embedded in the code, and I'm
exploring better ways to do this). 

Pull requests are of course welcome, but things will be moving fast
at first so they may not be accepted until I get this repo to a more
stable state - currently it is VERY fragile.

I should get it cleaned up soon (think a few days) in which case
I would love for help finishing this off :).

Examples
========

Super simple to use.

### To get a phone number

```go
num, err := libphonenumber.Parse("6502530000", "US")
```

### To format a number

```go
// num is a *libphonenumber.PhoneNumber
formattedNum := libphonenumber.Format(num, libphonenumber.NATIONAL)
```
