package controller

import (
	"database/sql"
	"errors"

	"github.com/yunfeizuo/liam/model"
)

// PackageController package apis.
type PackageController struct {
	DB *sql.DB
}

// GetPackage retrieve one package from DB by id
func (c *PackageController) GetPackage(id int64) (*model.Package, error) {
	pkg := model.Package{ID: id}
	err := c.DB.QueryRow(`SELECT 
			shipment_id, 
			barcode, 
			weight, 
			name,
			address, 
			cellphone 
		FROM package WHERE id = $1`,
		id).Scan(
		&pkg.ShipmentID,
		&pkg.Barcode,
		&pkg.Weight,
		&pkg.Name,
		&pkg.Address,
		&pkg.Cellphone,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &pkg, err
}

// CreatePackage insert a new row.
func (c *PackageController) CreatePackage(pkg *model.Package) error {
	if pkg.ID == 0 {
		query := `INSERT INTO package (shipment_id, barcode, weight, name, address, cellphone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
		return c.DB.QueryRow(query, pkg.ShipmentID, pkg.Barcode, pkg.Weight, pkg.Name, pkg.Address, pkg.Cellphone).Scan(&pkg.ID)
	}

	_, err := c.DB.Query(`INSERT INTO package (id, shipment_id, barcode, weight, name, address, cellphone) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		pkg.ID, pkg.ShipmentID, pkg.Barcode, pkg.Weight, pkg.Name, pkg.Address, pkg.Cellphone)
	return err
}

// UpdatePackage update the existing row.
func (c *PackageController) UpdatePackage(pkg *model.Package) error {
	if pkg.ID == 0 {
		return errors.New("Invalid input, requires a package id")
	}

	_, err := c.DB.Query(`UPDATE package SET shipment_id = $1, barcode = $2, weight = $3, name = $4, address = $5, cellphone = $6 WHERE id = $7`,
		pkg.ShipmentID, pkg.Barcode, pkg.Weight, pkg.Name, pkg.Address, pkg.Cellphone, pkg.ID)
	return err
}

// CreateOrUpdatePackage insert a new row if `pkg.ID == 0`, otherwise, update the existing row.
func (c *PackageController) CreateOrUpdatePackage(pkg *model.Package) error {
	if pkg.ID != 0 {
		return c.CreatePackage(pkg)
	}
	return c.UpdatePackage(pkg)
}

// Ship put a package into the next shipment by filling the shipment ID.
func (c *PackageController) Ship(pkg *model.Package) error {
	sc := ShipmentController{c.DB}
	shipment, err := sc.NextShipment()
	if err != nil {
		return err
	}

	pkg.ShipmentID = shipment.ID

	return nil
}
