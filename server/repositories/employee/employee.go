package repositories

import (
	"database/sql"
	"go-web-template/server/models"
)

type employeeRepo struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepo{
		DB: db,
	}
}

type EmployeeRepository interface {
	Save(employee *models.Employee) error
	FindAll() (*[]models.Employee, error)
	FindByID(id int) (*models.Employee, error)
	UpdateByID(id int) (*models.Employee, error)
	DeleteByID(id int) error
}

func (e *employeeRepo) Save(employee *models.Employee) error {
	// query ke DB
	// query ini menggunakan namanya params, yaitu $1, $2, dst...
	// jumlahnya sesuai dengan field yang diinginkan
	query := `
			INSERT INTO employees (
				 nip, name, address, created_at, updated_at
			) VALUES (
				$1, $2, $3, $4, $5,
			)
		`

	// kita akan melakukan prepare.
	// hal ini bertujuan untuk melakukan peningkatan performa
	// dan keamanan
	stmt, err := e.DB.Prepare(query)

	if err != nil {
		return err
	}

	// jangan lupa untuk menutup koneksi statements
	defer stmt.Close()

	// proses memasukkan parameter dari query tadi.
	// hal ini harus berurutan.
	_, err = stmt.Exec(
		employee.NIP, employee.Nama,
		employee.Address, employee.CreatedAt, employee.UpdatedAt,
	)

	return err

}

func (e *employeeRepo) FindAll() (*[]models.Employee, error) {
	query := `
		SELECT 
			id, nip, name, address, created_at, updated_at
		FROM
			employees
	`

	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// stmt.Query() berfungsi untuk mengambil semua data yang ada di database
	// jadi return nya akan lebih dari 1 data
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// mendefine employees sebagai slice of models.Employee
	var employees []models.Employee

	// melakukan perulangan
	// jadi karena data yang di return lebih dari 1 data, maka kita
	// perlu melakukan perulangan untuk mengambil datanya
	for rows.Next() {

		// ini berfungsi sebagai wadah untuk mengambil per 1 data
		var employee models.Employee

		// ini proses mengambil per 1 data, dan menambahkannya ke variable employee tadi
		// sebagai wadahnya
		err := rows.Scan(
			&employee.ID, &employee.NIP, &employee.Nama,
			&employee.Address, &employee.CreatedAt, &employee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// lalu kita akan menggunakan function append()
		// untuk menambah wadah tadi kedalam variable employees
		employees = append(employees, employee)
	}

	return &employees, nil

}

func (e *employeeRepo) FindByID(id int) (*models.Employee, error) {
	query := `
		SELECT 
			id, nip, name, address
		FROM
			employees
		WHERE
			id=$1
	`
	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// stmt.QueryRow()
	// berfungsi untuk menerima data dari database, dengan mengambil 1 data saja
	row := stmt.QueryRow(id)

	var employee models.Employee

	err = row.Scan(
		&employee.ID, &employee.NIP, &employee.Nama, &employee.Address,
	)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e *employeeRepo) UpdateByID(employee *models.Employee) error {
	query := `
		UPDATE employees
		SET name=$1, address=$2, nip=$3, updated_at=$4
		WHERE id=$5
	`

	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(employee.Nama, employee.Address, employee.NIP, employee.UpdatedAt, employee.ID)

	return err
}

func (e *employeeRepo) DeleteByID(id int) error {
	query := `
		DELETE FROM employees
		WHERE id=$1
	`

	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
