go-queue
========
[![Build Status](https://travis-ci.org/petrkotek/go-queue.svg?branch=master)](https://travis-ci.org/petrkotek/go-queue)
[![Coverage Status](https://coveralls.io/repos/petrkotek/go-queue/badge.svg?branch=master&service=github)](https://coveralls.io/github/petrkotek/go-vector?branch=master)

## About

Currently has only one implementation of a queue in golang:
- `OverflowingQueue`, which has specified capacity and starts discarding new items, when capacity is reached.

## License

Released under an MIT license. See `LICENSE.md` file for details.
