Docker:
    docker build -t template-matcher .
    docker run -p 8080:8080 <tag>

Heroku:
    heroku login
    heroku container:login
    heroku create
    heroku container:push web
    heroku container:release web
    heroku open