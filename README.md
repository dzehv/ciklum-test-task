# Interview Task Ciklum

The provided URLs each return JSON. The first returns an array of objects representing articles, and the second returns an array of objects representing content marketing (i.e. paid articles). We would like you to create an API that makes a request to the two URLs, and returns JSON containing an array created from the two responses.

The returned array should consist of the first five articles from the article response, followed by the first content marketing from the content marketing response. This pattern should repeat, taking the next five articles followed by the next content marketing. If no more content marketing is available, you should continue the pattern, but insert an object in the place of content marketing with a single field, Type, that has the value “Ad”. This entire pattern, consisting of all articles, content marketing, and ads, should be returned from a single API request.

### Article URL:​ ​

https://storage.googleapis.com/aller-structure-task/articles.json

### Content Marketing URL:

https://storage.googleapis.com/aller-structure-task/contentmarketing.json

# Usage

``` bash
go run main.go
curl -v http://localhost:8080/ciklum-test
```
