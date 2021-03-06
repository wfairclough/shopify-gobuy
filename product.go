package gobuy

import (
	"fmt"
)

const (
	getProductListingsPathTmpl     = "api/apps/{{.AppId}}/product_listings.json"
	getProductListingsTagsPathTmpl = "api/apps/{{.AppId}}/product_listings/tags.json"
)

type (
	Product struct {
		ProductID   int64            `json:"product_id"`
		CreatedAt   string           `json:"created_at"`
		UpdatedAt   string           `json:"updated_at"`
		BodyHTML    string           `json:"body_html"`
		Handle      string           `json:"handle"`
		ProductType string           `json:"product_type"`
		Title       string           `json:"title"`
		Vendor      string           `json:"vendor"`
		Available   bool             `json:"available"`
		Tags        string           `json:"tags"`
		PublishedAt string           `json:"published_at"`
		Images      []ProductImage   `json:"images"`
		Options     []ProductOption  `json:"options"`
		Variants    []ProductVariant `json:"variants"`
	}

	ProductImage struct {
		ID         int64         `json:"id"`
		CreatedAt  string        `json:"created_at"`
		Position   int           `json:"position"`
		UpdatedAt  string        `json:"updated_at"`
		ProductID  int64         `json:"product_id"`
		Src        string        `json:"src"`
		VariantIds []interface{} `json:"variant_ids"`
	}

	ProductOption struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		ProductID int64  `json:"product_id"`
		Position  int    `json:"position"`
	}

	ProductVariant struct {
		ID               int64                `json:"id"`
		Title            string               `json:"title"`
		OptionValues     []ProductOptionValue `json:"option_values"`
		Price            string               `json:"price"`
		CompareAtPrice   interface{}          `json:"compare_at_price"`
		FormattedPrice   string               `json:"formatted_price"`
		Grams            int                  `json:"grams"`
		RequiresShipping bool                 `json:"requires_shipping"`
		Sku              string               `json:"sku"`
		Barcode          string               `json:"barcode"`
		Taxable          bool                 `json:"taxable"`
		Position         int                  `json:"position"`
		Available        bool                 `json:"available"`
		CreatedAt        string               `json:"created_at"`
		UpdatedAt        string               `json:"updated_at"`
	}

	ProductOptionValue struct {
		OptionID int64  `json:"option_id"`
		Name     string `json:"name"`
		Value    string `json:"value"`
	}

	ProductService interface {
		GetProducts(page, limit int) ([]*Product, error)
	}

	ProductTag struct {
		ID    int64  `json:"id"`
		Title string `json:"title"`
	}
)

// GetProducts fetches the products with pagination
func (b *BuyClient) GetProducts(page, limit int) ([]*Product, error) {
	q := make(map[string]string)
	q["page"] = fmt.Sprintf("%d", page)
	q["limit"] = fmt.Sprintf("%d", limit)
	rsp, err := b.sendShopifyRequest(requestOptions{
		method:      "GET",
		urlTemplate: getProductListingsPathTmpl,
		queryParams: q,
	})
	if err != nil {
		return nil, err
	}
	list := &struct {
		ProductListings []*Product `json:"product_listings"`
	}{}
	err = rsp.JsonDecode(&list)
	return list.ProductListings, err
}

// GetProductByHandle fetches a product by a handle
func (b *BuyClient) GetProductByHandle(handle string) ([]*Product, error) {
	q := make(map[string]string)
	q["handle"] = handle
	rsp, err := b.sendShopifyRequest(requestOptions{
		method:      "GET",
		urlTemplate: getProductListingsPathTmpl,
		queryParams: q,
	})
	if err != nil {
		return nil, err
	}
	list := &struct {
		ProductListings []*Product `json:"product_listings"`
	}{}
	err = rsp.JsonDecode(&list)
	return list.ProductListings, err
}

// GetProductTags fetches the product tags with pagination
func (b *BuyClient) GetProductTags(page, limit int) ([]*ProductTag, error) {
	q := make(map[string]string)
	q["page"] = fmt.Sprintf("%d", page)
	q["limit"] = fmt.Sprintf("%d", limit)
	rsp, err := b.sendShopifyRequest(requestOptions{
		method:      "GET",
		urlTemplate: getProductListingsTagsPathTmpl,
		queryParams: q,
	})
	if err != nil {
		return nil, err
	}
	list := &struct {
		Tags []*ProductTag `json:"tags"`
	}{}
	err = rsp.JsonDecode(&list)
	return list.Tags, err
}
