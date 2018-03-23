package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		f, err := ioutil.TempFile("", "vimuar")
		if err != nil {
			c.Error(err)
			return err
		}
		f.Close()
		text, pattern := c.FormValue("text"), c.FormValue("pattern")
		cmd := exec.Command("vim", "-Z", "-u", "NONE", "-N", "--clean", "--cmd", "so vimuar.vim")
		cmd.Env = append(os.Environ(),
			"VIMUAR_FILE="+f.Name(),
			"VIMUAR_TEXT="+text,
			"VIMUAR_PATTERN="+pattern,
		)
		b, err := cmd.CombinedOutput()
		if err != nil {
			c.Error(err)
			return err
		}
		b, err = ioutil.ReadFile(f.Name())
		if err != nil {
			c.Error(err)
			return err
		}
		return c.String(http.StatusOK, string(b))
	})
	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
