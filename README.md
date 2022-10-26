# Writing Secure Go code

AppsFlyer ∴  2022 <br />


Miki Tebeka <i class="far fa-envelope"></i> [miki@353solutions.com](mailto:miki@353solutions.com), <i class="fab fa-twitter"></i> [@tebeka](https://twitter.com/tebeka), <i class="fab fa-linkedin-in"></i> [mikitebeka](https://www.linkedin.com/in/mikitebeka/), <i class="fab fa-blogger-b"></i> [blog](https://www.ardanlabs.com/blog/)  

#### Shameless Plugs

- [Go Brain Teasers](https://pragprog.com/titles/d-gobrain/go-brain-teasers/)
- [LinkedIn Learning](https://www.linkedin.com/learning/instructors/miki-tebeka)

# Agenda

- Common security threats (OWASP top 10)
- Avoiding injection
- Secure HTTP requests
- Avoiding sensitive data leak
- Handling secrets
- The security mindset and adding security to your development process

[Terminal Log](terminal.log)

# Links

- [Let's talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) by Dave Cheney
- [Go Security Policy](https://golang.org/security)
- [Awesome security tools](https://github.com/guardrailsio/awesome-golang-security)
- [How our security team handles secrets](https://monzo.com/blog/2019/10/11/how-our-security-team-handle-secrets)
- HTTPS
    - `openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj /CN=localhost`
    - [x/crypto/autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)
    - [Using Let's Encrypt in Go](https://marcofranssen.nl/build-a-go-webserver-on-http-2-using-letsencrypt)
- [Customizing Binaries with Build Tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags)
- Books
    - [Security with Go](https://www.packtpub.com/product/security-with-go/9781788627917)
    - [Black Hat Go](https://nostarch.com/blackhatgo) book
- [Search for AWS keys in GitHub](https://sourcegraph.com/search?q=context:global+AWS_SECRET_ACCESS_KEY%3D%5B%27%22%5D.%7B40%7D%5B%27%22%5D&patternType=regexp)
- [Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing#The_fallacies)
- [cue](https://cuelang.org/) - Language for data validation
- Serialization Vulnerabilities
    - [XML Billion Laughs](https://en.wikipedia.org/wiki/Billion_laughs_attack) attack
    - [Java Parse Float](https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/)
- [Understanding HTML templates in Go](https://blog.lu4p.xyz/posts/golang-template-turbo/)
- SQL
    - [database/sql](https://golang.org/pkg/database/sql/)
    - [sqlx](https://github.com/jmoiron/sqlx)
    - [gorm](https://gorm.io/index.html)
- [Resilient net/http servers](https://ieftimov.com/post/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/)
- [Context](https://blog.golang.org/context) on the Go blog
- [Customizing Binaries with Build Tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags)
- [Our Software Depenedcy Problem](https://research.swtch.com/deps)
- [Go's CVE List](https://www.cvedetails.com/vulnerability-list/vendor_id-14185/Golang.html)
- Static tools
    - [golangci-lint](https://golangci-lint.run/)
    - [gosec](https://github.com/securego/gosec)
    - [staticcheck](https://staticcheck.io/)
    - Use [x/tools/analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) to write your own (see [here](https://github.com/tebeka/recheck) for an example)
- The new[embed](https://golang.org/pkg/embed/) package
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [The Security Mindset](https://www.schneier.com/blog/archives/2008/03/the_security_mi_1.html) by Bruce Schneier
- [Effective Go](https://golang.org/doc/effective_go.html) - Read this!

# Data & Other

- entries
    - [add-1.json](_ws/add-1.json)
    - [add-2.json](_ws/add-2.json)
    - [add-3.json](_ws/add-3.json)
- `docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=s3cr3t postgres:14-alpine`
- `docker exec -it <ID> psql -U postgres`
     - or `pgcli -p 5432 -U postgres -h localhost`
- `curl -d@add-1.json http://localhost:8080/new`
- `openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj /CN=localhost`
