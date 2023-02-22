package main

import (
	"encoding/json"
	"fmt"
	data2 "frontNoSql/internal/data"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	//snippets, err := app.snippets.Latest()
	//if err != nil {
	//	if errors.Is(err, models.ErrNoRecord) {
	//		app.notFound(w)
	//	} else {
	//		app.serverError(w, err)
	//	}
	//	return
	//}
	//
	data := app.newTemplateData(r)
	//data.Snippets = snippets
	req, err := http.NewRequest(http.MethodGet, "http://localhost:4000/v1/books?page_size=5&page=2", nil)
	if err != nil {
		app.serverError(w, err)
	}
	req.Header.Add("Authorization", "Bearer N6NHJXNXEYVBZALDMYK3RTL4WY")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		app.serverError(w, err)
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	result := struct {
		Books []*data2.Book
		data2.Metadata
	}{}
	err = dec.Decode(&result)
	if err != nil {
		app.serverError(w, err)
	}
	data.Books = result.Books
	app.render(w, http.StatusOK, "index.html", data)

}

//func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
//	params := httprouter.ParamsFromContext(r.Context())
//
//	id, err := strconv.Atoi(params.ByName("id"))
//	if err != nil || id < 1 {
//		app.notFound(w)
//		return
//	}
//	snippet, err := app.snippets.Get(id)
//	if err != nil {
//		if errors.Is(err, models.ErrNoRecord) {
//			app.notFound(w)
//		} else {
//			app.serverError(w, err)
//		}
//		return
//	}
//
//	data := app.newTemplateData(r)
//	data.Snippet = snippet
//
//	app.render(w, http.StatusOK, "view.tmpl", data)
//
//}
//
//func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
//	data := app.newTemplateData(r)
//
//	data.Form = snippetCreateForm{
//		Expires: 365,
//	}
//
//	app.render(w, http.StatusOK, "create.tmpl", data)
//}
//
//func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
//
//	var form snippetCreateForm
//
//	err := app.decodePostForm(r, &form)
//	if err != nil {
//		app.clientError(w, http.StatusBadRequest)
//		return
//	}
//
//	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
//	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
//	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
//	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")
//
//	if !form.Valid() {
//		data := app.newTemplateData(r)
//		data.Form = form
//		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
//		return
//	}
//
//	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//
//	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
//}

func (app *application) list(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.serverError(w, err)
	}

	data := app.newTemplateData(r)
	//data.Snippets = snippets
	url := fmt.Sprintf("http://localhost:4000/v1/books?page_size=10&page=%d", id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		app.serverError(w, err)
	}
	req.Header.Add("Authorization", "Bearer N6NHJXNXEYVBZALDMYK3RTL4WY")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		app.serverError(w, err)
	}
	fmt.Println(req)
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	result := struct {
		Books []*data2.Book
		data2.Metadata
	}{}
	err = dec.Decode(&result)
	if err != nil {
		app.serverError(w, err)
	}
	data.Books = result.Books
	app.render(w, http.StatusOK, "incss.html", data)
}

func (app *application) showBook(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.serverError(w, err)
	}
	data := app.newTemplateData(r)
	url := fmt.Sprintf("http://localhost:4000/v1/books/%d", id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		app.serverError(w, err)
	}
	req.Header.Add("Authorization", "Bearer N6NHJXNXEYVBZALDMYK3RTL4WY")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		app.serverError(w, err)
	}
	fmt.Println(req)
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	result := struct {
		Book *data2.Book
	}{}
	err = dec.Decode(&result)
	data.Book = result.Book
	fmt.Println(data.Book)
	data.Book.Pages = int(data.Book.Size)
	fmt.Println(data.Book.Pages)
	app.render(w, http.StatusOK, "q7.html", data)
}

func (app *application) loginGet(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "login.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	http.Redirect(w, r, "/", 301)
}

func (app *application) regGet(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "reg.html", nil)
}

func (app *application) reg(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	http.Redirect(w, r, "/", 301)
}
