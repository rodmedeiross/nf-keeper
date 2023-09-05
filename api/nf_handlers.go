package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodmedeiross/nf-keeper/types"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "NF Keeper",
	})
}

func NfHandler(c *fiber.Ctx) error {
	nfs := &[]types.NfItem{
		{	Id: 1, Name: "NF 1", Number: 345252345325 },
		{	Id: 2, Name: "NF 2", Number: 32432523425324 },
		{	Id: 3, Name: "NF 3", Number: 5234325234235 },
	}

	return c.Render("nfs", fiber.Map{
		"Nfs": nfs,
	})
}
