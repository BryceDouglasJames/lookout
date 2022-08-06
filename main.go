package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"

	web_driver "github.com/brycedouglasjames/lookout/pool"
)

type link struct {
	URL string
}

func init() {
	err := os.Remove("activity.log")
	if err != nil {
		fmt.Println("Could not delete file, was not created yet.")
	}
}

func main() {
	//specify static path
	path, _ := filepath.Abs("./templates/html")

	//serve all layout and schema files
	engine := html.New(path, ".html")

	//satrt fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//serve static files
	app.Static("/js", "./templates/js")
	app.Static("/css", "./templates/css")

	//generate routes
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		fmt.Println("Getting index...")
		return c.Render("index", fiber.Map{}, "app")
	})

	app.Post("/ping", func(c *fiber.Ctx) error {

		//make sure proper json heading is applied to request
		if string(c.Request().Header.ContentType()) != "application/json" {
			log.Println("Wrong header..")
			c.Status(500).Send([]byte("Wrong header information..."))
			return nil
		}

		//clean json request
		te, err := json.Marshal(string(c.Body()))
		if err != nil {
			log.Println("Error when ingesting request")
			return nil
		}
		a, _ := strconv.Unquote(string(te))
		link_reader := strings.NewReader(a)
		link_decoder := json.NewDecoder(link_reader)
		link_decoder.DisallowUnknownFields()

		//map url to link type
		var l link
		parse := link_decoder.Decode(&l)
		if parse != nil {
			log.Println("Cannot map URL to json object")
			return nil
		}
		err = link_decoder.Decode(&struct{}{})
		if err != io.EOF {
			log.Println("Only expecting one json object..")
			return nil
		}
		//fmt.Printf("link: %+v", l)

		//Start ping job process
		web_driver.Ingest_Url(l.URL)

		c.Response().BodyWriter().Write([]byte("Succesful search"))
		return nil
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {

		//This route is a test for now. Making sure I can reliably update a template sockket connection for
		//Rendering crawler updates.
		//Once user sessions are instantiated, we need to figure out a way to dedicate each socket connection
		//to dedicated user logs or simply query against a DB to get user search updates.
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("WS ingress: %s", msg)
			fmt.Println(mt)
			file, err := ioutil.ReadFile("activity.log")
			if err != nil {
				os.Create("activity.log")
			}

			s := string(file)
			lines := strings.Split(s, "\n")
			//fmt.Println("Sending file content")

			if err = c.WriteJSON(lines); err != nil {
				log.Println("WS Error: ", err)
				break
			}
		}

	}))

	log.Fatal(app.Listen(":3000"))
}
