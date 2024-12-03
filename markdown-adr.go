package main

import (
    "os"
    "text/template"
)

type ADRData struct {
    ID            string
    Title         string
    Status        string
    Context       string
    Decision      string
    Consequences string
}

adrData := ADRData{
    ID:            "0001",
    Title:         "Adopt Markdown for ADRs",
    Status:        "Accepted",
    Context:       "We need a consistent and easy-to-use format for documenting architecture decisions.",
    Decision:      "We will use Markdown files for storing ADRs.",
    Consequences:  "ADRs will be easy to read and edit.",
}

func main() {
    // Create a new template and parse the template file
    tmpl, err := template.ParseFiles("adr_template.tmpl")
    if err != nil {
        panic(err)
    }

    // Execute the template with the ADR data, writing the output to stdout
    err = tmpl.ExecuteTemplate(os.Stdout, "adr", adrData)
    if err != nil {
        panic(err)
    }
}

