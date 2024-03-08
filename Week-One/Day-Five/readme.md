## overview

Learned about making css styles available to html elements


## keep in mind

http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))