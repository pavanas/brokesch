package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pavanas/brokestest/models"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Competitor)
// DB Table: Plural (competitors)
// Resource: Plural (Competitors)
// Path: Plural (/competitors)
// View Template Folder: Plural (/templates/competitors/)

// CompetitorsResource is the resource for the Competitor model
type CompetitorsResource struct {
	buffalo.Resource
}

// List gets all Competitors. This function is mapped to the path
// GET /competitors
func (v CompetitorsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	competitors := &models.Competitors{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Competitors from the DB
	if err := q.All(competitors); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the headers so clients know how to paginate.
	c.Response().Header().Set("X-Pagination", q.Paginator.String())

	return c.Render(200, r.JSON(competitors))
}

// Show gets the data for one Competitor. This function is mapped to
// the path GET /competitors/{competitor_id}
func (v CompetitorsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Competitor
	competitor := &models.Competitor{}

	// To find the Competitor the parameter competitor_id is used.
	if err := tx.Find(competitor, c.Param("competitor_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(competitor))
}

// New default implementation. Returns a 404
func (v CompetitorsResource) New(c buffalo.Context) error {
	return c.Error(404, errors.New("not available"))
}

// Create adds a Competitor to the DB. This function is mapped to the
// path POST /competitors
func (v CompetitorsResource) Create(c buffalo.Context) error {
	// Allocate an empty Competitor
	competitor := &models.Competitor{}

	// Bind competitor to the html form elements
	if err := c.Bind(competitor); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(competitor)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Render errors as JSON
		return c.Render(400, r.JSON(verrs))
	}

	return c.Render(201, r.JSON(competitor))
}

// Edit default implementation. Returns a 404
func (v CompetitorsResource) Edit(c buffalo.Context) error {
	return c.Error(404, errors.New("not available"))
}

// Update changes a Competitor in the DB. This function is mapped to
// the path PUT /competitors/{competitor_id}
func (v CompetitorsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Competitor
	competitor := &models.Competitor{}

	if err := tx.Find(competitor, c.Param("competitor_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Competitor to the html form elements
	if err := c.Bind(competitor); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(competitor)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Render errors as JSON
		return c.Render(400, r.JSON(verrs))
	}

	return c.Render(200, r.JSON(competitor))
}

// Destroy deletes a Competitor from the DB. This function is mapped
// to the path DELETE /competitors/{competitor_id}
func (v CompetitorsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Competitor
	competitor := &models.Competitor{}

	// To find the Competitor the parameter competitor_id is used.
	if err := tx.Find(competitor, c.Param("competitor_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(competitor); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON(competitor))
}
