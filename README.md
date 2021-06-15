# Data source
Random data collected from various websites and apis.

Currently this is ain't much, but honest work.

### Base url: `https://data.edwardbot.tk`
## Routes:
- `/` Health check route
- `/covid` Hungarian COVID-19 statistics collected by a web scraper
- `/meme/random` Returns a random meme from reddit(for some reason currently biased towards the first 10 entry)
- `/meme/count` Returns the count of memes on the server
- `/meme/all` Returns every meme(for some reason only returns 2/3 of it)

**DISCLAIMER**: I don't own any of this data. If you want me to remove something from it feel free to contact me on [Discord](https://dc.edwardbot.tk).


## Rate limits
There are no rate limits. Somewhat because I was too lazy to implement one.
As long as your requests don't take down the server, there is no limit.

## Stack
If you are too lazy to just open the `go.mod` file, here are what I used for this.
- [Golang](https://golang.org) for the language of it, because of its highly multithreaded nature
- [Gin-Gonic](https://gin-gonic.com) for the Web Server, because it offers speed and flexibility
- [goquery](https://github.com/PuerkitoBio/goquery) for the web scraper, because this was the fist result
- [Godotenv](https://github.com/joho/godotenv) because it is just cool

### Other stuff used
- [Reddit API](https://reddit.com) for the memes
- [koronavirus.gov.hu](https://koronavirus.gov.hu) for the COVID-19 stats
