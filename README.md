# Clone of the [Go-Blog](https://github.com/matt-west/go-blog) project.

Original documentation can be found there.

## Differences

1) Deployable to GAE

2) Blog posts can be generated using the Markdown

3) Posts can use templates

Used the [blackfriday](https://github.com/russross/blackfriday) library for parsing Markdown

## How to use post generation

I am using this blog under Windows, so if you need Shell scripts, you will have to create them.

I) newpost.bat test-first-post -> will crete the test-first-post.md in the "markdown" forlder.
If you have "Meditor" set in the app.json then it will execute [Meditor] test-first-post.md

I am using the [Markdownpad 2](http://markdownpad.com/) and like it a lot so far.

Open test-first-post.md in the editor you like and edit it:

<pre>
------
layout: post -> you can change the layout of the post, but you will have to create your own templte then
Title: "Insert your title here" -> Change this to your title
Slug: "test-first-post" -> leave this as it is
Date: "2012-10-17 22:38" -> Change this to the date you want to be shown. Probably I will add this automatically later
comments: yes -> leave this as it is. This is not realized yet. 
Tags:
- "java" -> I think this is quite obvious ;)
- "job"
------
Place your text here please -> Write your awesome text here
</pre>

II) generate.bat test-first-post -> will generate HTML and change the posts.json

This project is deployed as a Demo [here](http://go-blog-gae.appspot.com/)

My own [Home Pages/Blog](http://konakov.info) is hosted on the same code.

## License (this section is not changed section of the original project)

Licensed under MIT.

Copyright &copy; 2012 Matt West

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
