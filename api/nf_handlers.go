package api

import (
	"strings"
	"time"

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
		{	Id: 1, Name: "NF 1", DateIn: string(time.Now().Format("02/01/2006")), NfNumber: "12347654856574354757475774523" },
		{	Id: 2, Name: "NF 2", DateIn: string(time.Now().Format("02/01/2006")), NfNumber: "12347654856574354757475774523" },
		{	Id: 3, Name: "NF 3", DateIn: string(time.Now().Format("02/01/2006")), NfNumber: "12347654856574354757475774523" },
	}

	return c.Render("nfs", fiber.Map{
		"ModalId": "#cadastro",
		"Nfs": nfs,
	})
}

func ModalHandler(c *fiber.Ctx) error {
	id := c.Params("operation")
	title := strings.ToUpper(id[:1]) + id[1:]

	return c.Render("modal", fiber.Map{
		"Title": title,
		"ModalId": id,
	})
}
