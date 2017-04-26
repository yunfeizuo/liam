package controller

// type ProductController struct {
// 	DB *sql.DB
// }

// func (c *ProductController) CreateProduct(product *model.Product) error {
// 	urls := strings.Join(product.ImageURLs, ",")
// 	log.Println(product.ID, product.Title, urls)
// 	rows, err := c.DB.Query(`INSERT INTO products (id, title, img_urls) VALUES ($1, $2, $3)`,
// 		product.ID, product.Title, urls)
// 	if err != nil {
// 		log.Println("create product error", err)
// 		return err
// 	}
// 	defer rows.Close()
// 	return nil
// }

// func (c *ProductController) GetProducts() ([]model.Product, error) {
// 	products := make([]model.Product, 0, 256)
// 	rows, err := c.DB.Query(`SELECT id, title, img_urls FROM products LIMIT 256`)
// 	if err != nil {
// 		log.Println("query my product error", err)
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		product := model.Product{}
// 		var urls string
// 		if err = rows.Scan(&product.ID, &product.Title, &urls); err != nil {
// 			log.Println("scan my product error", err)
// 			return nil, err
// 		}
// 		product.ImageURLs = strings.Split(urls, ",")
// 		products = append(products, product)
// 	}

// 	return products, nil
// }
