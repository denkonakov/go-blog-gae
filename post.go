package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"os/exec"
	"regexp"
	"io/ioutil"
	"text/template"
	"encoding/json"
	"github.com/russross/blackfriday"
)

// We will parse *.md file here
type Mrkdwn map[string]string

var (
	HeaderRE = regexp.MustCompile("(?s)^(------\n(.+)\n------\n)")
	AttrsRE = regexp.MustCompile("(?Um)^([^\\:]+?)\\: (.+)$")
	TagsRE = regexp.MustCompile("(?m)^- (.+)$")

	IncludeRE = regexp.MustCompile("(?sU){% include (\\S+?) %}")

  	postTemplate *template.Template
)

func main() {
	args := os.Args;
	if (len(args) > 1) {
		
		post_template := config.Mfolder + "/" + config.Nptemplate;
		new_post := config.Mfolder + "/" + args[2]+".md"

		switch args[1] {
			case "newp":
				if (posts[args[2]] != nil) {
					fmt.Printf("Post %q is already exists. Please choose another name.", args[2])	
				} else {
					fmt.Printf("Copy %q to the %q\n", post_template, new_post);
					f, err := os.Open(post_template)
					if err != nil {
					    die("Cannot open the post template file.\n")
					}

					fout, err := os.Create(new_post)
					if err != nil {
					    die("Cannot open the target (new) template file.\n")
					}

					scanner := bufio.NewScanner(f)
					for scanner.Scan() {
						line := scanner.Text()
						fout.WriteString(strings.TrimSpace(line) + "\n")
					    if (strings.Contains(line, "Title:")) {
							fout.WriteString("Slug: \"" + args[2] + "\"\n")
					    } 
					}

					f.Close()
					fout.Close()

					// If the Editor is not set - do not try to start it
					if (len(config.Meditor) > 0) {
						fmt.Printf("Trying to open the Markdown editor with newly created post.\n")
						pwd, _ := os.Getwd()
						cmd := exec.Command(config.Meditor, pwd + "/" + new_post)
						er := cmd.Run()
						if er != nil {
							die("We  were not able to run default Markdown editor.\n")
						}
					}
				}

			case "generate":	
		        fmt.Printf("Parsing the markup %q\n", new_post)
						mrkp := load_markdown(new_post)

		        // does 'include' function
				funcs := template.FuncMap{"include": func(s string) string {
					return *load_file("templates/"+s)
				}}

				fmt.Printf("Loading the template %q\n", "templates/"+mrkp["layout"]+".html")
		        postTemplate := template.Must(template.New(cut_surrounding_quotes(mrkp["Slug"])).Funcs(
		        	funcs).Parse(*load_file("templates/"+mrkp["layout"]+".html")))

		        // Adds new one as a first one
		        if (posts[cut_surrounding_quotes(mrkp["Slug"])] != nil) {
		        	delete(posts, cut_surrounding_quotes(mrkp["Slug"]))
		        }

				p := []Post{Post{cut_surrounding_quotes(mrkp["Title"]),
					cut_surrounding_quotes(mrkp["Slug"]), cut_surrounding_quotes(mrkp["Date"]),
					cut_surrounding_quotes(mrkp["Date"]), "", cut_surrounding_quotes(mrkp["Title"]), 
					strings.Split(mrkp["Tags"], ", ")}}	
				
				for _, post := range posts {
					p = append(p, *post)
				}

				// Make it string and more readable
				mr, _ := json.Marshal(p)
				mr2 := string(mr)
				mr2 = strings.Replace(mr2, "[", "[\n", -1)	
				mr2 = strings.Replace(mr2, "{", "\t{\n", -1)	
				mr2 = strings.Replace(mr2, ",", ",\n", -1)
				mr2 = strings.Replace(mr2, "]", "]\n", -1)	
				mr2 = strings.Replace(mr2, "}", "\t}\n", -1)

				fmt.Printf("Saving posts map to the posts.json\n");

				postsf, err := os.Create("data/posts.json")
		        if err != nil {
		            die("Cannot create the posts.json file.\n")
		        }
		        postsf.WriteString(mr2)
		        postsf.Close()

		        htmlPostFile := "posts/"+cut_surrounding_quotes(mrkp["Slug"])+".html"
		        fmt.Printf("Creating the HTML post %q\n", htmlPostFile)
		        fout, err := os.Create(htmlPostFile)
		        if err != nil {
		            die("Cannot create the new post HTML file.\n")
		        }

		        postTemplate.Execute(fout, string(blackfriday.MarkdownCommon([]byte(mrkp["content"]))))
		        fout.Close()

			default:
				printUsage()	
		}	
	} else {
		printUsage()
	}
}

func load_markdown(name string) Mrkdwn {
	markup := Mrkdwn{}

	// Process the "---"-type of the header.
	content := *load_file(name)
	if m := HeaderRE.FindStringSubmatch(content); m != nil {		
		header := m[2]

		if m := AttrsRE.FindAllStringSubmatch(header, -1); m != nil {
			for _, pair := range m {
				markup[pair[1]] = pair[2]
			}
		} else {
		  	die2("Bad headers in [%#v]", header)
		}

		if m := TagsRE.FindAllStringSubmatch(header, -1); m != nil {
			tags := make([]string, 0)
			for _, g := range m {
				tags = append(tags, cut_surrounding_quotes(g[1]))
			}

			markup["Tags"] = strings.Join(tags, ", ")
		}

		content = HeaderRE.ReplaceAllLiteralString(content, "")
	}

	markup["content"] = content
	return markup
}

// This function loads up a fileand removes '\r' from its content.
func load_file(filename string) *string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		die2("Unable to read file [%s]", filename)
	}

	s := strings.Replace(string(bytes), "\r", "", -1)
	return &s
}

func printUsage() {
	fmt.Printf("--- usage example ---\n");
	fmt.Printf("run post.go [newp|generate] {post name}\n");
}

func die(msg string) {
	fmt.Printf(msg)
    os.Exit(1)
}

func die2(format string, v ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format+"\n", v...))
	os.Exit(1)
}

func cut_surrounding_quotes(s string) string {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") ||
		strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'") {
		return s[1 : len(s)-1]
	}

	return s
}