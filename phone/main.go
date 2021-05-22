package main

import (
	"bufio"
	//"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	db    = "numbers.db"
	input = "numbers.txt"
)

const schema = `
CREATE TABLE phonenumber (
	id integer primary key,
	number varchar(64)
);
`

func main() {
	d, err := newPhoneDB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	exists, err := d.exists()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !exists {
		numbers, err := load()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		err = upload(d, numbers)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err = normalizeAll(d); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func load() ([]string, error) {
	var numbers []string
	in, err := os.Open("numbers.txt")
	if err != nil {
		return numbers, err
	}
	defer in.Close()
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return numbers, nil
}

func upload(d *phoneDB, numbers []string) error {
	for _, n := range numbers {
		d.add(n)
	}
	return nil
}

func normalizeAll(d *phoneDB) error {
	numbers, err := d.getAll()
	if err != nil {
		return err
	}
	seen := make(map[string]struct{})
	for _, n := range numbers {
		number := normalize(n.Number)
		if _, ok := seen[number]; ok {
			d.delete(n.ID)
		} else {
			d.update(n.ID, number)
			seen[number] = struct{}{}
		}
	}
	return nil
}

func normalize(number string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(number, "")
}

type phoneDB struct {
	conn *sqlx.DB
}

type phone struct {
	ID     int
	Number string
}

func newPhoneDB() (*phoneDB, error) {
	d := phoneDB{}
	var err error
	d.conn, err = sqlx.Connect("sqlite3", db)
	return &d, err
}

func (d *phoneDB) exists() (bool, error) {
	_, err := d.conn.Exec(schema)
	if err == nil {
		return false, nil
	}
	if strings.Contains(err.Error(), "already exists") {
		return true, nil
	}
	return false, err
}

func (d *phoneDB) add(number string) {
	d.conn.MustExec(`INSERT INTO phonenumber (number) VALUES (?)`, number)
}

func (d *phoneDB) getAll() ([]phone, error) {
	numbers := []phone{}
	err := d.conn.Select(&numbers, "SELECT id,number FROM phonenumber")
	return numbers, err
}

func (d *phoneDB) delete(id int) {
	d.conn.MustExec(`DELETE FROM phonenumber WHERE id = ?`, id)
}

func (d *phoneDB) update(id int, number string) {
	d.conn.MustExec(`UPDATE phonenumber SET number = ? WHERE id = ?`, number, id)
}
