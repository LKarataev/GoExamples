# GoCompare

## A Golang tool for comparing databases and finding differences

This project demonstrates how to compare two databases in different formats (JSON or XML) using Golang. The databases contain cook recipes with various ingredients and cooking times. The program reports any changes in the recipes, such as added, removed or modified ingredients or cakes. The project consists of three utilities: ReadDB, CompareDB and CompareFS. ReadDB reads a database file and converts it to another format. CompareDB reads two database files and compares them. CompareFS reads two text files and compares them line by line.

This project is intended to showcase my Golang skills and knowledge of databases.

## Installation

To install this project, you need to have Golang installed on your system. You can download it from [here](https://golang.org/dl/).

Then, you can clone this repository using the following command:

```bash
git clone https://github.com/LKarataev/GoCompare.git
```

## Usage

To use this project, you need to have some database files in JSON or XML format that contain cook recipes. You can find some examples in the `data` folder of this repository.

To run the ReadDB utility, use the following command:

```bash
go run ReadDB.go -f data/cakes.json
```

This will read the `cakes.json` file and convert it to XML format. You can also specify a different file name or format using the `-f` flag.

To run the CompareDB utility, use the following command:

```bash
go run CompareDB.go -old data/cakes.json -new data/cakes.xml
```

This will read the `cakes.json` and `cakes.xml` files and compare them. You can also specify different file names using the `-old` and `-new` flags.

To run the CompareFS utility, use the following command:

```bash
go run CompareFS.go -old data/files.txt -new data/files2.txt
```

This will read the `files.txt` and `files2.txt` files and compare them line by line. You can also specify different file names using the `-old` and `-new` flags.
