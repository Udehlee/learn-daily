## Overview 


In Go, templating utilizes packages such as text/template or html/template to dynamically generate content by merging predefined formats (templates) with actual data.



## golang spec

The input text for a template is UTF-8-encoded text in any format. "Actions"--data evaluations or control structures--are delimited by "{{" and "}}"; all text outside actions is copied to the output unchanged.