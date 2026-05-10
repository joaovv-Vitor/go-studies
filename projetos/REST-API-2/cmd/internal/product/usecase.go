package product

type Usecase struct {
	repo *Repository
}

func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetProducts() ([]Product, error) {
	return u.repo.GetProducts()
}

func (u *Usecase) CreateProduct(product Product) (Product, error) {
	productId, err := u.repo.CreateProduct(product)
	if err != nil {
		return Product{}, err
	}
	product.ID = productId

	return product, nil
}

func (u *Usecase) GetProductById(id_product int) (*Product, error) {
	product, err := u.repo.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *Usecase) DeleteProductByID(id int) error {
	return u.repo.DeleteProductByID(id)
}

func (u *Usecase) UpdateProduct(product Product) error {
	return u.repo.UpdateProduct(product)

}
