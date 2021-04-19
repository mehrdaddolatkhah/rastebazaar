package business

import (
	"fmt"
	"html/template"
	"net/http"

	"rastebazaar/pkg/business/utils"
	"rastebazaar/pkg/domain"
	"rastebazaar/pkg/domain/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CategoryHandler will hold everything that controller needs
type CategoryHandler struct {
	categoryRepo domain.CategoryRepository
	tmpl         *template.Template
}

// NewCategoryHandler returns a new BaseHandler
func NewCategoryHandler(categoryRepo domain.CategoryRepository, tmpl *template.Template) *CategoryHandler {
	return &CategoryHandler{
		categoryRepo: categoryRepo,
		tmpl:         tmpl,
	}
}

// HelloCategory returns Hello, World
func (h *CategoryHandler) HelloCategory(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, Category"))
}

// GetCategoryListView , in this fucntion we will login as admin
func (h *CategoryHandler) GetCategoryListView(w http.ResponseWriter, r *http.Request) {

	err := h.tmpl.ExecuteTemplate(w, "admin-category-list.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GetAddCategoryView , in this fucntion we will login as admin
func (h *CategoryHandler) GetAddCategoryView(w http.ResponseWriter, r *http.Request) {

	var err error
	CategoriesForView := make(map[primitive.ObjectID]interface{})
	Categories, categoriesError := h.categoryRepo.GetCategories()

	if categoriesError != nil {
		fmt.Println(categoriesError.Error())
		return
	}

	if Categories == nil {
		err = h.tmpl.ExecuteTemplate(w, "admin-add-category.html", nil)
	} else {

		for _, categories := range Categories {

			CategoriesForView[categories.ID] = categories.Name

			// CategoriesForView = map[primitive.ObjectID]interface{}{
			// 	categories.ID: categories.Name,
			// }
		}

		fmt.Printf("%s", CategoriesForView)

		err = h.tmpl.ExecuteTemplate(w, "admin-add-category.html", CategoriesForView)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// PostAddCategoryView , in this fucntion we will login as admin
func (h *CategoryHandler) PostAddCategoryView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	categoryName := r.FormValue("et_category_name")
	parentCategory := r.Form["option_category"]

	fmt.Println(categoryName)
	fmt.Println(parentCategory)

	if categoryName == "" || parentCategory[0] == "" {
		utils.RespondJSON(w, 400, false, "all fields need filed", nil)
		return
	}

	category := database.Category{
		Name:   categoryName,
		Parent: parentCategory[0],
	}

	fmt.Printf("****** %s", category)

	_, err := h.categoryRepo.AddCategory(&category)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/secure/category/add-category", http.StatusSeeOther)
}
